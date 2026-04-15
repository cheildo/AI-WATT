import { ethers, upgrades, network } from "hardhat";

/**
 * deploy-phase4.ts — Phase 4 deployment: LendingPool
 *
 * Usage:
 *   npx hardhat run scripts/deploy-phase4.ts --network apothem
 *   npx hardhat run scripts/deploy-phase4.ts --network localhost
 *
 * Required env vars (from earlier deploy scripts):
 *   DEPLOYER_PRIVATE_KEY            — funded with testnet XDC
 *   ASSET_REGISTRY_PROXY_ADDRESS    — deployed in Phase 3
 *   HEALTH_ATTESTATION_PROXY_ADDRESS — deployed in Phase 3
 *   WATT_USD_PROXY_ADDRESS          — deployed in Phase 1
 *   SWATT_USD_PROXY_ADDRESS         — deployed in Phase 2
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  const chainId = (await ethers.provider.getNetwork()).chainId;

  console.log("=== AI WATT — Phase 4 Deploy ===");
  console.log(`Network : ${network.name} (chainId: ${chainId})`);
  console.log(`Deployer: ${deployer.address}`);
  console.log(`Balance : ${ethers.formatEther(await ethers.provider.getBalance(deployer.address))} XDC\n`);

  const assetRegistryAddress     = process.env.ASSET_REGISTRY_PROXY_ADDRESS;
  const healthAttestationAddress  = process.env.HEALTH_ATTESTATION_PROXY_ADDRESS;
  const wattUSDAddress            = process.env.WATT_USD_PROXY_ADDRESS;
  const sWattUSDAddress           = process.env.SWATT_USD_PROXY_ADDRESS;

  if (!assetRegistryAddress)    throw new Error("ASSET_REGISTRY_PROXY_ADDRESS env var not set");
  if (!healthAttestationAddress) throw new Error("HEALTH_ATTESTATION_PROXY_ADDRESS env var not set");
  if (!wattUSDAddress)           throw new Error("WATT_USD_PROXY_ADDRESS env var not set");
  if (!sWattUSDAddress)          throw new Error("SWATT_USD_PROXY_ADDRESS env var not set");

  // ── 1. Deploy LendingPool ─────────────────────────────────────────────────
  console.log("Deploying LendingPool...");
  const PoolFactory = await ethers.getContractFactory("LendingPool");
  const pool = await upgrades.deployProxy(
    PoolFactory,
    [
      deployer.address,
      assetRegistryAddress,
      healthAttestationAddress,
      wattUSDAddress,
      sWattUSDAddress,
    ],
    { initializer: "initialize", kind: "uups" }
  );
  await pool.waitForDeployment();
  const poolAddress = await pool.getAddress();
  console.log(`✓ LendingPool proxy   : ${poolAddress}`);
  console.log(`  Implementation      : ${await upgrades.erc1967.getImplementationAddress(poolAddress)}\n`);

  // ── 2. Wire roles ─────────────────────────────────────────────────────────
  const LENDINGPOOL_ROLE      = ethers.keccak256(ethers.toUtf8Bytes("LENDINGPOOL_ROLE"));
  const MINTER_ROLE           = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const YIELD_DISTRIBUTOR_ROLE = ethers.keccak256(ethers.toUtf8Bytes("YIELD_DISTRIBUTOR_ROLE"));

  console.log("Granting roles to LendingPool...");

  // AssetRegistry: grant LENDINGPOOL_ROLE so LendingPool can updateStatus on settle/liquidate
  const registry = await ethers.getContractAt("AssetRegistry", assetRegistryAddress);
  const tx1 = await (registry as any).grantRole(LENDINGPOOL_ROLE, poolAddress);
  await tx1.wait();
  console.log(`✓ LENDINGPOOL_ROLE on AssetRegistry (tx: ${tx1.hash})`);

  // WattUSD: grant MINTER_ROLE so LendingPool can mint principal to borrower on origination
  const wattUSD = await ethers.getContractAt("WattUSD", wattUSDAddress);
  const tx2 = await (wattUSD as any).grantRole(MINTER_ROLE, poolAddress);
  await tx2.wait();
  console.log(`✓ MINTER_ROLE on WattUSD (tx: ${tx2.hash})`);

  // sWattUSD: grant YIELD_DISTRIBUTOR_ROLE so LendingPool can call receiveYield
  const sWattUSD = await ethers.getContractAt("sWattUSD", sWattUSDAddress);
  const tx3 = await (sWattUSD as any).grantRole(YIELD_DISTRIBUTOR_ROLE, poolAddress);
  await tx3.wait();
  console.log(`✓ YIELD_DISTRIBUTOR_ROLE on sWattUSD (tx: ${tx3.hash})\n`);

  // ── Summary ───────────────────────────────────────────────────────────────
  console.log("=== Phase 4 Deployment complete ===");
  console.log(`LendingPool       : ${poolAddress}`);
  console.log(`AssetRegistry     : ${assetRegistryAddress}`);
  console.log(`HealthAttestation : ${healthAttestationAddress}`);
  console.log(`WattUSD           : ${wattUSDAddress}`);
  console.log(`sWattUSD          : ${sWattUSDAddress}`);
  console.log("\nAdd these to contracts/.env, then run verify.ts:");
  console.log(`LENDING_POOL_PROXY_ADDRESS=${poolAddress}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
