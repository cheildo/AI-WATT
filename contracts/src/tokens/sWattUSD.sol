// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ERC4626Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC4626Upgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PausableUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {IsWattUSD} from "../interfaces/IsWattUSD.sol";

/// @title sWattUSD
/// @notice ERC-4626 tokenized yield vault. Underlying asset: WATT (WattUSD).
///
///         Depositors stake WATT → receive sWATT shares.
///         As GPU loan repayments flow in via receiveYield(), totalAssets grows,
///         making each sWATT share redeemable for more WATT over time.
///
///         Large redemptions (above wevThreshold) are routed through WEVQueue
///         once that contract is deployed in Phase 5. Until then, wevQueue is
///         address(0) and all redemptions are immediate.
///
/// @dev UUPS upgradeable. Storage layout — only append new variables.
///      Inflation attack mitigation: deploy script seeds an initial deposit.
contract sWattUSD is
    Initializable,
    ERC4626Upgradeable,
    ERC20PausableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    IsWattUSD
{
    using SafeERC20 for IERC20;

    // ── Roles ─────────────────────────────────────────────────────────────────

    /// @notice Granted to LendingPool and TreasuryService — allows calling receiveYield().
    bytes32 public constant YIELD_DISTRIBUTOR_ROLE = keccak256("YIELD_DISTRIBUTOR_ROLE");

    /// @notice Allows pausing/unpausing the vault.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Allows updating WEVQueue address, threshold, and protocol parameters.
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Allows authorizing contract upgrades.
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    // ── Storage ───────────────────────────────────────────────────────────────

    /// @notice WEVQueue contract address. address(0) = no queue enforced yet.
    address public wevQueue;

    /// @notice Redemption size (in WATT, 6 decimals) above which WEVQueue is required.
    ///         Default: 100,000 WATT.
    uint256 public wevThreshold;

    // ── Constructor ───────────────────────────────────────────────────────────

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ── Initializer ───────────────────────────────────────────────────────────

    /// @notice Initializes the proxy. Called once by the deploy script.
    /// @param admin   Address granted admin roles. Should be a multisig in production.
    /// @param watt    Address of the WattUSD (WATT) token — the underlying ERC-4626 asset.
    function initialize(address admin, address watt) external initializer {
        if (admin == address(0) || watt == address(0)) revert ZeroAddress();

        __ERC20_init("Staked WattUSD", "sWATT");
        __ERC20Pausable_init();
        __ERC4626_init(IERC20(watt));
        __AccessControl_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(PAUSER_ROLE, admin);
        _grantRole(UPGRADER_ROLE, admin);

        wevThreshold = 100_000 * 1e6; // 100,000 WATT default
    }

    // ── IsWattUSD — yield ─────────────────────────────────────────────────────

    /// @inheritdoc IsWattUSD
    function receiveYield(uint256 amount) external onlyRole(YIELD_DISTRIBUTOR_ROLE) {
        if (amount == 0) revert ZeroAmount();
        IERC20(asset()).safeTransferFrom(msg.sender, address(this), amount);
        emit YieldReceived(amount, totalAssets());
    }

    // ── IsWattUSD — admin ─────────────────────────────────────────────────────

    /// @inheritdoc IsWattUSD
    function setWEVQueue(address newWEVQueue) external onlyRole(ADMIN_ROLE) {
        address old = wevQueue;
        wevQueue = newWEVQueue;
        emit WEVQueueUpdated(old, newWEVQueue);
    }

    /// @inheritdoc IsWattUSD
    function setWEVThreshold(uint256 newThreshold) external onlyRole(ADMIN_ROLE) {
        if (newThreshold == 0) revert ZeroAmount();
        uint256 old = wevThreshold;
        wevThreshold = newThreshold;
        emit WEVThresholdUpdated(old, newThreshold);
    }

    /// @inheritdoc IsWattUSD
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @inheritdoc IsWattUSD
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    // ── ERC-20 overrides ──────────────────────────────────────────────────────

    /// @dev Resolves diamond inheritance: both ERC20Upgradeable and ERC4626Upgradeable
    ///      declare decimals(). ERC4626 returns asset decimals; we keep that behavior.
    function decimals()
        public
        view
        override(ERC20Upgradeable, ERC4626Upgradeable)
        returns (uint8)
    {
        return super.decimals();
    }

    // ── IsWattUSD — view ──────────────────────────────────────────────────────

    /// @inheritdoc IsWattUSD
    function navPerShare() external view returns (uint256) {
        return convertToAssets(10 ** decimals());
    }

    // ── ERC-4626 overrides ────────────────────────────────────────────────────

    /// @notice Blocks withdrawals above wevThreshold when WEVQueue is configured.
    function maxWithdraw(address owner) public view override returns (uint256) {
        if (paused()) return 0;
        uint256 ownerMax = super.maxWithdraw(owner);
        if (wevQueue != address(0) && ownerMax > wevThreshold) return wevThreshold;
        return ownerMax;
    }

    /// @notice Blocks redemptions above the share equivalent of wevThreshold when WEVQueue is set.
    function maxRedeem(address owner) public view override returns (uint256) {
        if (paused()) return 0;
        uint256 ownerMax = super.maxRedeem(owner);
        if (wevQueue != address(0)) {
            uint256 thresholdShares = convertToShares(wevThreshold);
            if (ownerMax > thresholdShares) return thresholdShares;
        }
        return ownerMax;
    }

    /// @dev Reverts with a descriptive error when a large withdrawal is attempted
    ///      and the WEVQueue guard is active.
    function _withdraw(
        address caller,
        address receiver,
        address owner,
        uint256 assets,
        uint256 shares
    ) internal override {
        if (wevQueue != address(0) && assets > wevThreshold) {
            revert LargeRedemptionUseWEVQueue(assets, wevThreshold, wevQueue);
        }
        super._withdraw(caller, receiver, owner, assets, shares);
    }

    // ── Internal ──────────────────────────────────────────────────────────────

    /// @dev Resolves diamond inheritance: ERC4626 → ERC20 and ERC20Pausable both
    ///      define _update. super chain applies pause check before any transfer.
    function _update(
        address from,
        address to,
        uint256 value
    ) internal override(ERC20Upgradeable, ERC20PausableUpgradeable) {
        super._update(from, to, value);
    }

    /// @dev Only UPGRADER_ROLE can authorize an upgrade.
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADER_ROLE)
    {}
}
