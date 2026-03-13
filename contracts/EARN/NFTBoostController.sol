// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title NFTBoostController
 * @notice Computes an APY boost for Chapool Earn users based on their CPNFT holdings.
 *
 *  Rules (Phase 2 — "hold to boost"):
 *  - No staking or custody required; the user simply holds the NFT in their wallet.
 *  - Each user may register ONE "active" tokenId as their boost NFT.
 *  - Ownership is verified at read time: if the NFT is sold, boost automatically returns to 0.
 *  - Multiple NFTs are NOT cumulative — only the registered (or highest held) NFT counts.
 *  - NFT levels map to CPP boost:
 *      NORMAL / C  → 0 bps
 *      B (L1)      → 10 bps  (+0.1%)
 *      A (L2)      → 20 bps  (+0.2%)
 *      S (L3)      → 50 bps  (+0.5%)
 *      SS (L4)     → 350 bps (+3.5%)
 *      SSS         → 500 bps (+5.0%, future tier)
 *
 *  @dev CPNFT does not implement ERC721Enumerable, so we use a registration model:
 *       users call activateBoost(tokenId) to designate their boost NFT.
 *       getNFTBoostBps re-verifies ownership at call time.
 */
// =========================================================
// INTERFACES
// =========================================================

/// @dev Matches CPNFT.NFTLevel enum: NORMAL(0), C(1), B(2), A(3), S(4), SS(5), SSS(6)
interface ICPNFT {
    function ownerOf(uint256 tokenId) external view returns (address);
    function getTokenLevel(uint256 tokenId) external view returns (uint8);
    function isStaked(uint256 tokenId) external view returns (bool);
}

contract NFTBoostController is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    // =========================================================
    // CONSTANTS — boost in bps per NFT level index
    // =========================================================

    // Index corresponds to CPNFT.NFTLevel enum value
    uint256 private constant LEVEL_COUNT = 7;
    uint256[7] private _levelBoostBps;

    // Readable constants
    uint256 public constant BOOST_NORMAL_BPS = 0;    // NORMAL
    uint256 public constant BOOST_C_BPS      = 0;    // C
    uint256 public constant BOOST_B_BPS      = 10;   // B  = L1  (+0.1%)
    uint256 public constant BOOST_A_BPS      = 20;   // A  = L2  (+0.2%)
    uint256 public constant BOOST_S_BPS      = 50;   // S  = L3  (+0.5%)
    uint256 public constant BOOST_SS_BPS     = 350;  // SS = L4  (+3.5%)
    uint256 public constant BOOST_SSS_BPS    = 500;  // SSS      (+5.0%, future)

    // =========================================================
    // STORAGE
    // =========================================================

    ICPNFT public cpnft;

    /// @notice User's currently registered boost tokenId. 0 = not registered.
    mapping(address => uint256) public activeTokenId;
    mapping(address => bool)    public hasActiveToken;

    // =========================================================
    // EVENTS
    // =========================================================

    event BoostActivated(address indexed user, uint256 indexed tokenId, uint8 level, uint256 boostBps);
    event BoostDeactivated(address indexed user, uint256 indexed tokenId);

    // =========================================================
    // INITIALIZER
    // =========================================================

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _cpnft, address _owner) public initializer {
        require(_cpnft != address(0), "Zero cpnft");
        require(_owner != address(0), "Zero owner");

        __Ownable_init(_owner);
        __UUPSUpgradeable_init();

        cpnft = ICPNFT(_cpnft);

        // CPNFT.NFTLevel: NORMAL=0, C=1, B=2, A=3, S=4, SS=5, SSS=6
        _levelBoostBps[0] = BOOST_NORMAL_BPS;
        _levelBoostBps[1] = BOOST_C_BPS;
        _levelBoostBps[2] = BOOST_B_BPS;
        _levelBoostBps[3] = BOOST_A_BPS;
        _levelBoostBps[4] = BOOST_S_BPS;
        _levelBoostBps[5] = BOOST_SS_BPS;
        _levelBoostBps[6] = BOOST_SSS_BPS;
    }

    // =========================================================
    // USER ACTIONS
    // =========================================================

    /**
     * @notice Register a tokenId as your active boost NFT.
     *         You must own the NFT at the time of registration.
     * @param tokenId The CPNFT token ID to activate
     */
    function activateBoost(uint256 tokenId) external {
        require(cpnft.ownerOf(tokenId) == msg.sender, "Not token owner");
        require(!cpnft.isStaked(tokenId), "NFT is staked");

        activeTokenId[msg.sender] = tokenId;
        hasActiveToken[msg.sender] = true;

        uint8 level = cpnft.getTokenLevel(tokenId);
        uint256 boostBps = _levelBoostBps[level < LEVEL_COUNT ? level : 0];

        emit BoostActivated(msg.sender, tokenId, level, boostBps);
    }

    /**
     * @notice Deregister your active boost NFT (e.g., before selling it).
     */
    function deactivateBoost() external {
        require(hasActiveToken[msg.sender], "No active token");
        uint256 tokenId = activeTokenId[msg.sender];
        delete activeTokenId[msg.sender];
        delete hasActiveToken[msg.sender];
        emit BoostDeactivated(msg.sender, tokenId);
    }

    // =========================================================
    // VIEW
    // =========================================================

    /**
     * @notice Returns the APY boost in bps for a user.
     * @dev Re-verifies on-chain ownership. If user sold the NFT, returns 0.
     */
    function getNFTBoostBps(address user) external view returns (uint256) {
        if (!hasActiveToken[user]) return 0;

        uint256 tokenId = activeTokenId[user];

        // Verify current ownership
        try cpnft.ownerOf(tokenId) returns (address tokenOwner) {
            if (tokenOwner != user) return 0;
        } catch {
            return 0;
        }

        // Staked NFTs do not provide Earn boost (cannot double-dip Staking + Earn)
        try cpnft.isStaked(tokenId) returns (bool staked) {
            if (staked) return 0;
        } catch {}

        uint8 level = cpnft.getTokenLevel(tokenId);
        if (level >= LEVEL_COUNT) return 0;
        return _levelBoostBps[level];
    }

    /**
     * @notice Returns the NFT level of the active tokenId and its boost bps.
     *         Returns (0, 0) if no active token or ownership lost.
     */
    function getActiveNFT(address user)
        external
        view
        returns (uint256 tokenId, uint256 level, uint256 boostBps)
    {
        if (!hasActiveToken[user]) return (0, 0, 0);

        tokenId = activeTokenId[user];

        try cpnft.ownerOf(tokenId) returns (address tokenOwner) {
            if (tokenOwner != user) return (tokenId, 0, 0);
        } catch {
            return (tokenId, 0, 0);
        }

        try cpnft.isStaked(tokenId) returns (bool staked) {
            if (staked) return (tokenId, 0, 0);
        } catch {}

        uint8 lvl = cpnft.getTokenLevel(tokenId);
        level = lvl;
        boostBps = lvl < LEVEL_COUNT ? _levelBoostBps[lvl] : 0;
    }

    /**
     * @notice Returns the boost bps for a given NFT level index (read-only helper for UI).
     */
    function getLevelBoostBps(uint8 level) external view returns (uint256) {
        if (level >= LEVEL_COUNT) return 0;
        return _levelBoostBps[level];
    }

    /**
     * @notice Returns level of a specific tokenId.
     */
    function getNFTLevel(uint256 tokenId) external view returns (uint256) {
        return cpnft.getTokenLevel(tokenId);
    }

    // =========================================================
    // ADMIN
    // =========================================================

    /**
     * @notice Update boost bps for a specific level (owner only).
     *         Allows future adjustment without contract upgrade.
     */
    function setLevelBoostBps(uint8 level, uint256 boostBps) external onlyOwner {
        require(level < LEVEL_COUNT, "Invalid level");
        _levelBoostBps[level] = boostBps;
    }

    function setCPNFT(address _cpnft) external onlyOwner {
        require(_cpnft != address(0), "Zero address");
        cpnft = ICPNFT(_cpnft);
    }

    function _authorizeUpgrade(address newImpl) internal override onlyOwner {}
}
