import { run, upgrades, network } from "hardhat";

/**
 * verify.ts — Verify all deployed contracts on the XDC block explorer.
 *
 * For UUPS proxies we verify the implementation contract (the proxy bytecode
 * is a well-known ERC-1967 pattern that explorers recognise automatically).
 *
 * Usage:
 *   npx hardhat run scripts/verify.ts --network apothem
 *   npx hardhat run scripts/verify.ts --network xdc
 *
 * Required env vars (copy from deploy-proxy.ts output):
 *   WATT_USD_PROXY_ADDRESS
 *   MINT_ENGINE_PROXY_ADDRESS
 *   SWATT_USD_PROXY_ADDRESS
 *
 * Optional (copy from deploy-mocks.ts output):
 *   USDC_ADDRESS
 *   USDT_ADDRESS
 */
async function verifyImpl(name: string, proxyAddress: string) {
  const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log(`\n${name}`);
  console.log(`  Proxy          : ${proxyAddress}`);
  console.log(`  Implementation : ${implAddress}`);
  try {
    // Implementation has no constructor args (initializer replaces constructor)
    await run("verify:verify", {
      address: implAddress,
      constructorArguments: [],
    });
    console.log(`  ✓ Verified`);
  } catch (err: any) {
    if (err.message?.toLowerCase().includes("already verified")) {
      console.log(`  ✓ Already verified — skipping`);
    } else {
      console.error(`  ✗ Failed:`, err.message ?? err);
    }
  }
}

async function verifyPlain(
  name: string,
  address: string,
  constructorArgs: unknown[]
) {
  console.log(`\n${name} at ${address}`);
  try {
    await run("verify:verify", {
      address,
      constructorArguments: constructorArgs,
    });
    console.log(`  ✓ Verified`);
  } catch (err: any) {
    if (err.message?.toLowerCase().includes("already verified")) {
      console.log(`  ✓ Already verified — skipping`);
    } else {
      console.error(`  ✗ Failed:`, err.message ?? err);
    }
  }
}

async function main() {
  const explorerURL =
    network.name === "apothem"
      ? "https://testnet.xdcscan.com"
      : "https://xdcscan.com";

  console.log("=== AI WATT — Contract Verification ===");
  console.log(`Network : ${network.name}`);
  console.log(`Explorer: ${explorerURL}\n`);

  const wattUSDProxy = process.env.WATT_USD_PROXY_ADDRESS;
  const mintEngineProxy = process.env.MINT_ENGINE_PROXY_ADDRESS;
  const sWattUSDProxy = process.env.SWATT_USD_PROXY_ADDRESS;
  const assetRegistryProxy = process.env.ASSET_REGISTRY_PROXY_ADDRESS;
  const ocnftProxy = process.env.OCNFT_PROXY_ADDRESS;
  const healthAttestationProxy = process.env.HEALTH_ATTESTATION_PROXY_ADDRESS;
  const lendingPoolProxy = process.env.LENDING_POOL_PROXY_ADDRESS;
  const wevQueueProxy    = process.env.WEV_QUEUE_PROXY_ADDRESS;
  const usdcAddress = process.env.USDC_ADDRESS;
  const usdtAddress = process.env.USDT_ADDRESS;

  // ── UUPS proxies — verify implementation contract ─────────────────────────
  if (wattUSDProxy) {
    await verifyImpl("WattUSD", wattUSDProxy);
  } else {
    console.warn("WATT_USD_PROXY_ADDRESS not set — skipping WattUSD");
  }

  if (mintEngineProxy) {
    await verifyImpl("MintEngine", mintEngineProxy);
  } else {
    console.warn("MINT_ENGINE_PROXY_ADDRESS not set — skipping MintEngine");
  }

  if (sWattUSDProxy) {
    await verifyImpl("sWattUSD", sWattUSDProxy);
  } else {
    console.warn("SWATT_USD_PROXY_ADDRESS not set — skipping sWattUSD");
  }

  if (assetRegistryProxy) {
    await verifyImpl("AssetRegistry", assetRegistryProxy);
  } else {
    console.warn("ASSET_REGISTRY_PROXY_ADDRESS not set — skipping AssetRegistry");
  }

  if (ocnftProxy) {
    await verifyImpl("OCNFT", ocnftProxy);
  } else {
    console.warn("OCNFT_PROXY_ADDRESS not set — skipping OCNFT");
  }

  if (healthAttestationProxy) {
    await verifyImpl("HealthAttestation", healthAttestationProxy);
  } else {
    console.warn("HEALTH_ATTESTATION_PROXY_ADDRESS not set — skipping HealthAttestation");
  }

  if (lendingPoolProxy) {
    await verifyImpl("LendingPool", lendingPoolProxy);
  } else {
    console.warn("LENDING_POOL_PROXY_ADDRESS not set — skipping LendingPool");
  }

  if (wevQueueProxy) {
    await verifyImpl("WEVQueue", wevQueueProxy);
  } else {
    console.warn("WEV_QUEUE_PROXY_ADDRESS not set — skipping WEVQueue");
  }

  // ── Mock stablecoins — plain contracts, pass constructor args ─────────────
  if (usdcAddress) {
    await verifyPlain("MockUSDC", usdcAddress, ["USD Coin", "USDC", 6]);
  }
  if (usdtAddress) {
    await verifyPlain("MockUSDT", usdtAddress, ["Tether USD", "USDT", 6]);
  }

  console.log(`\n=== Done — view at ${explorerURL} ===`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
