import { ethers } from "hardhat";

export interface PackedUserOperation {
    sender: string;
    nonce: number;
    initCode: string;
    callData: string;
    accountGasLimits: string; // packed: verificationGasLimit + callGasLimit
    preVerificationGas: number;
    gasFees: string; // packed: maxPriorityFeePerGas + maxFeePerGas
    paymasterAndData: string;
    signature: string;
}

export function packAccountGasLimits(verificationGasLimit: number, callGasLimit: number): string {
    return ethers.concat([
        ethers.toBeHex(verificationGasLimit, 16),
        ethers.toBeHex(callGasLimit, 16)
    ]);
}

export function packGasFees(maxPriorityFeePerGas: number, maxFeePerGas: number): string {
    return ethers.concat([
        ethers.toBeHex(maxPriorityFeePerGas, 16),
        ethers.toBeHex(maxFeePerGas, 16)
    ]);
}

export function createUserOp(params: Partial<PackedUserOperation>): PackedUserOperation {
    return {
        sender: params.sender || ethers.ZeroAddress,
        nonce: params.nonce || 0,
        initCode: params.initCode || "0x",
        callData: params.callData || "0x",
        accountGasLimits: params.accountGasLimits || packAccountGasLimits(21000, 100000),
        preVerificationGas: params.preVerificationGas || 21000,
        gasFees: params.gasFees || packGasFees(
            Number(ethers.parseUnits("20", "gwei")),
            Number(ethers.parseUnits("30", "gwei"))
        ),
        paymasterAndData: params.paymasterAndData || "0x",
        signature: params.signature || "0x"
    };
}

export function getUserOpHash(userOp: PackedUserOperation, entryPointAddress: string, chainId: number): string {
    return ethers.keccak256(
        ethers.AbiCoder.defaultAbiCoder().encode(
            [
                "address", // sender
                "uint256", // nonce
                "bytes32", // initCode hash
                "bytes32", // callData hash
                "bytes32", // accountGasLimits
                "uint256", // preVerificationGas
                "bytes32", // gasFees
                "bytes32", // paymasterAndData hash
                "address", // entryPoint
                "uint256"  // chainId
            ],
            [
                userOp.sender,
                userOp.nonce,
                ethers.keccak256(userOp.initCode),
                ethers.keccak256(userOp.callData),
                userOp.accountGasLimits,
                userOp.preVerificationGas,
                userOp.gasFees,
                ethers.keccak256(userOp.paymasterAndData),
                entryPointAddress,
                chainId
            ]
        )
    );
}