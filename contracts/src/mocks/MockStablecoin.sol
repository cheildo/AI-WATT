// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// @title MockStablecoin
/// @notice Testnet-only ERC-20 stablecoin (USDC / USDT stand-in).
///         Anyone can mint freely — never deploy to mainnet.
/// @dev Deployed on XDC Apothem for Phase 1 integration testing.
contract MockStablecoin is ERC20 {
    uint8 private immutable _decimals;

    /// @notice Fixed amount dispensed per faucet() call (10,000 tokens).
    uint256 public constant FAUCET_AMOUNT = 10_000 * 1e6; // 10k with 6 decimals

    event FaucetDrip(address indexed to, uint256 amount);

    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) ERC20(name_, symbol_) {
        _decimals = decimals_;
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }

    /// @notice Mint any amount to any address. For scripts and test setup.
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    /// @notice Dispense 10,000 tokens to the caller. Callable from block explorer.
    function faucet() external {
        _mint(msg.sender, FAUCET_AMOUNT);
        emit FaucetDrip(msg.sender, FAUCET_AMOUNT);
    }
}
