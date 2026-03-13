// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title ChapoolRewardDistributor
 * @notice Operations contract for managing the CPP emission rate in ChapoolEarnVault.
 *
 *  Flow:
 *    Platform USDT revenue is collected off-chain.
 *    Operations wallet calls setVaultRewardRate() to update CPP/second emission.
 *    ChapoolEarnVault mints CPP on-demand when users call claimCPP().
 *
 *  No tokens are transferred through this contract — it is a pure control plane.
 *
 * @dev Two roles:
 *      - owner: configure vault address, rescue tokens
 *      - REWARD_DEPOSITOR_ROLE: call setVaultRewardRate (operations wallet)
 */
contract ChapoolRewardDistributor is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable
{
    using SafeERC20 for IERC20;

    // =========================================================
    // STORAGE
    // =========================================================

    address public earnVault;

    mapping(address => bool) public isRewardDepositor;

    // Analytics
    uint256 public totalCppAllocated;   // cumulative CPP/sec × duration notified

    // =========================================================
    // EVENTS
    // =========================================================

    event VaultRewardRateSet(address indexed sender, uint256 cppPerSecond, uint256 timestamp);
    event RewardDepositorSet(address indexed account, bool enabled);
    event EarnVaultSet(address indexed vault);

    // =========================================================
    // MODIFIERS
    // =========================================================

    modifier onlyRewardDepositor() {
        require(isRewardDepositor[msg.sender] || msg.sender == owner(), "Not depositor");
        _;
    }

    // =========================================================
    // INITIALIZER
    // =========================================================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _owner) public initializer {
        require(_owner != address(0), "Zero owner");

        __Ownable_init(_owner);
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
    }

    // =========================================================
    // REWARD RATE CONTROL
    // =========================================================

    /**
     * @notice Set the CPP emission rate on ChapoolEarnVault.
     *
     *         The operations team calculates the rate off-chain based on:
     *           cppPerSecond = (platformUSDTRevenue / 0.002) / periodSeconds
     *         (since 1 CPP = 0.002 USDT)
     *
     * @param cppPerSecond CPP (in wei, 18 decimals) emitted per second
     */
    function setVaultRewardRate(uint256 cppPerSecond) external onlyRewardDepositor {
        require(earnVault != address(0), "Vault not set");
        IEarnVault(earnVault).setRewardRate(cppPerSecond);
        emit VaultRewardRateSet(msg.sender, cppPerSecond, block.timestamp);
    }

    /**
     * @notice Pause the vault emission (set rate to 0).
     */
    function pauseEmission() external onlyRewardDepositor {
        require(earnVault != address(0), "Vault not set");
        IEarnVault(earnVault).setRewardRate(0);
        emit VaultRewardRateSet(msg.sender, 0, block.timestamp);
    }

    // =========================================================
    // VIEW
    // =========================================================

    function getVaultRewardRate() external view returns (uint256) {
        if (earnVault == address(0)) return 0;
        try IEarnVault(earnVault).rewardRate() returns (uint256 r) { return r; } catch { return 0; }
    }

    // =========================================================
    // ADMIN
    // =========================================================

    function setEarnVault(address _vault) external onlyOwner {
        require(_vault != address(0), "Zero address");
        earnVault = _vault;
        emit EarnVaultSet(_vault);
    }

    function setRewardDepositor(address account, bool enabled) external onlyOwner {
        isRewardDepositor[account] = enabled;
        emit RewardDepositorSet(account, enabled);
    }

    function rescueToken(address token, address to, uint256 amount) external onlyOwner {
        IERC20(token).safeTransfer(to, amount);
    }

    function _authorizeUpgrade(address newImpl) internal override onlyOwner {}
}

interface IEarnVault {
    function setRewardRate(uint256 cppPerSecond) external;
    function rewardRate() external view returns (uint256);
}
