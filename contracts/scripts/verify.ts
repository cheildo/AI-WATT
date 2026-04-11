import { run } from "hardhat";

/**
 * verify.ts — Verifies deployed contracts on XDC block explorer.
 *
 * Usage:
 *   npx hardhat run scripts/verify.ts --network apothem
 */
async function main() {
  // Example:
  // await run("verify:verify", {
  //   address: "0xYOUR_PROXY_ADDRESS",
  //   constructorArguments: [],
  // });
  console.log("No contracts to verify yet — add addresses in verify.ts");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
