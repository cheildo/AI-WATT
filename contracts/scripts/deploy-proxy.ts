import { ethers, upgrades } from "hardhat";

/**
 * deploy-proxy.ts — UUPS proxy deployment script.
 *
 * Usage:
 *   npx hardhat run scripts/deploy-proxy.ts --network apothem
 *
 * This script is a template. Uncomment and extend for each contract.
 * Always run hardhat-upgrades safety checks before deployment.
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying from:", deployer.address);
  console.log("Network:", (await ethers.provider.getNetwork()).name);

  // Example: deploy WattUSD (UUPS upgradeable)
  // const WattUSD = await ethers.getContractFactory("WattUSD");
  // const proxy = await upgrades.deployProxy(WattUSD, [deployer.address], {
  //   kind: "uups",
  //   initializer: "initialize",
  // });
  // await proxy.waitForDeployment();
  // console.log("WattUSD proxy deployed to:", await proxy.getAddress());

  console.log("No contracts to deploy yet — add them in deploy-proxy.ts");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
