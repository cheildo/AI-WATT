import { expect } from "chai";
import { ethers, upgrades, network } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { HealthAttestation } from "../../typechain-types";

const COOLDOWN = 12 * 60 * 60; // 12 hours in seconds

function randomAssetId(): string {
  return ethers.keccak256(ethers.randomBytes(32));
}

function randomHash(): string {
  return ethers.keccak256(ethers.randomBytes(32));
}

async function timeTravel(seconds: number): Promise<void> {
  await network.provider.send("evm_increaseTime", [seconds]);
  await network.provider.send("evm_mine");
}

describe("HealthAttestation", () => {
  let attestation: HealthAttestation;

  let admin: HardhatEthersSigner;
  let veriflowSigner: HardhatEthersSigner;
  let stranger: HardhatEthersSigner;

  const VERIFLOW_SIGNER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("VERIFLOW_SIGNER"));
  const UPGRADER_ROLE        = ethers.keccak256(ethers.toUtf8Bytes("UPGRADER_ROLE"));

  beforeEach(async () => {
    [admin, veriflowSigner, stranger] = await ethers.getSigners();

    const Factory = await ethers.getContractFactory("HealthAttestation");
    attestation = (await upgrades.deployProxy(Factory, [admin.address], {
      initializer: "initialize",
      kind: "uups",
    })) as unknown as HealthAttestation;

    await attestation.connect(admin).grantRole(VERIFLOW_SIGNER_ROLE, veriflowSigner.address);
  });

  // ── Deployment ─────────────────────────────────────────────────────────────

  describe("deployment", () => {
    it("sets COOLDOWN constant to 12 hours", async () => {
      expect(await attestation.COOLDOWN()).to.equal(COOLDOWN);
    });

    it("grants admin roles to admin", async () => {
      expect(await attestation.hasRole(ethers.ZeroHash, admin.address)).to.be.true;
      expect(await attestation.hasRole(UPGRADER_ROLE, admin.address)).to.be.true;
    });

    it("reverts on zero admin address", async () => {
      const Factory = await ethers.getContractFactory("HealthAttestation");
      await expect(
        upgrades.deployProxy(Factory, [ethers.ZeroAddress], {
          initializer: "initialize",
          kind: "uups",
        })
      ).to.be.revertedWithCustomError(attestation, "ZeroAddress");
    });
  });

  // ── submitAttestation ──────────────────────────────────────────────────────

  describe("submitAttestation", () => {
    it("accepts a valid attestation and emits AttestationSubmitted", async () => {
      const assetId    = randomAssetId();
      const healthHash = randomHash();
      const score      = 85;

      const tx = attestation.connect(veriflowSigner).submitAttestation(assetId, healthHash, score);
      await expect(tx).to.emit(attestation, "AttestationSubmitted");
    });

    it("stores the attestation as latest", async () => {
      const assetId    = randomAssetId();
      const healthHash = randomHash();
      const score      = 72;

      await attestation.connect(veriflowSigner).submitAttestation(assetId, healthHash, score);
      const latest = await attestation.getLatestAttestation(assetId);

      expect(latest.assetId).to.equal(assetId);
      expect(latest.healthHash).to.equal(healthHash);
      expect(latest.score).to.equal(score);
      expect(latest.timestamp).to.be.gt(0n);
    });

    it("accepts score = 0 (asset is down)", async () => {
      await expect(
        attestation.connect(veriflowSigner).submitAttestation(randomAssetId(), randomHash(), 0)
      ).to.not.be.reverted;
    });

    it("accepts score = 100 (perfect health)", async () => {
      await expect(
        attestation.connect(veriflowSigner).submitAttestation(randomAssetId(), randomHash(), 100)
      ).to.not.be.reverted;
    });

    it("reverts on score > 100", async () => {
      await expect(
        attestation.connect(veriflowSigner).submitAttestation(randomAssetId(), randomHash(), 101)
      ).to.be.revertedWithCustomError(attestation, "InvalidScore");
    });

    it("reverts if caller lacks VERIFLOW_SIGNER role", async () => {
      await expect(
        attestation.connect(stranger).submitAttestation(randomAssetId(), randomHash(), 80)
      ).to.be.revertedWithCustomError(attestation, "AccessControlUnauthorizedAccount");
    });

    it("reverts when attested again within cooldown period", async () => {
      const assetId = randomAssetId();
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 80);
      await expect(
        attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 85)
      ).to.be.revertedWithCustomError(attestation, "AttestationTooSoon");
    });

    it("accepts a second attestation after cooldown has passed", async () => {
      const assetId = randomAssetId();
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 80);
      await timeTravel(COOLDOWN + 1);
      await expect(
        attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 90)
      ).to.not.be.reverted;
    });

    it("separate assetIds are independent — no cross-cooldown", async () => {
      const asset1 = randomAssetId();
      const asset2 = randomAssetId();
      await attestation.connect(veriflowSigner).submitAttestation(asset1, randomHash(), 80);
      // asset2 should not be affected by asset1's cooldown
      await expect(
        attestation.connect(veriflowSigner).submitAttestation(asset2, randomHash(), 75)
      ).to.not.be.reverted;
    });
  });

  // ── hasAttestation ─────────────────────────────────────────────────────────

  describe("hasAttestation", () => {
    it("returns false before any attestation", async () => {
      expect(await attestation.hasAttestation(randomAssetId())).to.be.false;
    });

    it("returns true after first attestation", async () => {
      const assetId = randomAssetId();
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 80);
      expect(await attestation.hasAttestation(assetId)).to.be.true;
    });
  });

  // ── getLatestAttestation ───────────────────────────────────────────────────

  describe("getLatestAttestation", () => {
    it("returns zero struct for asset with no attestation", async () => {
      const latest = await attestation.getLatestAttestation(randomAssetId());
      expect(latest.score).to.equal(0);
      expect(latest.timestamp).to.equal(0);
    });

    it("reflects the most recent attestation after multiple submissions", async () => {
      const assetId = randomAssetId();
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 70);
      await timeTravel(COOLDOWN + 1);
      const newHash = randomHash();
      await attestation.connect(veriflowSigner).submitAttestation(assetId, newHash, 90);

      const latest = await attestation.getLatestAttestation(assetId);
      expect(latest.score).to.equal(90);
      expect(latest.healthHash).to.equal(newHash);
    });
  });

  // ── getAttestationHistory ──────────────────────────────────────────────────

  describe("getAttestationHistory", () => {
    let assetId: string;

    beforeEach(async () => {
      assetId = randomAssetId();
      // Submit 3 attestations spread over cooldown periods
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 60);
      await timeTravel(COOLDOWN + 1);
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 75);
      await timeTravel(COOLDOWN + 1);
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 90);
    });

    it("returns all history when limit=0", async () => {
      const history = await attestation.getAttestationHistory(assetId, 0);
      expect(history.length).to.equal(3);
    });

    it("returns newest-first ordering", async () => {
      const history = await attestation.getAttestationHistory(assetId, 0);
      expect(history[0].score).to.equal(90); // newest
      expect(history[1].score).to.equal(75);
      expect(history[2].score).to.equal(60); // oldest
    });

    it("respects limit parameter", async () => {
      const history = await attestation.getAttestationHistory(assetId, 2);
      expect(history.length).to.equal(2);
      expect(history[0].score).to.equal(90); // newest two
      expect(history[1].score).to.equal(75);
    });

    it("returns all when limit exceeds history length", async () => {
      const history = await attestation.getAttestationHistory(assetId, 100);
      expect(history.length).to.equal(3);
    });

    it("returns empty array for asset with no attestations", async () => {
      const history = await attestation.getAttestationHistory(randomAssetId(), 0);
      expect(history.length).to.equal(0);
    });
  });

  // ── Upgrade ────────────────────────────────────────────────────────────────

  describe("upgrades", () => {
    it("upgrades preserving attestation history", async () => {
      const assetId = randomAssetId();
      await attestation.connect(veriflowSigner).submitAttestation(assetId, randomHash(), 88);

      const NewFactory = await ethers.getContractFactory("HealthAttestation", admin);
      const upgraded = await upgrades.upgradeProxy(
        await attestation.getAddress(),
        NewFactory,
        { kind: "uups" }
      );

      const latest = await upgraded.getLatestAttestation(assetId);
      expect(latest.score).to.equal(88);
    });

    it("reverts upgrade if caller lacks UPGRADER_ROLE", async () => {
      const NewFactory = await ethers.getContractFactory("HealthAttestation", stranger);
      await expect(
        upgrades.upgradeProxy(await attestation.getAddress(), NewFactory, { kind: "uups" })
      ).to.be.reverted;
    });
  });
});
