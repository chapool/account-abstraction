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
 * @notice Lock CPOT to earn veCPOT weight and receive an APY boost on Chapool Earn.
 *
 *  This contract is a BOOST PROVIDER ONLY.
 *  CPP reward distribution and claiming are handled entirely by ChapoolEarnVault.
 *
 *  Key design points:
 *  - Each lock() call creates an independent position (lockId), allowing multiple concurrent locks.
 *  - veCPOT = sum of all active positions' veAmount.
 *  - APY boost = tier of the LONGEST active lock duration (not cumulative).
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

    // Boost in basis points per duration tier
    uint256 public constant BOOST_30_BPS  = 100;  // +1.0%
    uint256 public constant BOOST_90_BPS  = 200;  // +2.0%
    uint256 public constant BOOST_180_BPS = 350;  // +3.5%
    uint256 public constant BOOST_360_BPS = 500;  // +5.0%

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
     * @notice APY boost in bps — determined by the longest active lock duration.
     */
    function getBoostBps(address user) external view returns (uint256) {
        return _boostBpsForDuration(_getMaxDuration(user));
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
     * @notice Preview boost if user adds a new lock with newDurationDays.
     */
    function previewBoostBps(address user, uint256 newDurationDays)
        external
        view
        returns (uint256)
    {
        _requireValidDurationPure(newDurationDays);
        uint256 currentMax = _getMaxDuration(user);
        uint256 effectiveMax = newDurationDays > currentMax ? newDurationDays : currentMax;
        return _boostBpsForDuration(effectiveMax);
    }

    // =========================================================
    // ADMIN
    // =========================================================

    function setVault(address _vault) external onlyOwner {
        vault = _vault;
        emit VaultSet(_vault);
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

    function _getMaxDuration(address user) internal view returns (uint256 maxD) {
        LockPosition[] storage positions = lockPositions[user];
        for (uint256 i = 0; i < positions.length; i++) {
            if (positions[i].active && block.timestamp < positions[i].unlockTime) {
                if (positions[i].durationDays > maxD) maxD = positions[i].durationDays;
            }
        }
    }

    function _boostBpsForDuration(uint256 d) internal pure returns (uint256) {
        if (d >= DURATION_360) return BOOST_360_BPS;
        if (d >= DURATION_180) return BOOST_180_BPS;
        if (d >= DURATION_90)  return BOOST_90_BPS;
        if (d >= DURATION_30)  return BOOST_30_BPS;
        return 0;
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
