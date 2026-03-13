// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";

// =========================================================
// INTERFACES
// =========================================================

interface IEarnVaultReader {
    function totalVaultAssets() external view returns (uint256);
    function totalWeightedUSDT() external view returns (uint256);
    function totalUsers() external view returns (uint256);
    function rewardRate() external view returns (uint256);
    function usdtBalance(address user) external view returns (uint256);
    function weightedUSDT(address user) external view returns (uint256);
    function getPendingCPP(address user) external view returns (uint256);
    function getUserBoostBps(address user) external view returns (uint256);
    function estimatedDailyCPP(address user) external view returns (uint256);
    function positions(address user) external view returns (
        uint256 depositedAt,
        uint256 lastActionAt
    );
    function emergencyMode() external view returns (bool);
}

interface IVeCPOTLockerReader {
    struct LockPosition {
        uint256 lockId;
        uint256 amount;
        uint256 veAmount;
        uint256 startTime;
        uint256 unlockTime;
        uint256 durationDays;
        bool active;
    }
    function getTotalVeCPOT(address user) external view returns (uint256);
    function getBoostBps(address user) external view returns (uint256);
    function getLockPositions(address user) external view returns (LockPosition[] memory);
}

interface INFTBoostControllerReader {
    function getNFTBoostBps(address user) external view returns (uint256);
    function getActiveNFT(address user) external view returns (
        uint256 tokenId,
        uint256 level,
        uint256 boostBps
    );
}

interface IRewardDistributorReader {
    function getVaultRewardRate() external view returns (uint256);
}

/**
 * @title ChapoolVaultReader
 * @notice Aggregates read-only data from all Chapool Earn contracts into a single call.
 *         Designed for frontend dashboards and DefiLlama TVL adapters.
 *
 * @dev Non-upgradeable. Deploy a new Reader if underlying contract addresses change.
 */
contract ChapoolVaultReader is Ownable {

    // =========================================================
    // STORAGE
    // =========================================================

    IEarnVaultReader          public earnVault;
    IVeCPOTLockerReader       public vecpotLocker;
    INFTBoostControllerReader public nftBoostController;
    IRewardDistributorReader  public rewardDistributor;

    // =========================================================
    // RETURN STRUCTS
    // =========================================================

    struct ProtocolOverview {
        uint256 tvl;                // USDT locked in vault
        uint256 totalWeightedUSDT;  // Sum of all users' weighted USDT
        uint256 totalUsers;         // Depositor count
        uint256 rewardRate;         // CPP per second (global)
        uint256 dailyCPPEmission;   // rewardRate × 86400
        bool    emergencyMode;
    }

    struct UserDashboard {
        uint256 usdtBalance;        // Raw USDT deposited
        uint256 weightedUSDT;       // usdtBalance × boost multiplier
        uint256 pendingCPP;         // Claimable CPP
        uint256 estimatedDailyCPP;  // CPP earned per day at current rate
        uint256 boostBps;           // Total boost (veCPOT + NFT) in bps
        uint256 vecpotBoostBps;     // veCPOT boost component
        uint256 nftBoostBps;        // NFT boost component
        uint256 totalVeCPOT;        // Sum of veAmount across active locks
        uint256 lockedCPOT;         // Total CPOT locked across active positions
        uint256 depositedAt;        // First deposit timestamp
        uint256 lastActionAt;       // Most recent action timestamp
    }

    // =========================================================
    // CONSTRUCTOR
    // =========================================================

    constructor(
        address _earnVault,
        address _vecpotLocker,
        address _nftBoostController,
        address _rewardDistributor,
        address _owner
    ) Ownable(_owner) {
        earnVault          = IEarnVaultReader(_earnVault);
        vecpotLocker       = IVeCPOTLockerReader(_vecpotLocker);
        nftBoostController = INFTBoostControllerReader(_nftBoostController);
        rewardDistributor  = IRewardDistributorReader(_rewardDistributor);
    }

    // =========================================================
    // PROTOCOL OVERVIEW
    // =========================================================

    /**
     * @notice Global protocol snapshot. Primary entry point for DefiLlama and dashboards.
     */
    function getProtocolOverview() external view returns (ProtocolOverview memory o) {
        o.tvl               = _getTVL();
        o.totalWeightedUSDT = _safeTotalWeightedUSDT();
        o.totalUsers        = _safeTotalUsers();
        o.rewardRate        = _safeRewardRate();
        o.dailyCPPEmission  = o.rewardRate * 1 days;
        o.emergencyMode     = _safeEmergencyMode();
    }

    /**
     * @notice Minimal TVL — for DefiLlama adapter.
     */
    function getTVL() external view returns (uint256) {
        return _getTVL();
    }

    // =========================================================
    // USER DASHBOARD
    // =========================================================

    /**
     * @notice Full user position in a single RPC call.
     */
    function getUserDashboard(address user) external view returns (UserDashboard memory dash) {
        if (address(earnVault) != address(0)) {
            try earnVault.usdtBalance(user) returns (uint256 b) { dash.usdtBalance = b; } catch {}
            try earnVault.weightedUSDT(user) returns (uint256 w) { dash.weightedUSDT = w; } catch {}
            try earnVault.getPendingCPP(user) returns (uint256 c) { dash.pendingCPP = c; } catch {}
            try earnVault.estimatedDailyCPP(user) returns (uint256 d) { dash.estimatedDailyCPP = d; } catch {}
            try earnVault.getUserBoostBps(user) returns (uint256 b) { dash.boostBps = b; } catch {}
            try earnVault.positions(user) returns (uint256 depositedAt, uint256 lastActionAt) {
                dash.depositedAt  = depositedAt;
                dash.lastActionAt = lastActionAt;
            } catch {}
        }

        if (address(vecpotLocker) != address(0)) {
            try vecpotLocker.getBoostBps(user) returns (uint256 b) { dash.vecpotBoostBps = b; } catch {}
            try vecpotLocker.getTotalVeCPOT(user) returns (uint256 ve) { dash.totalVeCPOT = ve; } catch {}
            try vecpotLocker.getLockPositions(user) returns (
                IVeCPOTLockerReader.LockPosition[] memory lockList
            ) {
                for (uint256 i = 0; i < lockList.length; i++) {
                    if (lockList[i].active && block.timestamp < lockList[i].unlockTime) {
                        dash.lockedCPOT += lockList[i].amount;
                    }
                }
            } catch {}
        }

        if (address(nftBoostController) != address(0)) {
            try nftBoostController.getNFTBoostBps(user) returns (uint256 b) { dash.nftBoostBps = b; } catch {}
        }
    }

    /**
     * @notice Returns a user's CPOT lock positions (for lock list UI).
     */
    function getUserLockPositions(address user)
        external
        view
        returns (IVeCPOTLockerReader.LockPosition[] memory)
    {
        if (address(vecpotLocker) == address(0)) {
            return new IVeCPOTLockerReader.LockPosition[](0);
        }
        return vecpotLocker.getLockPositions(user);
    }

    /**
     * @notice Returns user's active NFT boost info.
     */
    function getUserNFTBoost(address user)
        external
        view
        returns (uint256 tokenId, uint256 level, uint256 boostBps)
    {
        if (address(nftBoostController) == address(0)) return (0, 0, 0);
        try nftBoostController.getActiveNFT(user) returns (
            uint256 tid, uint256 lvl, uint256 bps
        ) {
            return (tid, lvl, bps);
        } catch {
            return (0, 0, 0);
        }
    }

    // =========================================================
    // ADMIN
    // =========================================================

    function setEarnVault(address addr) external onlyOwner { earnVault = IEarnVaultReader(addr); }
    function setVecpotLocker(address addr) external onlyOwner { vecpotLocker = IVeCPOTLockerReader(addr); }
    function setNftBoostController(address addr) external onlyOwner { nftBoostController = INFTBoostControllerReader(addr); }
    function setRewardDistributor(address addr) external onlyOwner { rewardDistributor = IRewardDistributorReader(addr); }

    // =========================================================
    // INTERNAL HELPERS
    // =========================================================

    function _getTVL() internal view returns (uint256) {
        if (address(earnVault) == address(0)) return 0;
        try earnVault.totalVaultAssets() returns (uint256 v) { return v; } catch { return 0; }
    }

    function _safeTotalWeightedUSDT() internal view returns (uint256) {
        if (address(earnVault) == address(0)) return 0;
        try earnVault.totalWeightedUSDT() returns (uint256 v) { return v; } catch { return 0; }
    }

    function _safeTotalUsers() internal view returns (uint256) {
        if (address(earnVault) == address(0)) return 0;
        try earnVault.totalUsers() returns (uint256 u) { return u; } catch { return 0; }
    }

    function _safeRewardRate() internal view returns (uint256) {
        if (address(earnVault) == address(0)) return 0;
        try earnVault.rewardRate() returns (uint256 r) { return r; } catch { return 0; }
    }

    function _safeEmergencyMode() internal view returns (bool) {
        if (address(earnVault) == address(0)) return false;
        try earnVault.emergencyMode() returns (bool e) { return e; } catch { return false; }
    }
}
