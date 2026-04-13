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
  // XDC explorers are BlockScout-based and support the Etherscan-compatible API.
  // No API key required for BlockScout — use a dummy value.
  etherscan: {
    apiKey: {
      apothem: "no-api-key-required",
      xdc: "no-api-key-required",
    },
    customChains: [
      {
        network: "apothem",
        chainId: 51,
        urls: {
          apiURL: "https://apothem.xinfinscan.com/api",
          browserURL: "https://apothem.xinfinscan.com",
        },
      },
      {
        network: "xdc",
        chainId: 50,
        urls: {
          apiURL: "https://xinfinscan.com/api",
          browserURL: "https://xinfinscan.com",
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
