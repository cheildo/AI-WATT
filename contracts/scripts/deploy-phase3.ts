import { ethers, upgrades, network } from "hardhat";

/**
 * deploy-phase3.ts — Phase 3 deployment: AssetRegistry + OCNFT + HealthAttestation
 *
 * Usage:
 *   npx hardhat run scripts/deploy-phase3.ts --network apothem
 *   npx hardhat run scripts/deploy-phase3.ts --network localhost
 *
 * Required env vars:
 *   DEPLOYER_PRIVATE_KEY    — funded with testnet XDC
 *   BACKEND_SIGNER_ADDRESS  — hot wallet address granted REGISTRAR_ROLE, MINTER_ROLE, VERIFLOW_SIGNER
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  const chainId = (await ethers.provider.getNetwork()).chainId;

  console.log("=== AI WATT — Phase 3 Deploy ===");
  console.log(`Network : ${network.name} (chainId: ${chainId})`);
  console.log(`Deployer: ${deployer.address}`);
  console.log(`Balance : ${ethers.formatEther(await ethers.provider.getBalance(deployer.address))} XDC\n`);

  const backendSigner = process.env.BACKEND_SIGNER_ADDRESS;
  if (!backendSigner) throw new Error("BACKEND_SIGNER_ADDRESS env var not set");

  // ── 1. Deploy AssetRegistry ───────────────────────────────────────────────
  console.log("Deploying AssetRegistry...");
  const RegistryFactory = await ethers.getContractFactory("AssetRegistry");
  const registry = await upgrades.deployProxy(RegistryFactory, [deployer.address], {
    initializer: "initialize",
    kind: "uups",
  });
  await registry.waitForDeployment();
  const registryAddress = await registry.getAddress();
  console.log(`✓ AssetRegistry proxy : ${registryAddress}`);
  console.log(`  Implementation      : ${await upgrades.erc1967.getImplementationAddress(registryAddress)}\n`);

  // ── 2. Deploy OCNFT ───────────────────────────────────────────────────────
  console.log("Deploying OCNFT...");
  const NFTFactory = await ethers.getContractFactory("OCNFT");
  const nft = await upgrades.deployProxy(NFTFactory, [deployer.address], {
    initializer: "initialize",
    kind: "uups",
  });
  await nft.waitForDeployment();
  const nftAddress = await nft.getAddress();
  console.log(`✓ OCNFT proxy         : ${nftAddress}`);
  console.log(`  Implementation      : ${await upgrades.erc1967.getImplementationAddress(nftAddress)}\n`);

  // ── 3. Deploy HealthAttestation ───────────────────────────────────────────
  console.log("Deploying HealthAttestation...");
  const AttestationFactory = await ethers.getContractFactory("HealthAttestation");
  const attestation = await upgrades.deployProxy(AttestationFactory, [deployer.address], {
    initializer: "initialize",
    kind: "uups",
  });
  await attestation.waitForDeployment();
  const attestationAddress = await attestation.getAddress();
  console.log(`✓ HealthAttestation proxy : ${attestationAddress}`);
  console.log(`  Implementation          : ${await upgrades.erc1967.getImplementationAddress(attestationAddress)}\n`);

  // ── 4. Wire roles to backend signer ──────────────────────────────────────
  const REGISTRAR_ROLE    = ethers.keccak256(ethers.toUtf8Bytes("REGISTRAR_ROLE"));
  const MINTER_ROLE       = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const VERIFLOW_SIGNER   = ethers.keccak256(ethers.toUtf8Bytes("VERIFLOW_SIGNER"));

  console.log(`Granting roles to backend signer (${backendSigner})...`);

  const tx1 = await (registry as any).grantRole(REGISTRAR_ROLE, backendSigner);
  await tx1.wait();
  console.log(`✓ REGISTRAR_ROLE on AssetRegistry (tx: ${tx1.hash})`);

  const tx2 = await (nft as any).grantRole(MINTER_ROLE, backendSigner);
  await tx2.wait();
  console.log(`✓ MINTER_ROLE on OCNFT (tx: ${tx2.hash})`);

  const tx3 = await (attestation as any).grantRole(VERIFLOW_SIGNER, backendSigner);
  await tx3.wait();
  console.log(`✓ VERIFLOW_SIGNER on HealthAttestation (tx: ${tx3.hash})\n`);

  // Note: LENDINGPOOL_ROLE on AssetRegistry is granted in Phase 4 deploy script
  // after LendingPool is deployed.

  // ── Summary ───────────────────────────────────────────────────────────────
  console.log("=== Phase 3 Deployment complete ===");
  console.log(`AssetRegistry     : ${registryAddress}`);
  console.log(`OCNFT             : ${nftAddress}`);
  console.log(`HealthAttestation : ${attestationAddress}`);
  console.log(`Backend signer    : ${backendSigner}`);
  console.log("\nAdd these to contracts/.env, then run verify-phase3.ts:");
  console.log(`ASSET_REGISTRY_PROXY_ADDRESS=${registryAddress}`);
  console.log(`OCNFT_PROXY_ADDRESS=${nftAddress}`);
  console.log(`HEALTH_ATTESTATION_PROXY_ADDRESS=${attestationAddress}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
