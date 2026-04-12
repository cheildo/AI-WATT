import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { WattUSD } from "../typechain-types";

describe("WattUSD", () => {
  let wattUSD: WattUSD;
  let admin: SignerWithAddress;
  let minter: SignerWithAddress;
  let user: SignerWithAddress;
  let other: SignerWithAddress;

  const MINTER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const PAUSER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("PAUSER_ROLE"));
  const UPGRADER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("UPGRADER_ROLE"));

  const ONE_WATT = ethers.parseUnits("1", 6);
  const TEN_WATT = ethers.parseUnits("10", 6);

  beforeEach(async () => {
    [admin, minter, user, other] = await ethers.getSigners();

    const WattUSDFactory = await ethers.getContractFactory("WattUSD");
    wattUSD = (await upgrades.deployProxy(WattUSDFactory, [admin.address], {
      kind: "uups",
      initializer: "initialize",
    })) as unknown as WattUSD;
    await wattUSD.waitForDeployment();

    // Grant MINTER_ROLE to the minter signer (simulates MintEngine)
    await wattUSD.connect(admin).grantRole(MINTER_ROLE, minter.address);
  });

  // ── Deployment ────────────────────────────────────────────────────────────

  describe("Deployment", () => {
    it("sets name and symbol", async () => {
      expect(await wattUSD.name()).to.equal("WattUSD");
      expect(await wattUSD.symbol()).to.equal("WATT");
    });

    it("has 6 decimals", async () => {
      expect(await wattUSD.decimals()).to.equal(6);
    });

    it("grants DEFAULT_ADMIN_ROLE to admin", async () => {
      const DEFAULT_ADMIN = ethers.ZeroHash;
      expect(await wattUSD.hasRole(DEFAULT_ADMIN, admin.address)).to.be.true;
    });

    it("grants PAUSER_ROLE and UPGRADER_ROLE to admin", async () => {
      expect(await wattUSD.hasRole(PAUSER_ROLE, admin.address)).to.be.true;
      expect(await wattUSD.hasRole(UPGRADER_ROLE, admin.address)).to.be.true;
    });

    it("does NOT grant MINTER_ROLE to admin on deploy", async () => {
      expect(await wattUSD.hasRole(MINTER_ROLE, admin.address)).to.be.false;
    });

    it("starts with zero supply", async () => {
      expect(await wattUSD.totalSupply()).to.equal(0n);
    });

    it("reverts initialize with zero admin address", async () => {
      const WattUSDFactory = await ethers.getContractFactory("WattUSD");
      await expect(
        upgrades.deployProxy(WattUSDFactory, [ethers.ZeroAddress], {
          kind: "uups",
          initializer: "initialize",
        })
      ).to.be.revertedWithCustomError(wattUSD, "ZeroAddress");
    });
  });

  // ── Minting ───────────────────────────────────────────────────────────────

  describe("mint()", () => {
    it("minter can mint tokens", async () => {
      await wattUSD.connect(minter).mint(user.address, TEN_WATT);
      expect(await wattUSD.balanceOf(user.address)).to.equal(TEN_WATT);
      expect(await wattUSD.totalSupply()).to.equal(TEN_WATT);
    });

    it("emits WattMinted event", async () => {
      await expect(wattUSD.connect(minter).mint(user.address, ONE_WATT))
        .to.emit(wattUSD, "WattMinted")
        .withArgs(user.address, ONE_WATT);
    });

    it("reverts if caller lacks MINTER_ROLE", async () => {
      await expect(
        wattUSD.connect(other).mint(user.address, ONE_WATT)
      ).to.be.revertedWithCustomError(wattUSD, "AccessControlUnauthorizedAccount");
    });

    it("reverts on zero amount", async () => {
      await expect(
        wattUSD.connect(minter).mint(user.address, 0)
      ).to.be.revertedWithCustomError(wattUSD, "ZeroAmount");
    });
  });

  // ── Burning ───────────────────────────────────────────────────────────────

  describe("burn()", () => {
    beforeEach(async () => {
      await wattUSD.connect(minter).mint(user.address, TEN_WATT);
    });

    it("minter can burn tokens from an address", async () => {
      await wattUSD.connect(minter).burn(user.address, ONE_WATT);
      expect(await wattUSD.balanceOf(user.address)).to.equal(TEN_WATT - ONE_WATT);
      expect(await wattUSD.totalSupply()).to.equal(TEN_WATT - ONE_WATT);
    });

    it("emits WattBurned event", async () => {
      await expect(wattUSD.connect(minter).burn(user.address, ONE_WATT))
        .to.emit(wattUSD, "WattBurned")
        .withArgs(user.address, ONE_WATT);
    });

    it("reverts if caller lacks MINTER_ROLE", async () => {
      await expect(
        wattUSD.connect(other).burn(user.address, ONE_WATT)
      ).to.be.revertedWithCustomError(wattUSD, "AccessControlUnauthorizedAccount");
    });

    it("reverts on zero amount", async () => {
      await expect(
        wattUSD.connect(minter).burn(user.address, 0)
      ).to.be.revertedWithCustomError(wattUSD, "ZeroAmount");
    });

    it("reverts if balance is insufficient", async () => {
      await expect(
        wattUSD.connect(minter).burn(user.address, TEN_WATT + ONE_WATT)
      ).to.be.reverted;
    });
  });

  // ── Pause ─────────────────────────────────────────────────────────────────

  describe("pause / unpause", () => {
    it("pauser can pause and unpause", async () => {
      await wattUSD.connect(admin).pause();
      expect(await wattUSD.paused()).to.be.true;
      await wattUSD.connect(admin).unpause();
      expect(await wattUSD.paused()).to.be.false;
    });

    it("blocks transfers when paused", async () => {
      await wattUSD.connect(minter).mint(user.address, TEN_WATT);
      await wattUSD.connect(admin).pause();
      await expect(
        wattUSD.connect(user).transfer(other.address, ONE_WATT)
      ).to.be.revertedWithCustomError(wattUSD, "EnforcedPause");
    });

    it("blocks minting when paused", async () => {
      await wattUSD.connect(admin).pause();
      await expect(
        wattUSD.connect(minter).mint(user.address, ONE_WATT)
      ).to.be.revertedWithCustomError(wattUSD, "EnforcedPause");
    });

    it("reverts pause if caller lacks PAUSER_ROLE", async () => {
      await expect(
        wattUSD.connect(other).pause()
      ).to.be.revertedWithCustomError(wattUSD, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Upgrade ───────────────────────────────────────────────────────────────

  describe("upgradeability", () => {
    it("admin (UPGRADER_ROLE) can upgrade the implementation", async () => {
      const WattUSDV2Factory = await ethers.getContractFactory("WattUSD");
      const upgraded = await upgrades.upgradeProxy(
        await wattUSD.getAddress(),
        WattUSDV2Factory
      );
      // State must be preserved after upgrade
      await wattUSD.connect(minter).mint(user.address, TEN_WATT);
      expect(await upgraded.balanceOf(user.address)).to.equal(TEN_WATT);
    });

    it("reverts upgrade from non-UPGRADER_ROLE account", async () => {
      const WattUSDV2Factory = await ethers.getContractFactory("WattUSD", other);
      await expect(
        upgrades.upgradeProxy(await wattUSD.getAddress(), WattUSDV2Factory)
      ).to.be.revertedWithCustomError(wattUSD, "AccessControlUnauthorizedAccount");
    });
  });
});
