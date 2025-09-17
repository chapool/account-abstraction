// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title Payment
 * @notice Simple payment contract for offline orders
 * @dev Supports ETH and whitelisted ERC20 tokens
 */
contract Payment is ReentrancyGuardTransient {
    struct PaymentInfo {
        uint256 orderId;        // 链下生成的订单ID
        address payer;          // 付款人
        address token;          // 支付代币地址 (ETH为address(0))
        uint256 amount;         // 支付金额
        uint256 timestamp;      // 支付时间
    }
    
    mapping(uint256 => PaymentInfo) public payments;    // 订单ID -> 支付信息
    mapping(address => bool) public allowedTokens;      // 代币白名单
    address public owner;
    
    event PaymentMade(
        uint256 indexed orderId,
        address indexed payer,
        address indexed token,
        uint256 amount,
        uint256 timestamp
    );
    
    event TokenAdded(address indexed token);
    event TokenRemoved(address indexed token);
    event ETHWithdrawn(address indexed owner, uint256 amount);
    event TokenWithdrawn(address indexed owner, address indexed token, uint256 amount);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event RefundProcessed(
        uint256 indexed orderId,
        address indexed user,
        address indexed token,
        uint256 amount
    );
    
    error OnlyOwner();
    error TokenNotAllowed();
    error OrderAlreadyPaid();
    error InvalidAmount();
    error InvalidToken();
    error InsufficientBalance();
    error TransferFailed();
    error InvalidOwner();
    error ArrayLengthMismatch();
    error InvalidUser();
    error RefundFailed();
    
    modifier onlyOwner() {
        if (msg.sender != owner) revert OnlyOwner();
        _;
    }
    
    constructor() {
        owner = msg.sender;
        // ETH is allowed by default
        allowedTokens[address(0)] = true;
        emit TokenAdded(address(0));
    }
    
    /**
     * @notice Pay for an offline order
     * @param orderId Order ID generated off-chain
     * @param token Token address (use address(0) for ETH)
     * @param amount Amount to pay
     */
    function pay(uint256 orderId, address token, uint256 amount) 
        external 
        payable 
        nonReentrant 
    {
        if (!allowedTokens[token]) revert TokenNotAllowed();
        if (payments[orderId].payer != address(0)) revert OrderAlreadyPaid();
        if (amount == 0) revert InvalidAmount();
        
        if (token == address(0)) {
            // ETH payment
            if (msg.value != amount) revert InvalidAmount();
        } else {
            // ERC20 payment
            if (msg.value != 0) revert InvalidAmount();
            IERC20(token).transferFrom(msg.sender, address(this), amount);
        }
        
        payments[orderId] = PaymentInfo({
            orderId: orderId,
            payer: msg.sender,
            token: token,
            amount: amount,
            timestamp: block.timestamp
        });
        
        emit PaymentMade(orderId, msg.sender, token, amount, block.timestamp);
    }
    
    /**
     * @notice Add token to whitelist
     * @param token Token address to add (use address(0) for ETH)
     */
    function addAllowedToken(address token) external onlyOwner {
        allowedTokens[token] = true;
        emit TokenAdded(token);
    }
    
    /**
     * @notice Remove token from whitelist
     * @param token Token address to remove
     */
    function removeAllowedToken(address token) external onlyOwner {
        if (token == address(0)) revert InvalidToken(); // Cannot remove ETH
        allowedTokens[token] = false;
        emit TokenRemoved(token);
    }
    
    /**
     * @notice Withdraw ETH from contract
     * @param amount Amount to withdraw
     */
    function withdrawETH(uint256 amount) external onlyOwner nonReentrant {
        if (amount > address(this).balance) revert InsufficientBalance();
        
        (bool success, ) = payable(owner).call{value: amount}("");
        if (!success) revert TransferFailed();
        
        emit ETHWithdrawn(owner, amount);
    }
    
    /**
     * @notice Withdraw all ETH from contract
     */
    function withdrawAllETH() external onlyOwner nonReentrant {
        uint256 balance = address(this).balance;
        if (balance == 0) revert InsufficientBalance();
        
        (bool success, ) = payable(owner).call{value: balance}("");
        if (!success) revert TransferFailed();
        
        emit ETHWithdrawn(owner, balance);
    }
    
    /**
     * @notice Withdraw ERC20 tokens from contract
     * @param token Token address
     * @param amount Amount to withdraw
     */
    function withdrawToken(address token, uint256 amount) external onlyOwner nonReentrant {
        if (token == address(0)) revert InvalidToken();
        
        IERC20 tokenContract = IERC20(token);
        if (amount > tokenContract.balanceOf(address(this))) revert InsufficientBalance();
        
        tokenContract.transfer(owner, amount);
        emit TokenWithdrawn(owner, token, amount);
    }
    
    /**
     * @notice Withdraw all tokens of a specific type
     * @param token Token address
     */
    function withdrawAllTokens(address token) external onlyOwner nonReentrant {
        if (token == address(0)) revert InvalidToken();
        
        IERC20 tokenContract = IERC20(token);
        uint256 balance = tokenContract.balanceOf(address(this));
        if (balance == 0) revert InsufficientBalance();
        
        tokenContract.transfer(owner, balance);
        emit TokenWithdrawn(owner, token, balance);
    }
    
    /**
     * @notice Transfer ownership to a new owner
     * @param newOwner New owner address
     */
    function transferOwnership(address newOwner) external onlyOwner {
        if (newOwner == address(0)) revert InvalidOwner();
        
        address previousOwner = owner;
        owner = newOwner;
        emit OwnershipTransferred(previousOwner, newOwner);
    }
    
    /**
     * @notice Batch refund for multiple orders
     * @param orderIds Array of order IDs
     * @param users Array of user addresses to refund
     * @param amounts Array of amounts to refund
     * @param tokens Array of token addresses (use address(0) for ETH)
     */
    function batchRefund(
        uint256[] calldata orderIds,
        address[] calldata users,
        uint256[] calldata amounts,
        address[] calldata tokens
    ) external onlyOwner nonReentrant {
        // Validate array lengths
        if (orderIds.length != users.length || 
            users.length != amounts.length || 
            amounts.length != tokens.length) {
            revert ArrayLengthMismatch();
        }
        
        if (orderIds.length == 0) revert InvalidAmount();
        
        // Process each refund
        for (uint256 i = 0; i < orderIds.length; i++) {
            _processRefund(orderIds[i], users[i], amounts[i], tokens[i]);
        }
    }
    
    /**
     * @notice Internal function to process individual refund
     * @param orderId Order ID (for tracking purposes only)
     * @param user User address to refund
     * @param amount Amount to refund
     * @param token Token address (use address(0) for ETH)
     */
    function _processRefund(
        uint256 orderId,
        address user,
        uint256 amount,
        address token
    ) internal {
        // Validate inputs
        if (user == address(0)) revert InvalidUser();
        if (amount == 0) revert InvalidAmount();
        if (!allowedTokens[token]) revert TokenNotAllowed();
        
        // Process refund based on token type
        if (token == address(0)) {
            // ETH refund
            if (amount > address(this).balance) revert InsufficientBalance();
            
            (bool success, ) = payable(user).call{value: amount}("");
            if (!success) revert RefundFailed();
        } else {
            // ERC20 refund
            IERC20 tokenContract = IERC20(token);
            if (amount > tokenContract.balanceOf(address(this))) revert InsufficientBalance();
            
            bool success = tokenContract.transfer(user, amount);
            if (!success) revert RefundFailed();
        }
        
        // Emit refund event
        emit RefundProcessed(orderId, user, token, amount);
    }
    
    /**
     * @notice Get payment information for an order
     * @param orderId Order ID
     * @return PaymentInfo struct
     */
    function getPayment(uint256 orderId) external view returns (PaymentInfo memory) {
        return payments[orderId];
    }
    
    /**
     * @notice Check if token is allowed for payments
     * @param token Token address
     * @return bool True if allowed
     */
    function isTokenAllowed(address token) external view returns (bool) {
        return allowedTokens[token];
    }
    
    /**
     * @notice Check if order has been paid
     * @param orderId Order ID
     * @return bool True if paid
     */
    function isOrderPaid(uint256 orderId) external view returns (bool) {
        return payments[orderId].payer != address(0);
    }
    
    /**
     * @notice Get contract ETH balance
     * @return uint256 ETH balance
     */
    function getETHBalance() external view returns (uint256) {
        return address(this).balance;
    }
    
    /**
     * @notice Get contract token balance
     * @param token Token address
     * @return uint256 Token balance
     */
    function getTokenBalance(address token) external view returns (uint256) {
        if (token == address(0)) return address(this).balance;
        return IERC20(token).balanceOf(address(this));
    }
}