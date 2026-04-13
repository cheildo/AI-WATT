import { ethers, upgrades, network } from "hardhat";

/**
 * deploy-proxy.ts — Phase 1 + Phase 2 deployment: WattUSD + MintEngine + sWattUSD
 *
 * Usage:
 *   npx hardhat run scripts/deploy-proxy.ts --network apothem
 *   npx hardhat run scripts/deploy-proxy.ts --network localhost
 *
 * Required env vars:
 *   DEPLOYER_PRIVATE_KEY  — funded with testnet XDC
 *   TREASURY_ADDRESS      — address that receives protocol fees
 *   USDC_ADDRESS          — USDC token address on target network
 *   USDT_ADDRESS          — USDT token address on target network (optional)
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  const chainId = (await ethers.provider.getNetwork()).chainId;

  console.log("=== AI WATT — Phase 1 + 2 Deploy ===");
  console.log(`Network : ${network.name} (chainId: ${chainId})`);
  console.log(`Deployer: ${deployer.address}`);
  console.log(`Balance : ${ethers.formatEther(await ethers.provider.getBalance(deployer.address))} XDC\n`);

  const treasuryAddress = process.env.TREASURY_ADDRESS;
  const usdcAddress = process.env.USDC_ADDRESS;

  if (!treasuryAddress) throw new Error("TREASURY_ADDRESS env var not set");
  if (!usdcAddress) throw new Error("USDC_ADDRESS env var not set");

  // ── 1. Deploy WattUSD (UUPS proxy) ───────────────────────────────────────
  console.log("Deploying WattUSD...");
  const WattUSDFactory = await ethers.getContractFactory("WattUSD");
  const wattUSD = await upgrades.deployProxy(WattUSDFactory, [deployer.address], {
    kind: "uups",
    initializer: "initialize",
  });
  await wattUSD.waitForDeployment();
  const wattUSDAddress = await wattUSD.getAddress();
  console.log(`✓ WattUSD proxy    : ${wattUSDAddress}`);
  console.log(`  Implementation   : ${await upgrades.erc1967.getImplementationAddress(wattUSDAddress)}\n`);

  // ── 2. Deploy MintEngine (UUPS proxy) ────────────────────────────────────
  console.log("Deploying MintEngine...");
  const MintEngineFactory = await ethers.getContractFactory("MintEngine");
  const mintEngine = await upgrades.deployProxy(
    MintEngineFactory,
    [deployer.address, wattUSDAddress, treasuryAddress],
    { kind: "uups", initializer: "initialize" }
  );
  await mintEngine.waitForDeployment();
  const mintEngineAddress = await mintEngine.getAddress();
  console.log(`✓ MintEngine proxy : ${mintEngineAddress}`);
  console.log(`  Implementation   : ${await upgrades.erc1967.getImplementationAddress(mintEngineAddress)}\n`);

  // ── 3. Wire: grant MINTER_ROLE on WattUSD to MintEngine ──────────────────
  console.log("Granting MINTER_ROLE on WattUSD to MintEngine...");
  const MINTER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("MINTER_ROLE"));
  const tx1 = await (wattUSD as any).grantRole(MINTER_ROLE, mintEngineAddress);
  await tx1.wait();
  console.log(`✓ MINTER_ROLE granted (tx: ${tx1.hash})\n`);

  // ── 4. Accept USDC in MintEngine ─────────────────────────────────────────
  console.log(`Accepting USDC (${usdcAddress}) in MintEngine...`);
  const tx2 = await (mintEngine as any).setAcceptedStablecoin(usdcAddress, true);
  await tx2.wait();
  console.log(`✓ USDC accepted (tx: ${tx2.hash})`);

  const usdtAddress = process.env.USDT_ADDRESS;
  if (usdtAddress) {
    console.log(`Accepting USDT (${usdtAddress}) in MintEngine...`);
    const tx3 = await (mintEngine as any).setAcceptedStablecoin(usdtAddress, true);
    await tx3.wait();
    console.log(`✓ USDT accepted (tx: ${tx3.hash})`);
  }

  // ── 5. Deploy sWattUSD (UUPS proxy) ──────────────────────────────────────
  console.log("\nDeploying sWattUSD...");
  const sWattFactory = await ethers.getContractFactory("sWattUSD");
  const sWattUSD = await upgrades.deployProxy(
    sWattFactory,
    [deployer.address, wattUSDAddress],
    { kind: "uups", initializer: "initialize" }
  );
  await sWattUSD.waitForDeployment();
  const sWattUSDAddress = await sWattUSD.getAddress();
  console.log(`✓ sWattUSD proxy   : ${sWattUSDAddress}`);
  console.log(`  Implementation   : ${await upgrades.erc1967.getImplementationAddress(sWattUSDAddress)}\n`);

  // ── 6. Seed initial deposit — inflation attack mitigation ─────────────────
  // Deploy script mints 1 WATT into sWattUSD so the exchange rate is anchored
  // before any user deposits. The deployer receives the corresponding sWATT shares.
  const SEED_AMOUNT = 1_000_000n; // 1 WATT (6 decimals)
  console.log("Seeding initial deposit (1 WATT) into sWattUSD...");

  // Grant MINTER_ROLE on WattUSD to deployer temporarily for seed mint
  const tx4 = await (wattUSD as any).grantRole(MINTER_ROLE, deployer.address);
  await tx4.wait();

  // Mint 1 WATT to deployer
  const tx5 = await (wattUSD as any).mint(deployer.address, SEED_AMOUNT);
  await tx5.wait();

  // Approve and deposit into sWattUSD
  const tx6 = await (wattUSD as any).approve(sWattUSDAddress, SEED_AMOUNT);
  await tx6.wait();
  const tx7 = await (sWattUSD as any).deposit(SEED_AMOUNT, deployer.address);
  await tx7.wait();
  console.log(`✓ Seed deposit complete (1 WATT → sWATT shares to deployer)\n`);

  // Grant YIELD_DISTRIBUTOR_ROLE to MintEngine address (LendingPool wired in Phase 4)
  // Skipped here — grant individually when LendingPool is deployed

  // ── Summary ───────────────────────────────────────────────────────────────
  console.log("=== Deployment complete ===");
  console.log(`WattUSD   : ${wattUSDAddress}`);
  console.log(`MintEngine: ${mintEngineAddress}`);
  console.log(`sWattUSD  : ${sWattUSDAddress}`);
  console.log(`Treasury  : ${treasuryAddress}`);
  console.log("\nAdd these to contracts/.env, then run verify.ts:");
  console.log(`WATT_USD_PROXY_ADDRESS=${wattUSDAddress}`);
  console.log(`MINT_ENGINE_PROXY_ADDRESS=${mintEngineAddress}`);
  console.log(`SWATT_USD_PROXY_ADDRESS=${sWattUSDAddress}`);
  console.log("\nNext steps:");
  console.log("  1. npm run verify:apothem");
  console.log("  2. Transfer DEFAULT_ADMIN_ROLE to a multisig");
  console.log("  3. Update backend/.env with contract addresses");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
