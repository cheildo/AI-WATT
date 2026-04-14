import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { AssetRegistry } from "../../typechain-types";

const GPU_CLUSTER  = 0;
const ROBOTICS     = 1;
const AI_ENERGY    = 2;

const PENDING    = 0;
const ACTIVE     = 1;
const FLAGGED    = 2;
const LIQUIDATED = 3;

const LTV_70 = 7000; // 70%
const LTV_80 = 8000; // 80%
const MAX_LTV = 9000; // 90%

function randomAssetId(): string {
  return ethers.keccak256(ethers.randomBytes(32));
}

describe("AssetRegistry", () => {
  let registry: AssetRegistry;

  let admin: HardhatEthersSigner;
  let registrar: HardhatEthersSigner;
  let lendingPool: HardhatEthersSigner;
  let borrower: HardhatEthersSigner;
  let stranger: HardhatEthersSigner;

  const REGISTRAR_ROLE    = ethers.keccak256(ethers.toUtf8Bytes("REGISTRAR_ROLE"));
  const LENDINGPOOL_ROLE  = ethers.keccak256(ethers.toUtf8Bytes("LENDINGPOOL_ROLE"));
  const ADMIN_ROLE        = ethers.keccak256(ethers.toUtf8Bytes("ADMIN_ROLE"));
  const UPGRADER_ROLE     = ethers.keccak256(ethers.toUtf8Bytes("UPGRADER_ROLE"));

  beforeEach(async () => {
    [admin, registrar, lendingPool, borrower, stranger] = await ethers.getSigners();

    const Factory = await ethers.getContractFactory("AssetRegistry");
    registry = (await upgrades.deployProxy(Factory, [admin.address], {
      initializer: "initialize",
      kind: "uups",
    })) as unknown as AssetRegistry;

    await registry.connect(admin).grantRole(REGISTRAR_ROLE, registrar.address);
    await registry.connect(admin).grantRole(LENDINGPOOL_ROLE, lendingPool.address);
  });

  // ── Deployment ─────────────────────────────────────────────────────────────

  describe("deployment", () => {
    it("grants admin roles to admin", async () => {
      expect(await registry.hasRole(ethers.ZeroHash, admin.address)).to.be.true;
      expect(await registry.hasRole(ADMIN_ROLE, admin.address)).to.be.true;
      expect(await registry.hasRole(UPGRADER_ROLE, admin.address)).to.be.true;
    });

    it("MAX_LTV is 9000 (90%)", async () => {
      expect(await registry.MAX_LTV()).to.equal(9000n);
    });

    it("reverts on zero admin address", async () => {
      const Factory = await ethers.getContractFactory("AssetRegistry");
      await expect(
        upgrades.deployProxy(Factory, [ethers.ZeroAddress], {
          initializer: "initialize",
          kind: "uups",
        })
      ).to.be.revertedWithCustomError(registry, "ZeroAddress");
    });
  });

  // ── registerAsset ──────────────────────────────────────────────────────────

  describe("registerAsset", () => {
    it("registers a GPU_CLUSTER asset and emits AssetRegistered", async () => {
      const assetId = randomAssetId();
      await expect(
        registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70)
      )
        .to.emit(registry, "AssetRegistered")
        .withArgs(assetId, GPU_CLUSTER, borrower.address, LTV_70);
    });

    it("stores asset fields correctly", async () => {
      const assetId = randomAssetId();
      await registry.connect(registrar).registerAsset(assetId, ROBOTICS, borrower.address, LTV_80);
      const asset = await registry.getAsset(assetId);
      expect(asset.assetId).to.equal(assetId);
      expect(asset.assetType).to.equal(ROBOTICS);
      expect(asset.borrower).to.equal(borrower.address);
      expect(asset.ltv).to.equal(LTV_80);
      expect(asset.status).to.equal(PENDING);
      expect(asset.registeredAt).to.be.gt(0n);
    });

    it("reverts on duplicate registration", async () => {
      const assetId = randomAssetId();
      await registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70);
      await expect(
        registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70)
      ).to.be.revertedWithCustomError(registry, "AssetAlreadyRegistered");
    });

    it("reverts on zero borrower address", async () => {
      await expect(
        registry.connect(registrar).registerAsset(randomAssetId(), GPU_CLUSTER, ethers.ZeroAddress, LTV_70)
      ).to.be.revertedWithCustomError(registry, "ZeroAddress");
    });

    it("reverts on zero LTV", async () => {
      await expect(
        registry.connect(registrar).registerAsset(randomAssetId(), GPU_CLUSTER, borrower.address, 0)
      ).to.be.revertedWithCustomError(registry, "InvalidLTV");
    });

    it("reverts when LTV exceeds MAX_LTV (9001 bps)", async () => {
      await expect(
        registry.connect(registrar).registerAsset(randomAssetId(), GPU_CLUSTER, borrower.address, 9001)
      ).to.be.revertedWithCustomError(registry, "InvalidLTV");
    });

    it("accepts MAX_LTV exactly", async () => {
      await expect(
        registry.connect(registrar).registerAsset(randomAssetId(), GPU_CLUSTER, borrower.address, MAX_LTV)
      ).to.not.be.reverted;
    });

    it("reverts if caller lacks REGISTRAR_ROLE", async () => {
      await expect(
        registry.connect(stranger).registerAsset(randomAssetId(), GPU_CLUSTER, borrower.address, LTV_70)
      ).to.be.revertedWithCustomError(registry, "AccessControlUnauthorizedAccount");
    });
  });

  // ── updateLTV ──────────────────────────────────────────────────────────────

  describe("updateLTV", () => {
    let assetId: string;

    beforeEach(async () => {
      assetId = randomAssetId();
      await registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70);
    });

    it("updates LTV and emits LTVUpdated", async () => {
      await expect(registry.connect(admin).updateLTV(assetId, LTV_80))
        .to.emit(registry, "LTVUpdated")
        .withArgs(assetId, LTV_70, LTV_80);
      expect((await registry.getAsset(assetId)).ltv).to.equal(LTV_80);
    });

    it("reverts on unregistered assetId", async () => {
      await expect(
        registry.connect(admin).updateLTV(randomAssetId(), LTV_80)
      ).to.be.revertedWithCustomError(registry, "AssetNotFound");
    });

    it("reverts on invalid LTV", async () => {
      await expect(
        registry.connect(admin).updateLTV(assetId, 0)
      ).to.be.revertedWithCustomError(registry, "InvalidLTV");
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        registry.connect(stranger).updateLTV(assetId, LTV_80)
      ).to.be.revertedWithCustomError(registry, "AccessControlUnauthorizedAccount");
    });
  });

  // ── updateStatus ───────────────────────────────────────────────────────────

  describe("updateStatus", () => {
    let assetId: string;

    beforeEach(async () => {
      assetId = randomAssetId();
      await registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70);
    });

    it("ADMIN_ROLE can set status to ACTIVE and emits StatusChanged", async () => {
      await expect(registry.connect(admin).updateStatus(assetId, ACTIVE))
        .to.emit(registry, "StatusChanged")
        .withArgs(assetId, PENDING, ACTIVE);
      expect((await registry.getAsset(assetId)).status).to.equal(ACTIVE);
    });

    it("LENDINGPOOL_ROLE can update status", async () => {
      await registry.connect(admin).updateStatus(assetId, ACTIVE);
      await expect(
        registry.connect(lendingPool).updateStatus(assetId, FLAGGED)
      ).to.not.be.reverted;
      expect((await registry.getAsset(assetId)).status).to.equal(FLAGGED);
    });

    it("LENDINGPOOL_ROLE can mark asset LIQUIDATED", async () => {
      await registry.connect(admin).updateStatus(assetId, ACTIVE);
      await registry.connect(lendingPool).updateStatus(assetId, LIQUIDATED);
      expect((await registry.getAsset(assetId)).status).to.equal(LIQUIDATED);
    });

    it("reverts on unregistered assetId", async () => {
      await expect(
        registry.connect(admin).updateStatus(randomAssetId(), ACTIVE)
      ).to.be.revertedWithCustomError(registry, "AssetNotFound");
    });

    it("reverts if caller lacks both ADMIN_ROLE and LENDINGPOOL_ROLE", async () => {
      await expect(
        registry.connect(stranger).updateStatus(assetId, ACTIVE)
      ).to.be.revertedWithCustomError(registry, "AccessControlUnauthorizedAccount");
    });
  });

  // ── isActive ───────────────────────────────────────────────────────────────

  describe("isActive", () => {
    let assetId: string;

    beforeEach(async () => {
      assetId = randomAssetId();
      await registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70);
    });

    it("returns false when status is PENDING", async () => {
      expect(await registry.isActive(assetId)).to.be.false;
    });

    it("returns true when status is ACTIVE", async () => {
      await registry.connect(admin).updateStatus(assetId, ACTIVE);
      expect(await registry.isActive(assetId)).to.be.true;
    });

    it("returns false when status is FLAGGED", async () => {
      await registry.connect(admin).updateStatus(assetId, ACTIVE);
      await registry.connect(admin).updateStatus(assetId, FLAGGED);
      expect(await registry.isActive(assetId)).to.be.false;
    });

    it("returns false for unregistered assetId", async () => {
      expect(await registry.isActive(randomAssetId())).to.be.false;
    });
  });

  // ── getAsset ───────────────────────────────────────────────────────────────

  describe("getAsset", () => {
    it("reverts for unregistered assetId", async () => {
      await expect(
        registry.getAsset(randomAssetId())
      ).to.be.revertedWithCustomError(registry, "AssetNotFound");
    });
  });

  // ── Upgrade ────────────────────────────────────────────────────────────────

  describe("upgrades", () => {
    it("upgrades preserving all state", async () => {
      const assetId = randomAssetId();
      await registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, LTV_70);
      await registry.connect(admin).updateStatus(assetId, ACTIVE);

      const NewFactory = await ethers.getContractFactory("AssetRegistry", admin);
      const upgraded = await upgrades.upgradeProxy(await registry.getAddress(), NewFactory, { kind: "uups" });
      expect(await upgraded.isActive(assetId)).to.be.true;
    });

    it("reverts upgrade if caller lacks UPGRADER_ROLE", async () => {
      const NewFactory = await ethers.getContractFactory("AssetRegistry", stranger);
      await expect(
        upgrades.upgradeProxy(await registry.getAddress(), NewFactory, { kind: "uups" })
      ).to.be.reverted;
    });
  });
});
