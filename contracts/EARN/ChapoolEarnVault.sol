// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../CPP/interfaces/ICPPToken.sol";
import "../CPP/interfaces/IAccountManager.sol";

/**
 * @title ChapoolEarnVault
 * @notice Pure TVL custody vault. Users deposit USDT and earn CPP rewards.
 *
 *  Reward mechanism (Synthetix staking pattern):
 *  - A global `rewardRate` (CPP per second) is set by admin.
 *  - Each user accrues CPP proportional to their "weighted USDT":
 *
 *      userWeightedUSDT = usdtBalance × (10000 + veCPOTBoost + nftBoost) / 10000
 *      userCPP/sec      = rewardRate × userWeightedUSDT / totalWeightedUSDT
 *
 *  - CPP is minted on-demand when the user calls claimCPP().
 *  - CPP is minted directly to the user's AA account (mirrors Staking contract).
 *  - No USDT yield is distributed — deposited USDT is returned 1:1 on withdrawal.
 *
 *  Boost sources (read-only, no custody):
 *  - VeCPOTLocker:       +1% / +2% / +3.5% / +5% by lock duration
 *  - NFTBoostController: +0.5% / +1% / +2% / +3.5% by NFT level (unstaked only)
 *
 *  Emergency withdrawal: owner can extract TVL after a 24-hour timelock.
 */
contract ChapoolEarnVault is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable
{
    using SafeERC20 for IERC20;

    // =========================================================
    // CONSTANTS
    // =========================================================

    uint256 public constant PRECISION = 1e18;
    uint256 public constant BPS_DENOMINATOR = 10000;
    uint256 public constant EMERGENCY_TIMELOCK = 24 hours;

    // =========================================================
    // STORAGE
    // =========================================================

    /// @notice Underlying asset (USDT)
    IERC20 public asset;

    /// @notice CPP token — minted to users on claimCPP()
    ICPPToken public cppToken;

    /// @notice AccountManager — resolves user EOA → AA account
    IAccountManager public accountManagerContract;

    /// @notice VeCPOTLocker — queried for per-user boost
    address public vecpotLocker;

    /// @notice NFTBoostController — queried for per-user NFT boost
    address public nftBoostController;

    // ---- CPP reward accumulator (Synthetix pattern) ----

    /// @notice CPP minted per second across all weighted-USDT (set by admin)
    uint256 public rewardRate;

    /// @notice Timestamp of last accumulator update
    uint256 public lastUpdateTime;

    /// @notice Accumulated CPP per weighted-USDT unit (× PRECISION)
    uint256 public accCPPPerWeightedUSDT;

    /// @notice Sum of all users' weighted USDT
    uint256 public totalWeightedUSDT;

    // ---- Per-user state ----

    /// @notice Raw USDT balance per user (1:1 with deposited amount)
    mapping(address => uint256) public usdtBalance;

    /// @notice Per-user weighted USDT (cached, updated on deposit/withdraw/syncBoost)
    mapping(address => uint256) public weightedUSDT;

    /// @notice Accumulator snapshot at last user settlement
    mapping(address => uint256) public rewardDebt;

    /// @notice Accrued but unclaimed CPP per user
    mapping(address => uint256) public pendingCPP;

    // ---- Position metadata (for display) ----

    struct UserVaultPosition {
        uint256 depositedAt;   // timestamp of first deposit
        uint256 lastActionAt;  // timestamp of most recent action
    }

    mapping(address => UserVaultPosition) public positions;

    /// @notice Number of unique depositors
    uint256 public totalUsers;

    // ---- Emergency withdrawal ----

    bool public emergencyMode;

    struct PendingWithdraw {
        uint256 amount;
        address to;
        uint256 scheduledAt;
        bool pending;
    }

    PendingWithdraw public pendingEmergencyWithdraw;

    // =========================================================
    // EVENTS
    // =========================================================

    event Deposited(address indexed user, uint256 assets, uint256 timestamp);
    event Withdrawn(address indexed user, uint256 assets, uint256 timestamp);
    event CPPRewardClaimed(address indexed user, address indexed aaAccount, uint256 amount, uint256 timestamp);
    event RewardRateSet(uint256 cppPerSecond, uint256 timestamp);
    event BoostSynced(address indexed user, uint256 oldWeighted, uint256 newWeighted);
    event VecpotLockerSet(address indexed locker);
    event NftBoostControllerSet(address indexed controller);
    event AccountManagerSet(address indexed accountManager);

    event EmergencyModeSet(bool enabled, uint256 timestamp);
    event EmergencyWithdrawInitiated(uint256 amount, address indexed to, uint256 executeAfter);
    event EmergencyWithdrawExecuted(uint256 amount, address indexed to, uint256 timestamp);
    event EmergencyWithdrawCancelled(uint256 timestamp);

    // =========================================================
    // INITIALIZER
    // =========================================================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @param _asset          USDT token address
     * @param _cppToken       CPP token address (must grant MINTER_ROLE to this contract)
     * @param _accountManager AccountManager contract address
     * @param _owner          Contract owner
     */
    function initialize(
        address _asset,
        address _cppToken,
        address _accountManager,
        address _owner
    ) public initializer {
        require(_asset != address(0), "Zero asset");
        require(_cppToken != address(0), "Zero cpp");
        require(_accountManager != address(0), "Zero accountManager");
        require(_owner != address(0), "Zero owner");

        __Ownable_init(_owner);
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __Pausable_init();

        asset = IERC20(_asset);
        cppToken = ICPPToken(_cppToken);
        accountManagerContract = IAccountManager(_accountManager);
        lastUpdateTime = block.timestamp;
    }

    // =========================================================
    // DEPOSIT / WITHDRAW
    // =========================================================

    /**
     * @notice Deposit USDT. Starts accruing CPP immediately.
     */
    function deposit(uint256 assets, address receiver)
        external
        nonReentrant
        whenNotPaused
        returns (uint256)
    {
        require(assets > 0, "Zero assets");
        require(receiver != address(0), "Zero receiver");

        _updateAccumulator();
        _settleUser(receiver);

        asset.safeTransferFrom(msg.sender, address(this), assets);
        usdtBalance[receiver] += assets;

        _syncUserWeight(receiver);

        UserVaultPosition storage pos = positions[receiver];
        if (pos.depositedAt == 0) {
            pos.depositedAt = block.timestamp;
            totalUsers += 1;
        }
        pos.lastActionAt = block.timestamp;

        emit Deposited(receiver, assets, block.timestamp);
        return assets;
    }

    /**
     * @notice Withdraw USDT. Returns exactly the deposited amount (no yield on USDT).
     */
    function withdraw(uint256 assets, address receiver)
        external
        nonReentrant
        returns (uint256)
    {
        require(assets > 0, "Zero assets");
        require(receiver != address(0), "Zero receiver");
        require(usdtBalance[msg.sender] >= assets, "Insufficient balance");

        _updateAccumulator();
        _settleUser(msg.sender);

        usdtBalance[msg.sender] -= assets;
        _syncUserWeight(msg.sender);

        positions[msg.sender].lastActionAt = block.timestamp;

        asset.safeTransfer(receiver, assets);
        emit Withdrawn(msg.sender, assets, block.timestamp);
        return assets;
    }

    // =========================================================
    // CPP REWARDS
    // =========================================================

    /**
     * @notice Claim all accrued CPP. Minted directly to the caller's AA account.
     */
    function claimCPP() external nonReentrant returns (uint256 amount) {
        _updateAccumulator();
        _settleUser(msg.sender);

        amount = pendingCPP[msg.sender];
        require(amount > 0, "No CPP to claim");

        pendingCPP[msg.sender] = 0;
        positions[msg.sender].lastActionAt = block.timestamp;

        address aaAccount = _getAAAccount(msg.sender);
        cppToken.mint(aaAccount, amount);

        emit CPPRewardClaimed(msg.sender, aaAccount, amount, block.timestamp);
    }

    /**
     * @notice Preview claimable CPP for a user (includes unaccrued since lastUpdateTime).
     */
    function getPendingCPP(address user) external view returns (uint256) {
        uint256 acc = accCPPPerWeightedUSDT;
        if (totalWeightedUSDT > 0) {
            uint256 elapsed = block.timestamp - lastUpdateTime;
            acc += (elapsed * rewardRate * PRECISION) / totalWeightedUSDT;
        }
        uint256 wUsdt = weightedUSDT[user];
        if (wUsdt == 0) return pendingCPP[user];
        return pendingCPP[user] + (wUsdt * (acc - rewardDebt[user])) / PRECISION;
    }

    // =========================================================
    // BOOST SYNC
    // =========================================================

    /**
     * @notice Re-sync a user's weighted USDT after their boost changes.
     *         Called automatically by VeCPOTLocker and NFTBoostController,
     *         or manually by the user after activating/deactivating an NFT.
     */
    function syncBoost(address user) external {
        _updateAccumulator();
        _settleUser(user);
        _syncUserWeight(user);
    }

    // =========================================================
    // VIEW
    // =========================================================

    /**
     * @notice Total USDT locked in the vault (= TVL for DefiLlama).
     */
    function totalVaultAssets() external view returns (uint256) {
        return asset.balanceOf(address(this));
    }

    /**
     * @notice Current boost in bps for a user (veCPOT + NFT combined).
     */
    function getUserBoostBps(address user) external view returns (uint256) {
        return _getUserBoostBps(user);
    }

    /**
     * @notice Estimated daily CPP for a user at current rewardRate.
     */
    function estimatedDailyCPP(address user) external view returns (uint256) {
        if (totalWeightedUSDT == 0) return 0;
        return (rewardRate * weightedUSDT[user] * 1 days) / totalWeightedUSDT;
    }

    // =========================================================
    // EMERGENCY WITHDRAWAL
    // =========================================================

    function setEmergencyMode(bool enabled) external onlyOwner {
        emergencyMode = enabled;
        if (enabled) {
            _pause();
        } else {
            _unpause();
            if (pendingEmergencyWithdraw.pending) {
                delete pendingEmergencyWithdraw;
                emit EmergencyWithdrawCancelled(block.timestamp);
            }
        }
        emit EmergencyModeSet(enabled, block.timestamp);
    }

    function initiateEmergencyWithdraw(uint256 amount, address to) external onlyOwner {
        require(emergencyMode, "Not in emergency mode");
        require(amount > 0, "Zero amount");
        require(to != address(0), "Zero destination");
        require(!pendingEmergencyWithdraw.pending, "Withdraw already pending");
        require(amount <= asset.balanceOf(address(this)), "Exceeds vault assets");

        pendingEmergencyWithdraw = PendingWithdraw({
            amount: amount,
            to: to,
            scheduledAt: block.timestamp,
            pending: true
        });

        emit EmergencyWithdrawInitiated(amount, to, block.timestamp + EMERGENCY_TIMELOCK);
    }

    function executeEmergencyWithdraw() external onlyOwner nonReentrant {
        PendingWithdraw memory pw = pendingEmergencyWithdraw;
        require(pw.pending, "No pending withdraw");
        require(block.timestamp >= pw.scheduledAt + EMERGENCY_TIMELOCK, "Timelock not elapsed");

        delete pendingEmergencyWithdraw;
        asset.safeTransfer(pw.to, pw.amount);

        emit EmergencyWithdrawExecuted(pw.amount, pw.to, block.timestamp);
    }

    function cancelEmergencyWithdraw() external onlyOwner {
        require(pendingEmergencyWithdraw.pending, "No pending withdraw");
        delete pendingEmergencyWithdraw;
        emit EmergencyWithdrawCancelled(block.timestamp);
    }

    // =========================================================
    // ADMIN
    // =========================================================

    /**
     * @notice Set the CPP emission rate. Updates accumulator first.
     * @param cppPerSecond CPP (in wei) emitted per second across all depositors
     */
    function setRewardRate(uint256 cppPerSecond) external onlyOwner {
        _updateAccumulator();
        rewardRate = cppPerSecond;
        emit RewardRateSet(cppPerSecond, block.timestamp);
    }

    function pause() external onlyOwner { _pause(); }
    function unpause() external onlyOwner { _unpause(); }

    function setVecpotLocker(address locker) external onlyOwner {
        vecpotLocker = locker;
        emit VecpotLockerSet(locker);
    }

    function setNftBoostController(address controller) external onlyOwner {
        nftBoostController = controller;
        emit NftBoostControllerSet(controller);
    }

    function setAccountManager(address _accountManager) external onlyOwner {
        require(_accountManager != address(0), "Zero address");
        accountManagerContract = IAccountManager(_accountManager);
        emit AccountManagerSet(_accountManager);
    }

    function _authorizeUpgrade(address newImpl) internal override onlyOwner {}

    // =========================================================
    // INTERNAL — ACCUMULATOR
    // =========================================================

    function _updateAccumulator() internal {
        if (totalWeightedUSDT == 0) {
            lastUpdateTime = block.timestamp;
            return;
        }
        uint256 elapsed = block.timestamp - lastUpdateTime;
        if (elapsed > 0) {
            accCPPPerWeightedUSDT += (elapsed * rewardRate * PRECISION) / totalWeightedUSDT;
            lastUpdateTime = block.timestamp;
        }
    }

    function _settleUser(address user) internal {
        uint256 wUsdt = weightedUSDT[user];
        if (wUsdt > 0 && accCPPPerWeightedUSDT > rewardDebt[user]) {
            pendingCPP[user] += (wUsdt * (accCPPPerWeightedUSDT - rewardDebt[user])) / PRECISION;
        }
        rewardDebt[user] = accCPPPerWeightedUSDT;
    }

    /**
     * @dev Recalculate and apply user's weighted USDT based on current boosts.
     *      Must be called AFTER _settleUser to avoid reward miscalculation.
     */
    function _syncUserWeight(address user) internal {
        uint256 oldWeighted = weightedUSDT[user];
        uint256 boostBps = _getUserBoostBps(user);
        uint256 newWeighted = (usdtBalance[user] * (BPS_DENOMINATOR + boostBps)) / BPS_DENOMINATOR;

        totalWeightedUSDT = totalWeightedUSDT - oldWeighted + newWeighted;
        weightedUSDT[user] = newWeighted;
        rewardDebt[user] = accCPPPerWeightedUSDT;

        emit BoostSynced(user, oldWeighted, newWeighted);
    }

    // =========================================================
    // INTERNAL — BOOST QUERIES
    // =========================================================

    function _getUserBoostBps(address user) internal view returns (uint256 boostBps) {
        if (vecpotLocker != address(0)) {
            try IVeCPOTLocker(vecpotLocker).getBoostBps(user) returns (uint256 b) {
                boostBps += b;
            } catch {}
        }
        if (nftBoostController != address(0)) {
            try INFTBoostController(nftBoostController).getNFTBoostBps(user) returns (uint256 b) {
                boostBps += b;
            } catch {}
        }
    }

    // =========================================================
    // INTERNAL — AA ACCOUNT RESOLUTION
    // =========================================================

    function _getAAAccount(address user) internal view returns (address) {
        address masterSigner = accountManagerContract.getDefaultMasterSigner();
        address aaAccount = accountManagerContract.getAccountAddress(user, masterSigner);
        require(aaAccount != address(0), "AA account not found");
        return aaAccount;
    }
}

// =========================================================
// MINIMAL INTERFACES (cross-contract calls)
// =========================================================

interface IVeCPOTLocker {
    function getBoostBps(address user) external view returns (uint256);
}

interface INFTBoostController {
    function getNFTBoostBps(address user) external view returns (uint256);
}
