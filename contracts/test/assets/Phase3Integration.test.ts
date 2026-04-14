import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { AssetRegistry, OCNFT, HealthAttestation } from "../../typechain-types";

/// @dev Full Phase 3 flow: register asset → mint OC-NFT → submit attestation → query all three.

describe("Phase 3 Integration", () => {
  let registry: AssetRegistry;
  let nft: OCNFT;
  let attestation: HealthAttestation;

  let admin: HardhatEthersSigner;
  let backendSigner: HardhatEthersSigner;  // holds REGISTRAR_ROLE + MINTER_ROLE + VERIFLOW_SIGNER
  let borrower: HardhatEthersSigner;

  const REGISTRAR_ROLE       = ethers.keccak256(ethers.toUtf8Bytes("REGISTRAR_ROLE"));
  const LENDINGPOOL_ROLE     = ethers.keccak256(ethers.toUtf8Bytes("LENDINGPOOL_ROLE"));
  const MINTER_ROLE          = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const VERIFLOW_SIGNER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("VERIFLOW_SIGNER"));

  before(async () => {
    [admin, backendSigner, borrower] = await ethers.getSigners();

    // Deploy all three contracts
    const RegistryFactory = await ethers.getContractFactory("AssetRegistry");
    registry = (await upgrades.deployProxy(RegistryFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as AssetRegistry;

    const NFTFactory = await ethers.getContractFactory("OCNFT");
    nft = (await upgrades.deployProxy(NFTFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as OCNFT;

    const AttestationFactory = await ethers.getContractFactory("HealthAttestation");
    attestation = (await upgrades.deployProxy(AttestationFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as HealthAttestation;

    // Grant roles to the backend signer (mirrors deploy script)
    await registry.connect(admin).grantRole(REGISTRAR_ROLE, backendSigner.address);
    await registry.connect(admin).grantRole(LENDINGPOOL_ROLE, backendSigner.address);
    await nft.connect(admin).grantRole(MINTER_ROLE, backendSigner.address);
    await attestation.connect(admin).grantRole(VERIFLOW_SIGNER_ROLE, backendSigner.address);
  });

  it("full lifecycle: register → activate → mint OC-NFT → attest → query all three", async () => {
    // ── Step 1: Generate assetId (done off-chain in production) ─────────────
    const assetId = ethers.keccak256(
      ethers.solidityPacked(
        ["string", "address", "uint256"],
        ["GPU-SN-12345", borrower.address, Date.now()]
      )
    );
    const metadataURI = "ipfs://QmAssetMetadata";
    const healthHash  = ethers.keccak256(
      ethers.solidityPacked(["bytes32", "uint8", "uint256"], [assetId, 88, Math.floor(Date.now() / 1000)])
    );

    // ── Step 2: Register asset in AssetRegistry ──────────────────────────────
    await registry.connect(backendSigner).registerAsset(
      assetId,
      0, // GPU_CLUSTER
      borrower.address,
      7000 // 70% LTV
    );

    // Verify registration
    const asset = await registry.getAsset(assetId);
    expect(asset.borrower).to.equal(borrower.address);
    expect(asset.status).to.equal(0); // PENDING
    expect(await registry.isActive(assetId)).to.be.false;

    // ── Step 3: Activate asset (admin or LendingPool after KYC/due diligence) ─
    await registry.connect(backendSigner).updateStatus(assetId, 1); // ACTIVE
    expect(await registry.isActive(assetId)).to.be.true;

    // ── Step 4: Mint OC-NFT for the borrower ────────────────────────────────
    const tx = await nft.connect(backendSigner).mintOCNFT(borrower.address, assetId, metadataURI);
    const receipt = await tx.wait();

    // Verify NFT
    const tokenId = await nft.getTokenId(assetId);
    expect(tokenId).to.equal(1n);
    expect(await nft.ownerOf(tokenId)).to.equal(borrower.address);
    expect(await nft.getAssetId(tokenId)).to.equal(assetId);
    expect(await nft.tokenURI(tokenId)).to.equal(metadataURI);

    // ── Step 5: Submit Veriflow health attestation ───────────────────────────
    await attestation.connect(backendSigner).submitAttestation(assetId, healthHash, 88);

    // Verify attestation
    expect(await attestation.hasAttestation(assetId)).to.be.true;
    const latest = await attestation.getLatestAttestation(assetId);
    expect(latest.assetId).to.equal(assetId);
    expect(latest.score).to.equal(88);
    expect(latest.healthHash).to.equal(healthHash);

    // ── Step 6: All three contracts agree on the assetId ────────────────────
    const registeredAsset = await registry.getAsset(assetId);
    const nftLinked       = await nft.getAssetId(tokenId);
    const attestedAsset   = (await attestation.getLatestAttestation(assetId)).assetId;

    expect(registeredAsset.assetId).to.equal(assetId);
    expect(nftLinked).to.equal(assetId);
    expect(attestedAsset).to.equal(assetId);

    // ── Step 7: LendingPool readiness check (simulated) ─────────────────────
    // Phase 4 LendingPool will check both of these before originating a loan
    const isActive    = await registry.isActive(assetId);
    const latestScore = (await attestation.getLatestAttestation(assetId)).score;

    expect(isActive).to.be.true;
    expect(latestScore).to.be.gte(60); // minimum score for loan origination
  });
});
