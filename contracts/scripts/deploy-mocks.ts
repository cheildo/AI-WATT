import { ethers, network } from "hardhat";

/**
 * deploy-mocks.ts — Deploy MockStablecoin (USDC + USDT) on Apothem testnet.
 *
 * Usage:
 *   npx hardhat run scripts/deploy-mocks.ts --network apothem
 *   npx hardhat run scripts/deploy-mocks.ts --network localhost
 *
 * After running, copy the printed addresses into contracts/.env:
 *   USDC_ADDRESS=0x...
 *   USDT_ADDRESS=0x...
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  const chainId = (await ethers.provider.getNetwork()).chainId;

  console.log("=== AI WATT — Deploy Test Stablecoins ===");
  console.log(`Network : ${network.name} (chainId: ${chainId})`);
  console.log(`Deployer: ${deployer.address}`);
  console.log(
    `Balance : ${ethers.formatEther(
      await ethers.provider.getBalance(deployer.address)
    )} XDC\n`
  );

  const MockStablecoin = await ethers.getContractFactory("MockStablecoin");

  // ── Deploy mock USDC ───────────────────────────────────────────────────────
  console.log("Deploying MockUSDC...");
  const mockUSDC = await MockStablecoin.deploy("USD Coin", "USDC", 6);
  await mockUSDC.waitForDeployment();
  const usdcAddress = await mockUSDC.getAddress();
  console.log(`✓ MockUSDC : ${usdcAddress}`);

  // ── Deploy mock USDT ───────────────────────────────────────────────────────
  console.log("Deploying MockUSDT...");
  const mockUSDT = await MockStablecoin.deploy("Tether USD", "USDT", 6);
  await mockUSDT.waitForDeployment();
  const usdtAddress = await mockUSDT.getAddress();
  console.log(`✓ MockUSDT : ${usdtAddress}`);

  // ── Mint initial supply to deployer for testing ────────────────────────────
  const initialSupply = ethers.parseUnits("1000000", 6); // 1M each

  console.log("\nMinting 1,000,000 USDC to deployer...");
  const tx1 = await mockUSDC.mint(deployer.address, initialSupply);
  await tx1.wait();
  console.log(`✓ tx: ${tx1.hash}`);

  console.log("Minting 1,000,000 USDT to deployer...");
  const tx2 = await mockUSDT.mint(deployer.address, initialSupply);
  await tx2.wait();
  console.log(`✓ tx: ${tx2.hash}`);

  // ── Summary ────────────────────────────────────────────────────────────────
  console.log("\n=== Done. Add these to contracts/.env ===");
  console.log(`USDC_ADDRESS=${usdcAddress}`);
  console.log(`USDT_ADDRESS=${usdtAddress}`);
  console.log("\nAnyone can call faucet() on either contract to get 10,000 tokens.");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
