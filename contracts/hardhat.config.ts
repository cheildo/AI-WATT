import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";
import * as dotenv from "dotenv";

dotenv.config();

const DEPLOYER_PRIVATE_KEY = process.env.DEPLOYER_PRIVATE_KEY ?? "";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.24",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
      viaIR: false,
    },
  },

  networks: {
    hardhat: {
      chainId: 31337,
    },
    localhost: {
      url: "http://127.0.0.1:8545",
      chainId: 31337,
    },

    // XDC Apothem Testnet — chainId: 51
    apothem: {
      url: process.env.XDC_RPC_URL ?? "https://erpc.apothem.network",
      chainId: 51,
      accounts: DEPLOYER_PRIVATE_KEY ? [DEPLOYER_PRIVATE_KEY] : [],
      gasPrice: "auto",
    },

    // XDC Mainnet — chainId: 50
    xdc: {
      url: "https://erpc.xinfin.network",
      chainId: 50,
      accounts: DEPLOYER_PRIVATE_KEY ? [DEPLOYER_PRIVATE_KEY] : [],
      gasPrice: "auto",
    },
  },

  // ── Block explorer verification ────────────────────────────────────────────
  // Both explorers are BlockScout-based — no API key required.
  // apiKey must be a flat string (Etherscan v2 format); "placeholder" satisfies
  // hardhat-verify's non-empty check without being used for auth.
  etherscan: {
    apiKey: "placeholder",
    customChains: [
      {
        network: "apothem",
        chainId: 51,
        urls: {
          // BlockScout instance for XDC Apothem testnet
          apiURL: "https://explorer.apothem.network/api",
          browserURL: "https://explorer.apothem.network",
        },
      },
      {
        network: "xdc",
        chainId: 50,
        urls: {
          // BlockScout instance for XDC mainnet
          apiURL: "https://explorer.xinfin.network/api",
          browserURL: "https://explorer.xinfin.network",
        },
      },
    ],
  },

  paths: {
    sources: "./src",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts",
  },
};

export default config;
