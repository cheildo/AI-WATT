import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { WattUSD, MintEngine, MockERC20 } from "../typechain-types";

describe("MintEngine", () => {
  let wattUSD: WattUSD;
  let mintEngine: MintEngine;
  let mockUSDC: MockERC20;
  let mockUSDT: MockERC20;

  let admin: SignerWithAddress;
  let treasury: SignerWithAddress;
  let user: SignerWithAddress;
  let other: SignerWithAddress;

  const MINTER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const ADMIN_ROLE = ethers.keccak256(ethers.toUtf8Bytes("ADMIN_ROLE"));

  // 1,000 USDC (6 decimals)
  const ONE_THOUSAND = ethers.parseUnits("1000", 6);
  // Expected fee: 0.1% of 1000 = 1 USDC
  const EXPECTED_FEE = ethers.parseUnits("1", 6);
  // Expected WATT minted: 1000 - 1 = 999
  const EXPECTED_WATT = ethers.parseUnits("999", 6);

  beforeEach(async () => {
    [admin, treasury, user, other] = await ethers.getSigners();

    // Deploy mock stablecoins
    const MockERC20Factory = await ethers.getContractFactory("MockERC20");
    mockUSDC = (await MockERC20Factory.deploy("USD Coin", "USDC", 6)) as unknown as MockERC20;
    mockUSDT = (await MockERC20Factory.deploy("Tether USD", "USDT", 6)) as unknown as MockERC20;

    // Deploy WattUSD proxy
    const WattUSDFactory = await ethers.getContractFactory("WattUSD");
    wattUSD = (await upgrades.deployProxy(WattUSDFactory, [admin.address], {
      kind: "uups",
      initializer: "initialize",
    })) as unknown as WattUSD;
    await wattUSD.waitForDeployment();

    // Deploy MintEngine proxy
    const MintEngineFactory = await ethers.getContractFactory("MintEngine");
    mintEngine = (await upgrades.deployProxy(
      MintEngineFactory,
      [admin.address, await wattUSD.getAddress(), treasury.address],
      { kind: "uups", initializer: "initialize" }
    )) as unknown as MintEngine;
    await mintEngine.waitForDeployment();

    // Grant MINTER_ROLE on WattUSD to MintEngine
    await wattUSD.connect(admin).grantRole(MINTER_ROLE, await mintEngine.getAddress());

    // Accept USDC in MintEngine
    await mintEngine.connect(admin).setAcceptedStablecoin(await mockUSDC.getAddress(), true);
    await mintEngine.connect(admin).setAcceptedStablecoin(await mockUSDT.getAddress(), true);

    // Fund user with stablecoins
    await mockUSDC.mint(user.address, ONE_THOUSAND * 10n);
    await mockUSDT.mint(user.address, ONE_THOUSAND * 10n);

    // Approve MintEngine to spend user's stablecoins
    await mockUSDC.connect(user).approve(await mintEngine.getAddress(), ethers.MaxUint256);
    await mockUSDT.connect(user).approve(await mintEngine.getAddress(), ethers.MaxUint256);
  });

  // ── Deployment ────────────────────────────────────────────────────────────

  describe("Deployment", () => {
    it("sets wattUSD, treasury correctly", async () => {
      expect(await mintEngine.wattUSD()).to.equal(await wattUSD.getAddress());
      expect(await mintEngine.treasury()).to.equal(treasury.address);
    });

    it("has FEE_BPS of 10 (0.1%)", async () => {
      expect(await mintEngine.FEE_BPS()).to.equal(10n);
    });

    it("accepts USDC and USDT after setup", async () => {
      expect(await mintEngine.isAcceptedStablecoin(await mockUSDC.getAddress())).to.be.true;
      expect(await mintEngine.isAcceptedStablecoin(await mockUSDT.getAddress())).to.be.true;
    });

    it("reverts initialize with zero address", async () => {
      const MintEngineFactory = await ethers.getContractFactory("MintEngine");
      await expect(
        upgrades.deployProxy(
          MintEngineFactory,
          [ethers.ZeroAddress, await wattUSD.getAddress(), treasury.address],
          { kind: "uups", initializer: "initialize" }
        )
      ).to.be.revertedWithCustomError(mintEngine, "ZeroAddress");
    });
  });

  // ── Minting ───────────────────────────────────────────────────────────────

  describe("mint()", () => {
    it("mints WATT 1:1 minus 0.1% fee", async () => {
      await mintEngine.connect(user).mint(await mockUSDC.getAddress(), ONE_THOUSAND);

      expect(await wattUSD.balanceOf(user.address)).to.equal(EXPECTED_WATT);
    });

    it("sends 0.1% fee in stablecoin to treasury", async () => {
      const beforeBalance = await mockUSDC.balanceOf(treasury.address);
      await mintEngine.connect(user).mint(await mockUSDC.getAddress(), ONE_THOUSAND);
      expect(await mockUSDC.balanceOf(treasury.address)).to.equal(beforeBalance + EXPECTED_FEE);
    });

    it("holds net collateral in MintEngine", async () => {
      await mintEngine.connect(user).mint(await mockUSDC.getAddress(), ONE_THOUSAND);
      expect(await mintEngine.collateralBalance(await mockUSDC.getAddress())).to.equal(EXPECTED_WATT);
    });

    it("emits Minted event with correct args", async () => {
      await expect(
        mintEngine.connect(user).mint(await mockUSDC.getAddress(), ONE_THOUSAND)
      )
        .to.emit(mintEngine, "Minted")
        .withArgs(
          user.address,
          await mockUSDC.getAddress(),
          ONE_THOUSAND,
          EXPECTED_WATT,
          EXPECTED_FEE
        );
    });

    it("works with USDT as well", async () => {
      await mintEngine.connect(user).mint(await mockUSDT.getAddress(), ONE_THOUSAND);
      expect(await wattUSD.balanceOf(user.address)).to.equal(EXPECTED_WATT);
    });

    it("reverts on zero amount", async () => {
      await expect(
        mintEngine.connect(user).mint(await mockUSDC.getAddress(), 0)
      ).to.be.revertedWithCustomError(mintEngine, "ZeroAmount");
    });

    it("reverts on unaccepted stablecoin", async () => {
      const MockERC20Factory = await ethers.getContractFactory("MockERC20");
      const rogue = await MockERC20Factory.deploy("Rogue", "RGE", 6);
      await expect(
        mintEngine.connect(user).mint(await rogue.getAddress(), ONE_THOUSAND)
      ).to.be.revertedWithCustomError(mintEngine, "StablecoinNotAccepted");
    });

    it("reverts when paused", async () => {
      await mintEngine.connect(admin).pause();
      await expect(
        mintEngine.connect(user).mint(await mockUSDC.getAddress(), ONE_THOUSAND)
      ).to.be.revertedWithCustomError(mintEngine, "EnforcedPause");
    });
  });

  // ── Redemption ────────────────────────────────────────────────────────────

  describe("redeem()", () => {
    beforeEach(async () => {
      // User mints first so they have WATT and MintEngine has collateral
      await mintEngine.connect(user).mint(await mockUSDC.getAddress(), ONE_THOUSAND);
    });

    it("burns WATT and returns stablecoin minus 0.1% fee", async () => {
      const wattBalance = await wattUSD.balanceOf(user.address); // 999 WATT
      const fee = (wattBalance * 10n) / 10_000n;
      const expectedReturn = wattBalance - fee;

      const usdcBefore = await mockUSDC.balanceOf(user.address);
      await mintEngine.connect(user).redeem(await mockUSDC.getAddress(), wattBalance);

      expect(await wattUSD.balanceOf(user.address)).to.equal(0n);
      expect(await mockUSDC.balanceOf(user.address)).to.equal(usdcBefore + expectedReturn);
    });

    it("sends redemption fee to treasury", async () => {
      const wattBalance = await wattUSD.balanceOf(user.address);
      const fee = (wattBalance * 10n) / 10_000n;
      const treasuryBefore = await mockUSDC.balanceOf(treasury.address);

      await mintEngine.connect(user).redeem(await mockUSDC.getAddress(), wattBalance);

      expect(await mockUSDC.balanceOf(treasury.address)).to.equal(treasuryBefore + fee);
    });

    it("emits Redeemed event", async () => {
      const wattBalance = await wattUSD.balanceOf(user.address);
      const fee = (wattBalance * 10n) / 10_000n;
      const stablecoinReturned = wattBalance - fee;

      await expect(
        mintEngine.connect(user).redeem(await mockUSDC.getAddress(), wattBalance)
      )
        .to.emit(mintEngine, "Redeemed")
        .withArgs(
          user.address,
          await mockUSDC.getAddress(),
          wattBalance,
          stablecoinReturned,
          fee
        );
    });

    it("reverts on zero amount", async () => {
      await expect(
        mintEngine.connect(user).redeem(await mockUSDC.getAddress(), 0)
      ).to.be.revertedWithCustomError(mintEngine, "ZeroAmount");
    });

    it("reverts on unaccepted stablecoin", async () => {
      const MockERC20Factory = await ethers.getContractFactory("MockERC20");
      const rogue = await MockERC20Factory.deploy("Rogue", "RGE", 6);
      await expect(
        mintEngine.connect(user).redeem(await rogue.getAddress(), ONE_THOUSAND)
      ).to.be.revertedWithCustomError(mintEngine, "StablecoinNotAccepted");
    });

    it("reverts when paused", async () => {
      await mintEngine.connect(admin).pause();
      await expect(
        mintEngine.connect(user).redeem(await mockUSDC.getAddress(), ONE_THOUSAND)
      ).to.be.revertedWithCustomError(mintEngine, "EnforcedPause");
    });

    it("reverts on insufficient collateral", async () => {
      // Try to redeem more WATT than there is collateral for
      await wattUSD.connect(admin).grantRole(
        ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE")),
        admin.address
      );
      // Mint extra WATT without depositing collateral
      await wattUSD.connect(admin).mint(user.address, ONE_THOUSAND * 100n);

      await expect(
        mintEngine.connect(user).redeem(await mockUSDC.getAddress(), ONE_THOUSAND * 100n)
      ).to.be.revertedWithCustomError(mintEngine, "InsufficientCollateral");
    });
  });

  // ── Admin ─────────────────────────────────────────────────────────────────

  describe("setAcceptedStablecoin()", () => {
    it("admin can remove an accepted stablecoin", async () => {
      await mintEngine.connect(admin).setAcceptedStablecoin(await mockUSDC.getAddress(), false);
      expect(await mintEngine.isAcceptedStablecoin(await mockUSDC.getAddress())).to.be.false;
    });

    it("emits StablecoinUpdated event", async () => {
      await expect(
        mintEngine.connect(admin).setAcceptedStablecoin(await mockUSDC.getAddress(), false)
      )
        .to.emit(mintEngine, "StablecoinUpdated")
        .withArgs(await mockUSDC.getAddress(), false);
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        mintEngine.connect(other).setAcceptedStablecoin(await mockUSDC.getAddress(), false)
      ).to.be.revertedWithCustomError(mintEngine, "AccessControlUnauthorizedAccount");
    });

    it("reverts on zero address", async () => {
      await expect(
        mintEngine.connect(admin).setAcceptedStablecoin(ethers.ZeroAddress, true)
      ).to.be.revertedWithCustomError(mintEngine, "ZeroAddress");
    });
  });

  describe("setTreasury()", () => {
    it("admin can update treasury", async () => {
      await mintEngine.connect(admin).setTreasury(other.address);
      expect(await mintEngine.treasury()).to.equal(other.address);
    });

    it("emits TreasuryUpdated event", async () => {
      await expect(mintEngine.connect(admin).setTreasury(other.address))
        .to.emit(mintEngine, "TreasuryUpdated")
        .withArgs(treasury.address, other.address);
    });

    it("reverts on zero address", async () => {
      await expect(
        mintEngine.connect(admin).setTreasury(ethers.ZeroAddress)
      ).to.be.revertedWithCustomError(mintEngine, "ZeroAddress");
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        mintEngine.connect(other).setTreasury(other.address)
      ).to.be.revertedWithCustomError(mintEngine, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Upgrade ───────────────────────────────────────────────────────────────

  describe("upgradeability", () => {
    it("admin can upgrade the implementation", async () => {
      const MintEngineV2Factory = await ethers.getContractFactory("MintEngine");
      const upgraded = await upgrades.upgradeProxy(
        await mintEngine.getAddress(),
        MintEngineV2Factory
      );
      // State preserved
      expect(await upgraded.treasury()).to.equal(treasury.address);
    });

    it("reverts upgrade from non-UPGRADER_ROLE", async () => {
      const MintEngineV2Factory = await ethers.getContractFactory("MintEngine", other);
      await expect(
        upgrades.upgradeProxy(await mintEngine.getAddress(), MintEngineV2Factory)
      ).to.be.revertedWithCustomError(mintEngine, "AccessControlUnauthorizedAccount");
    });
  });
});
