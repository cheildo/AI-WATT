import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { OCNFT } from "../../typechain-types";

function randomAssetId(): string {
  return ethers.keccak256(ethers.randomBytes(32));
}

const SAMPLE_URI = "ipfs://QmSampleMetadataHash";

describe("OCNFT", () => {
  let nft: OCNFT;

  let admin: HardhatEthersSigner;
  let minter: HardhatEthersSigner;
  let borrower: HardhatEthersSigner;
  let alice: HardhatEthersSigner;
  let stranger: HardhatEthersSigner;

  const MINTER_ROLE   = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const UPGRADER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("UPGRADER_ROLE"));

  beforeEach(async () => {
    [admin, minter, borrower, alice, stranger] = await ethers.getSigners();

    const Factory = await ethers.getContractFactory("OCNFT");
    nft = (await upgrades.deployProxy(Factory, [admin.address], {
      initializer: "initialize",
      kind: "uups",
    })) as unknown as OCNFT;

    await nft.connect(admin).grantRole(MINTER_ROLE, minter.address);
  });

  // ── Deployment ─────────────────────────────────────────────────────────────

  describe("deployment", () => {
    it("sets ERC-721 name and symbol", async () => {
      expect(await nft.name()).to.equal("AI WATT Ownership Certificate");
      expect(await nft.symbol()).to.equal("OC-NFT");
    });

    it("grants admin roles to admin", async () => {
      expect(await nft.hasRole(ethers.ZeroHash, admin.address)).to.be.true;
      expect(await nft.hasRole(UPGRADER_ROLE, admin.address)).to.be.true;
    });

    it("reverts on zero admin address", async () => {
      const Factory = await ethers.getContractFactory("OCNFT");
      await expect(
        upgrades.deployProxy(Factory, [ethers.ZeroAddress], {
          initializer: "initialize",
          kind: "uups",
        })
      ).to.be.revertedWithCustomError(nft, "ZeroAddress");
    });
  });

  // ── mintOCNFT ──────────────────────────────────────────────────────────────

  describe("mintOCNFT", () => {
    it("mints a token and emits OCNFTMinted", async () => {
      const assetId = randomAssetId();
      await expect(nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI))
        .to.emit(nft, "OCNFTMinted")
        .withArgs(1n, borrower.address, assetId);
    });

    it("assigns tokenId starting at 1", async () => {
      const assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);
      expect(await nft.ownerOf(1n)).to.equal(borrower.address);
    });

    it("increments tokenId for each mint", async () => {
      await nft.connect(minter).mintOCNFT(borrower.address, randomAssetId(), SAMPLE_URI);
      await nft.connect(minter).mintOCNFT(alice.address, randomAssetId(), SAMPLE_URI);
      expect(await nft.ownerOf(1n)).to.equal(borrower.address);
      expect(await nft.ownerOf(2n)).to.equal(alice.address);
    });

    it("sets the tokenURI from metadataURI", async () => {
      const assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);
      expect(await nft.tokenURI(1n)).to.equal(SAMPLE_URI);
    });

    it("links tokenId to assetId bidirectionally", async () => {
      const assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);
      expect(await nft.getAssetId(1n)).to.equal(assetId);
      expect(await nft.getTokenId(assetId)).to.equal(1n);
    });

    it("reverts on duplicate mint for same assetId", async () => {
      const assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);
      await expect(
        nft.connect(minter).mintOCNFT(alice.address, assetId, SAMPLE_URI)
      ).to.be.revertedWithCustomError(nft, "AlreadyMinted");
    });

    it("reverts on zero recipient address", async () => {
      await expect(
        nft.connect(minter).mintOCNFT(ethers.ZeroAddress, randomAssetId(), SAMPLE_URI)
      ).to.be.revertedWithCustomError(nft, "ZeroAddress");
    });

    it("reverts if caller lacks MINTER_ROLE", async () => {
      await expect(
        nft.connect(stranger).mintOCNFT(borrower.address, randomAssetId(), SAMPLE_URI)
      ).to.be.revertedWithCustomError(nft, "AccessControlUnauthorizedAccount");
    });
  });

  // ── burnOCNFT ──────────────────────────────────────────────────────────────

  describe("burnOCNFT", () => {
    let assetId: string;

    beforeEach(async () => {
      assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);
    });

    it("burns the token and emits OCNFTBurned", async () => {
      await expect(nft.connect(minter).burnOCNFT(1n))
        .to.emit(nft, "OCNFTBurned")
        .withArgs(1n, assetId);
    });

    it("clears both mappings after burn", async () => {
      await nft.connect(minter).burnOCNFT(1n);
      // getAssetId should revert (token no longer exists)
      await expect(nft.getAssetId(1n)).to.be.revertedWithCustomError(nft, "TokenNotFound");
      // getTokenId returns 0 (mapping cleared)
      expect(await nft.getTokenId(assetId)).to.equal(0n);
    });

    it("allows reminting the same assetId after burn", async () => {
      await nft.connect(minter).burnOCNFT(1n);
      // After clearing the mapping, a new mint should succeed
      await expect(
        nft.connect(minter).mintOCNFT(alice.address, assetId, SAMPLE_URI)
      ).to.not.be.reverted;
    });

    it("reverts on non-existent tokenId", async () => {
      await expect(nft.connect(minter).burnOCNFT(999n))
        .to.be.revertedWithCustomError(nft, "TokenNotFound");
    });

    it("reverts if caller lacks MINTER_ROLE", async () => {
      await expect(nft.connect(stranger).burnOCNFT(1n))
        .to.be.revertedWithCustomError(nft, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Soulbound (transfer restriction) ──────────────────────────────────────

  describe("soulbound transfer restriction", () => {
    let assetId: string;

    beforeEach(async () => {
      assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);
    });

    it("blocks transferFrom by token owner", async () => {
      await expect(
        nft.connect(borrower).transferFrom(borrower.address, alice.address, 1n)
      ).to.be.revertedWithCustomError(nft, "TransferRestricted");
    });

    it("blocks safeTransferFrom by token owner", async () => {
      await expect(
        nft.connect(borrower)["safeTransferFrom(address,address,uint256)"](
          borrower.address, alice.address, 1n
        )
      ).to.be.revertedWithCustomError(nft, "TransferRestricted");
    });

    it("blocks transfer even with approval", async () => {
      await nft.connect(borrower).approve(alice.address, 1n);
      await expect(
        nft.connect(alice).transferFrom(borrower.address, alice.address, 1n)
      ).to.be.revertedWithCustomError(nft, "TransferRestricted");
    });

    it("allows transfer by MINTER_ROLE (settlement flow)", async () => {
      // MINTER_ROLE can transfer the OC-NFT on settlement
      await nft.connect(minter).transferFrom(borrower.address, alice.address, 1n);
      expect(await nft.ownerOf(1n)).to.equal(alice.address);
    });
  });

  // ── View functions ─────────────────────────────────────────────────────────

  describe("view functions", () => {
    it("getAssetId reverts for non-existent token", async () => {
      await expect(nft.getAssetId(99n)).to.be.revertedWithCustomError(nft, "TokenNotFound");
    });

    it("getTokenId returns 0 for unminted assetId", async () => {
      expect(await nft.getTokenId(randomAssetId())).to.equal(0n);
    });

    it("supportsInterface: ERC-721 and ERC-165", async () => {
      expect(await nft.supportsInterface("0x80ac58cd")).to.be.true; // ERC-721
      expect(await nft.supportsInterface("0x01ffc9a7")).to.be.true; // ERC-165
    });
  });

  // ── Upgrade ────────────────────────────────────────────────────────────────

  describe("upgrades", () => {
    it("upgrades preserving token state", async () => {
      const assetId = randomAssetId();
      await nft.connect(minter).mintOCNFT(borrower.address, assetId, SAMPLE_URI);

      const NewFactory = await ethers.getContractFactory("OCNFT", admin);
      const upgraded = await upgrades.upgradeProxy(await nft.getAddress(), NewFactory, { kind: "uups" });
      expect(await upgraded.ownerOf(1n)).to.equal(borrower.address);
      expect(await upgraded.getAssetId(1n)).to.equal(assetId);
    });

    it("reverts upgrade if caller lacks UPGRADER_ROLE", async () => {
      const NewFactory = await ethers.getContractFactory("OCNFT", stranger);
      await expect(
        upgrades.upgradeProxy(await nft.getAddress(), NewFactory, { kind: "uups" })
      ).to.be.reverted;
    });
  });
});
