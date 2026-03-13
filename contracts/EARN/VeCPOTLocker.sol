// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title VeCPOTLocker
 * @notice Lock CPOT to earn veCPOT weight and receive a CPP reward boost on Chapool Earn.
 *
 *  This contract is a BOOST PROVIDER ONLY.
 *  CPP reward distribution and claiming are handled entirely by ChapoolEarnVault.
 *
 *  Key design points:
 *  - Each lock() call creates an independent position (lockId), allowing multiple concurrent locks.
 *  - veCPOT = sum of all active positions' (amount × durationDays / 360).
 *  - Boost (bps) = min(totalVeCPOT_units × boostPerVeUnit, maxVecpotBoostBps).
 *    Both lock amount AND duration contribute — more CPOT locked longer → more boost.
 *  - After any lock state change, the vault's syncBoost(user) is called to update
 *    the global weighted-USDT accumulator in ChapoolEarnVault.
 */
contract VeCPOTLocker is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable
{
    using SafeERC20 for IERC20;

    // =========================================================
    // CONSTANTS
    // =========================================================

    uint256 public constant DURATION_30  = 30;
    uint256 public constant DURATION_90  = 90;
    uint256 public constant DURATION_180 = 180;
    uint256 public constant DURATION_360 = 360;

    /// @dev Precision divisor for the boost formula.
    ///      boostBps = (veUnits × boostPerVeUnit) / BOOST_VE_PRECISION
    ///      At default boostPerVeUnit=1 this means 1 bps per 10 veCPOT units,
    ///      requiring 10× more CPOT to reach the same boost compared to precision=1.
    uint256 public constant BOOST_VE_PRECISION = 10;

    // =========================================================
    // STRUCTS
    // =========================================================

    struct LockPosition {
        uint256 lockId;       // User-scoped monotonic ID
        uint256 amount;       // CPOT locked
        uint256 veAmount;     // veCPOT weight for this position
        uint256 startTime;
        uint256 unlockTime;
        uint256 durationDays;
        bool active;
    }

    // =========================================================
    // STORAGE
    // =========================================================

    IERC20 public cpot;

    /// @notice ChapoolEarnVault — notified after every lock state change
    address public vault;

    /// @notice Boost earned per 1 veCPOT unit (1 unit = 1e18 wei), in basis points.
    ///         Default: 1 bps per veCPOT unit. Admin-tunable.
    uint256 public boostPerVeUnit;

    /// @notice Maximum CPOT boost cap in basis points (default 500 = +5%).
    uint256 public maxVecpotBoostBps;

    mapping(address => LockPosition[]) public lockPositions;
    mapping(address => uint256) public nextLockId;

    // =========================================================
    // EVENTS
    // =========================================================

    event LockedCPOT(
        address indexed user,
        uint256 lockId,
        uint256 amount,
        uint256 unlockTime,
        uint256 veAmount,
        uint256 durationDays
    );
    event UnlockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 timestamp);
    event MoreCPOTLocked(address indexed user, uint256 lockId, uint256 additionalAmount, uint256 newVeAmount);
    event VaultSet(address indexed vault);
    event BoostParamsSet(uint256 boostPerVeUnit, uint256 maxVecpotBoostBps);

    // =========================================================
    // INITIALIZER
    // =========================================================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @param _cpot  CPOT token address
     * @param _owner Contract owner
     */
    function initialize(address _cpot, address _owner) public initializer {
        require(_cpot != address(0), "Zero cpot");
        require(_owner != address(0), "Zero owner");

        __Ownable_init(_owner);
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();

        cpot = IERC20(_cpot);

        // Default: 1 bps per veCPOT unit, capped at +5%
        boostPerVeUnit    = 1;
        maxVecpotBoostBps = 500;
    }

    // =========================================================
    // LOCK
    // =========================================================

    /**
     * @notice Lock CPOT for a chosen duration. Creates an independent position.
     * @param amount      CPOT amount to lock
     * @param durationDays 30 / 90 / 180 / 360
     * @return lockId  New position ID (user-scoped)
     */
    function lock(uint256 amount, uint256 durationDays)
        external
        nonReentrant
        returns (uint256 lockId)
    {
        require(amount > 0, "Zero amount");
        _requireValidDuration(durationDays);

        cpot.safeTransferFrom(msg.sender, address(this), amount);

        uint256 veAmount = _calcVeAmount(amount, durationDays);
        uint256 unlockTime = block.timestamp + durationDays * 1 days;
        lockId = nextLockId[msg.sender]++;

        lockPositions[msg.sender].push(LockPosition({
            lockId: lockId,
            amount: amount,
            veAmount: veAmount,
            startTime: block.timestamp,
            unlockTime: unlockTime,
            durationDays: durationDays,
            active: true
        }));

        _notifyVault(msg.sender);

        emit LockedCPOT(msg.sender, lockId, amount, unlockTime, veAmount, durationDays);
    }

    // =========================================================
    // UNLOCK
    // =========================================================

    /**
     * @notice Unlock an expired position and return CPOT to sender.
     */
    function unlock(uint256 lockId) external nonReentrant {
        (uint256 idx, LockPosition storage pos) = _findActiveLock(msg.sender, lockId);
        require(block.timestamp >= pos.unlockTime, "Lock not expired");

        uint256 amount = pos.amount;
        pos.active = false;
        _removePosition(msg.sender, idx);

        _notifyVault(msg.sender);

        cpot.safeTransfer(msg.sender, amount);
        emit UnlockedCPOT(msg.sender, lockId, amount, block.timestamp);
    }

    // =========================================================
    // LOCK MORE
    // =========================================================

    /**
     * @notice Add more CPOT to an existing position. Unlock time is unchanged;
     *         veCPOT is recalculated on remaining duration.
     */
    function lockMore(uint256 lockId, uint256 additionalAmount) external nonReentrant {
        require(additionalAmount > 0, "Zero amount");

        (, LockPosition storage pos) = _findActiveLock(msg.sender, lockId);
        require(block.timestamp < pos.unlockTime, "Lock expired");

        cpot.safeTransferFrom(msg.sender, address(this), additionalAmount);

        pos.amount += additionalAmount;
        uint256 remainingDays = (pos.unlockTime - block.timestamp) / 1 days;
        if (remainingDays == 0) remainingDays = 1;
        pos.veAmount = _calcVeAmount(pos.amount, remainingDays);

        _notifyVault(msg.sender);

        emit MoreCPOTLocked(msg.sender, lockId, additionalAmount, pos.veAmount);
    }

    // =========================================================
    // VIEW — veCPOT & BOOST
    // =========================================================

    /**
     * @notice Sum of veAmounts for all active (non-expired) positions.
     */
    function getTotalVeCPOT(address user) external view returns (uint256) {
        return _calcTotalVeCPOT(user);
    }

    /**
     * @notice Boost in bps based on total veCPOT.
     *         boostBps = min(totalVeCPOT_units × boostPerVeUnit, maxVecpotBoostBps)
     *         where totalVeCPOT_units = totalVeCPOT / 1e18.
     */
    function getBoostBps(address user) external view returns (uint256) {
        return _calcBoostBps(_calcTotalVeCPOT(user));
    }

    /**
     * @notice All lock positions for a user.
     */
    function getLockPositions(address user) external view returns (LockPosition[] memory) {
        return lockPositions[user];
    }

    /**
     * @notice A specific lock position by lockId.
     */
    function getLockPosition(address user, uint256 lockId)
        external
        view
        returns (LockPosition memory)
    {
        LockPosition[] storage positions = lockPositions[user];
        for (uint256 i = 0; i < positions.length; i++) {
            if (positions[i].lockId == lockId) return positions[i];
        }
        revert("Lock not found");
    }

    /**
     * @notice Preview veCPOT for a hypothetical lock.
     */
    function previewVeCPOT(uint256 amount, uint256 durationDays) external pure returns (uint256) {
        _requireValidDurationPure(durationDays);
        return _calcVeAmount(amount, durationDays);
    }

    /**
     * @notice Preview boost if user adds a new lock of `additionalAmount` for `durationDays`.
     *         Returns the projected total boost bps (current positions + hypothetical new lock).
     */
    function previewBoostBps(address user, uint256 additionalAmount, uint256 durationDays)
        external
        view
        returns (uint256)
    {
        _requireValidDurationPure(durationDays);
        uint256 currentVe  = _calcTotalVeCPOT(user);
        uint256 newVe      = _calcVeAmount(additionalAmount, durationDays);
        return _calcBoostBps(currentVe + newVe);
    }

    // =========================================================
    // ADMIN
    // =========================================================

    function setVault(address _vault) external onlyOwner {
        vault = _vault;
        emit VaultSet(_vault);
    }

    /**
     * @notice Tune the veCPOT boost formula parameters.
     * @param _boostPerVeUnit    Bps earned per 1 veCPOT unit (1e18 wei = 1 unit).
     * @param _maxVecpotBoostBps Maximum CPOT boost cap in bps.
     */
    function setBoostParams(uint256 _boostPerVeUnit, uint256 _maxVecpotBoostBps) external onlyOwner {
        require(_boostPerVeUnit > 0, "Zero coeff");
        require(_maxVecpotBoostBps > 0 && _maxVecpotBoostBps <= 2000, "Invalid cap");
        boostPerVeUnit    = _boostPerVeUnit;
        maxVecpotBoostBps = _maxVecpotBoostBps;
        emit BoostParamsSet(_boostPerVeUnit, _maxVecpotBoostBps);
    }

    function _authorizeUpgrade(address newImpl) internal override onlyOwner {}

    // =========================================================
    // INTERNAL — VAULT NOTIFICATION
    // =========================================================

    /**
     * @dev Notify ChapoolEarnVault to re-sync the user's boost weight.
     *      Uses a try/catch so locker never reverts if vault is not yet set.
     */
    function _notifyVault(address user) internal {
        if (vault == address(0)) return;
        try IEarnVault(vault).syncBoost(user) {} catch {}
    }

    // =========================================================
    // INTERNAL — POSITION MANAGEMENT
    // =========================================================

    function _findActiveLock(address user, uint256 lockId)
        internal
        view
        returns (uint256 idx, LockPosition storage pos)
    {
        LockPosition[] storage positions = lockPositions[user];
        for (uint256 i = 0; i < positions.length; i++) {
            if (positions[i].lockId == lockId && positions[i].active) {
                return (i, positions[i]);
            }
        }
        revert("Lock not found or inactive");
    }

    function _removePosition(address user, uint256 idx) internal {
        LockPosition[] storage arr = lockPositions[user];
        uint256 last = arr.length - 1;
        if (idx != last) arr[idx] = arr[last];
        arr.pop();
    }

    // =========================================================
    // INTERNAL — MATH
    // =========================================================

    function _calcVeAmount(uint256 amount, uint256 durationDays) internal pure returns (uint256) {
        return (amount * durationDays) / 360;
    }

    function _calcTotalVeCPOT(address user) internal view returns (uint256 total) {
        LockPosition[] storage positions = lockPositions[user];
        for (uint256 i = 0; i < positions.length; i++) {
            if (positions[i].active && block.timestamp < positions[i].unlockTime) {
                total += positions[i].veAmount;
            }
        }
    }

    /**
     * @dev boostBps = min((veUnits × boostPerVeUnit) / BOOST_VE_PRECISION, maxVecpotBoostBps)
     *      Default: 1 bps per 10 veCPOT units → need 5,000 veCPOT units to reach +5% cap.
     */
    function _calcBoostBps(uint256 totalVe) internal view returns (uint256) {
        uint256 veUnits = totalVe / 1e18;
        uint256 bps     = (veUnits * boostPerVeUnit) / BOOST_VE_PRECISION;
        return bps > maxVecpotBoostBps ? maxVecpotBoostBps : bps;
    }

    function _requireValidDuration(uint256 d) internal pure {
        require(
            d == DURATION_30 || d == DURATION_90 || d == DURATION_180 || d == DURATION_360,
            "Invalid duration"
        );
    }

    function _requireValidDurationPure(uint256 d) internal pure {
        require(d == 30 || d == 90 || d == 180 || d == 360, "Invalid duration");
    }
}

interface IEarnVault {
    function syncBoost(address user) external;
}
