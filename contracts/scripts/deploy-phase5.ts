import { ethers, upgrades, network } from "hardhat";

/**
 * deploy-phase5.ts — Phase 5 deployment: WEVQueue
 *
 * Usage:
 *   npx hardhat run scripts/deploy-phase5.ts --network apothem
 *   npx hardhat run scripts/deploy-phase5.ts --network localhost
 *
 * Required env vars:
 *   DEPLOYER_PRIVATE_KEY       — funded with testnet XDC
 *   SWATT_USD_PROXY_ADDRESS    — deployed in Phase 2
 *   WATT_USD_PROXY_ADDRESS     — deployed in Phase 1
 *   BACKEND_SIGNER_ADDRESS     — hot wallet granted PROCESSOR_ROLE
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  const chainId = (await ethers.provider.getNetwork()).chainId;

  console.log("=== AI WATT — Phase 5 Deploy ===");
  console.log(`Network : ${network.name} (chainId: ${chainId})`);
  console.log(`Deployer: ${deployer.address}`);
  console.log(`Balance : ${ethers.formatEther(await ethers.provider.getBalance(deployer.address))} XDC\n`);

  const sWattUSDAddress   = process.env.SWATT_USD_PROXY_ADDRESS;
  const wattUSDAddress    = process.env.WATT_USD_PROXY_ADDRESS;
  const backendSigner     = process.env.BACKEND_SIGNER_ADDRESS;

  if (!sWattUSDAddress) throw new Error("SWATT_USD_PROXY_ADDRESS env var not set");
  if (!wattUSDAddress)  throw new Error("WATT_USD_PROXY_ADDRESS env var not set");
  if (!backendSigner)   throw new Error("BACKEND_SIGNER_ADDRESS env var not set");

  // ── 1. Deploy WEVQueue ────────────────────────────────────────────────────
  console.log("Deploying WEVQueue...");
  const QueueFactory = await ethers.getContractFactory("WEVQueue");
  const queue = await upgrades.deployProxy(
    QueueFactory,
    [deployer.address, sWattUSDAddress, wattUSDAddress],
    { initializer: "initialize", kind: "uups" }
  );
  await queue.waitForDeployment();
  const queueAddress = await queue.getAddress();
  console.log(`✓ WEVQueue proxy      : ${queueAddress}`);
  console.log(`  Implementation      : ${await upgrades.erc1967.getImplementationAddress(queueAddress)}\n`);

  // ── 2. Wire WEVQueue into sWattUSD ────────────────────────────────────────
  console.log(`Calling sWattUSD.setWEVQueue(${queueAddress})...`);
  const sWattUSD = await ethers.getContractAt("sWattUSD", sWattUSDAddress);
  const tx1 = await (sWattUSD as any).setWEVQueue(queueAddress);
  await tx1.wait();
  console.log(`✓ sWattUSD.wevQueue updated (tx: ${tx1.hash})\n`);

  // ── 3. Grant PROCESSOR_ROLE to backend signer ─────────────────────────────
  const PROCESSOR_ROLE = ethers.keccak256(ethers.toUtf8Bytes("PROCESSOR_ROLE"));
  console.log(`Granting PROCESSOR_ROLE to backend signer (${backendSigner})...`);
  const tx2 = await (queue as any).grantRole(PROCESSOR_ROLE, backendSigner);
  await tx2.wait();
  console.log(`✓ PROCESSOR_ROLE granted (tx: ${tx2.hash})\n`);

  // ── Summary ───────────────────────────────────────────────────────────────
  console.log("=== Phase 5 Deployment complete ===");
  console.log(`WEVQueue          : ${queueAddress}`);
  console.log(`sWattUSD          : ${sWattUSDAddress}  (wevQueue now set)`);
  console.log(`WattUSD           : ${wattUSDAddress}`);
  console.log(`Processor (keeper): ${backendSigner}`);
  console.log("\nAdd these to contracts/.env, then run verify.ts:");
  console.log(`WEV_QUEUE_PROXY_ADDRESS=${queueAddress}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
