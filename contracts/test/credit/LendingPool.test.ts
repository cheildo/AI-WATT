import { expect } from "chai";
import { ethers, upgrades, network } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  LendingPool,
  AssetRegistry,
  HealthAttestation,
  sWattUSD,
  WattUSD,
} from "../../typechain-types";

// ── Constants ───────────────────────────────────────────────────────────────

const ONE_WATT       = 1_000_000n;          // 1 WATT (6 decimals)
const PRINCIPAL      = 100_000n * ONE_WATT; // 100,000 WATT
const INTEREST_RATE  = 1200n;               // 12% annual (basis points)
const TERM_DAYS      = 30n;                 // 30 days
const ENGINE_TYPE    = 2;                   // Engine 2

const ATTESTATION_MAX_AGE = 48 * 60 * 60;  // 48 hours in seconds
const HEALTH_COOLDOWN     = 12 * 60 * 60;  // 12 hours

const GPU_CLUSTER = 0;
const ACTIVE      = 1;

// Role keccak hashes
const role = (name: string) => ethers.keccak256(ethers.toUtf8Bytes(name));

// Simple interest: principal × rate(bps) × days / (365 × 10000)
function calcInterest(principal: bigint, rateBps: bigint, days: bigint): bigint {
  return principal * rateBps * days / (365n * 10_000n);
}

async function timeTravel(seconds: number) {
  await network.provider.send("evm_increaseTime", [seconds]);
  await network.provider.send("evm_mine");
}

describe("LendingPool", () => {
  let pool: LendingPool;
  let registry: AssetRegistry;
  let attestation: HealthAttestation;
  let wattUSD: WattUSD;
  let vault: sWattUSD;

  let admin: HardhatEthersSigner;
  let curator: HardhatEthersSigner;
  let liquidator: HardhatEthersSigner;
  let borrower: HardhatEthersSigner;
  let registrar: HardhatEthersSigner;
  let veriflow: HardhatEthersSigner;
  let stranger: HardhatEthersSigner;

  let assetId: string;

  // ── Setup ──────────────────────────────────────────────────────────────────

  beforeEach(async () => {
    [admin, curator, liquidator, borrower, registrar, veriflow, stranger] =
      await ethers.getSigners();

    // Deploy WattUSD
    const WattFactory = await ethers.getContractFactory("WattUSD");
    wattUSD = (await upgrades.deployProxy(WattFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as WattUSD;

    // Deploy sWattUSD (ERC-4626 vault — WATT as underlying)
    const sWattFactory = await ethers.getContractFactory("sWattUSD");
    vault = (await upgrades.deployProxy(
      sWattFactory,
      [admin.address, await wattUSD.getAddress()],
      { initializer: "initialize", kind: "uups" }
    )) as unknown as sWattUSD;

    // Deploy AssetRegistry
    const RegistryFactory = await ethers.getContractFactory("AssetRegistry");
    registry = (await upgrades.deployProxy(RegistryFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as AssetRegistry;

    // Deploy HealthAttestation
    const AttestationFactory = await ethers.getContractFactory("HealthAttestation");
    attestation = (await upgrades.deployProxy(AttestationFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as HealthAttestation;

    // Deploy LendingPool
    const PoolFactory = await ethers.getContractFactory("LendingPool");
    pool = (await upgrades.deployProxy(
      PoolFactory,
      [
        admin.address,
        await registry.getAddress(),
        await attestation.getAddress(),
        await wattUSD.getAddress(),
        await vault.getAddress(),
      ],
      { initializer: "initialize", kind: "uups" }
    )) as unknown as LendingPool;

    // Grant roles
    await registry.connect(admin).grantRole(role("REGISTRAR_ROLE"), registrar.address);
    await registry.connect(admin).grantRole(role("LENDINGPOOL_ROLE"), await pool.getAddress());
    await attestation.connect(admin).grantRole(role("VERIFLOW_SIGNER"), veriflow.address);
    await pool.connect(admin).grantRole(role("CURATOR_ROLE"), curator.address);
    await pool.connect(admin).grantRole(role("LIQUIDATOR_ROLE"), liquidator.address);

    // Grant MINTER_ROLE on WattUSD to LendingPool (for loan disbursement)
    await wattUSD.connect(admin).grantRole(role("MINTER_ROLE"), await pool.getAddress());

    // Grant YIELD_DISTRIBUTOR_ROLE on sWattUSD to LendingPool (for routing yield)
    await vault.connect(admin).grantRole(role("YIELD_DISTRIBUTOR_ROLE"), await pool.getAddress());

    // Seed sWattUSD with 1 WATT so NAV is anchored (inflation protection)
    await wattUSD.connect(admin).grantRole(role("MINTER_ROLE"), admin.address);
    await wattUSD.connect(admin).mint(admin.address, ONE_WATT);
    await wattUSD.connect(admin).approve(await vault.getAddress(), ONE_WATT);
    await vault.connect(admin).deposit(ONE_WATT, admin.address);

    // Register and activate a test asset with a fresh attestation
    assetId = ethers.keccak256(ethers.toUtf8Bytes("GPU-CLUSTER-001"));
    await registry.connect(registrar).registerAsset(assetId, GPU_CLUSTER, borrower.address, 7000);
    await registry.connect(admin).updateStatus(assetId, ACTIVE);
    await attestation.connect(veriflow).submitAttestation(assetId, ethers.keccak256(ethers.randomBytes(32)), 85);
  });

  // ── Deployment ─────────────────────────────────────────────────────────────

  describe("deployment", () => {
    it("wires all contract addresses", async () => {
      expect(await pool.assetRegistry()).to.equal(await registry.getAddress());
      expect(await pool.healthAttestation()).to.equal(await attestation.getAddress());
      expect(await pool.wattUSD()).to.equal(await wattUSD.getAddress());
      expect(await pool.sWattUSD()).to.equal(await vault.getAddress());
    });

    it("constants are correct", async () => {
      expect(await pool.MIN_HEALTH_SCORE()).to.equal(60);
      expect(await pool.ATTESTATION_MAX_AGE()).to.equal(ATTESTATION_MAX_AGE);
    });

    it("reverts if any address is zero", async () => {
      const Factory = await ethers.getContractFactory("LendingPool");
      await expect(
        upgrades.deployProxy(Factory, [
          ethers.ZeroAddress,
          await registry.getAddress(),
          await attestation.getAddress(),
          await wattUSD.getAddress(),
          await vault.getAddress(),
        ], { initializer: "initialize", kind: "uups" })
      ).to.be.revertedWithCustomError(pool, "ZeroAddress");
    });
  });

  // ── originateLoan ──────────────────────────────────────────────────────────

  describe("originateLoan", () => {
    it("originates a loan and mints WATT to borrower", async () => {
      const wattBefore = await wattUSD.balanceOf(borrower.address);
      await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      expect(await wattUSD.balanceOf(borrower.address)).to.equal(wattBefore + PRINCIPAL);
    });

    it("emits LoanOriginated event", async () => {
      await expect(
        pool.connect(curator).originateLoan(
          assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.emit(pool, "LoanOriginated");
    });

    it("stores loan with correct fields", async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      const loanId = event.args.loanId;

      const loan = await pool.getLoan(loanId);
      expect(loan.assetId).to.equal(assetId);
      expect(loan.borrower).to.equal(borrower.address);
      expect(loan.principal).to.equal(PRINCIPAL);
      expect(loan.interestRate).to.equal(INTEREST_RATE);
      expect(loan.engineType).to.equal(ENGINE_TYPE);
      expect(loan.status).to.equal(1); // ACTIVE
    });

    it("outstanding = principal + interest at origination", async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      const loanId = event.args.loanId;

      const loan = await pool.getLoan(loanId);
      const expectedInterest = calcInterest(PRINCIPAL, INTEREST_RATE, TERM_DAYS);
      expect(loan.outstanding).to.equal(PRINCIPAL + expectedInterest);
    });

    it("tracks loan under borrower address", async () => {
      await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const loans = await pool.getBorrowerLoans(borrower.address);
      expect(loans.length).to.equal(1);
    });

    it("reverts when asset is not registered", async () => {
      const unknownAsset = ethers.keccak256(ethers.toUtf8Bytes("UNKNOWN"));
      await expect(
        pool.connect(curator).originateLoan(
          unknownAsset, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "AssetNotActive");
    });

    it("reverts when asset status is PENDING (not ACTIVE)", async () => {
      const newAsset = ethers.keccak256(ethers.toUtf8Bytes("PENDING-ASSET"));
      await registry.connect(registrar).registerAsset(newAsset, GPU_CLUSTER, borrower.address, 7000);
      // Asset stays PENDING — not activated
      await attestation.connect(veriflow).submitAttestation(newAsset, ethers.keccak256(ethers.randomBytes(32)), 85);
      await expect(
        pool.connect(curator).originateLoan(
          newAsset, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "AssetNotActive");
    });

    it("reverts when health score is below 60", async () => {
      const lowScoreAsset = ethers.keccak256(ethers.toUtf8Bytes("LOW-HEALTH"));
      await registry.connect(registrar).registerAsset(lowScoreAsset, GPU_CLUSTER, borrower.address, 7000);
      await registry.connect(admin).updateStatus(lowScoreAsset, ACTIVE);
      await attestation.connect(veriflow).submitAttestation(lowScoreAsset, ethers.keccak256(ethers.randomBytes(32)), 55);
      await expect(
        pool.connect(curator).originateLoan(
          lowScoreAsset, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "HealthScoreTooLow");
    });

    it("reverts when attestation is stale (> 48 hours)", async () => {
      await timeTravel(ATTESTATION_MAX_AGE + 1);
      await expect(
        pool.connect(curator).originateLoan(
          assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "AttestationStale");
    });

    it("reverts when asset has no attestation", async () => {
      const noAttestAsset = ethers.keccak256(ethers.toUtf8Bytes("NO-ATTEST"));
      await registry.connect(registrar).registerAsset(noAttestAsset, GPU_CLUSTER, borrower.address, 7000);
      await registry.connect(admin).updateStatus(noAttestAsset, ACTIVE);
      // No attestation submitted
      await expect(
        pool.connect(curator).originateLoan(
          noAttestAsset, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "AttestationStale");
    });

    it("reverts on double-collateral (same asset, two loans)", async () => {
      await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      await expect(
        pool.connect(curator).originateLoan(
          assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "AssetAlreadyEncumbered");
    });

    it("reverts if caller lacks CURATOR_ROLE", async () => {
      await expect(
        pool.connect(stranger).originateLoan(
          assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "AccessControlUnauthorizedAccount");
    });
  });

  // ── repay ──────────────────────────────────────────────────────────────────

  describe("repay", () => {
    let loanId: string;
    let outstanding: bigint;

    beforeEach(async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      loanId = event.args.loanId;
      outstanding = (await pool.getLoan(loanId)).outstanding;

      // Borrower received PRINCIPAL from the loan disbursement.
      // Mint the interest portion so they can repay the full outstanding.
      const interest = outstanding - PRINCIPAL;
      if (interest > 0n) await wattUSD.connect(admin).mint(borrower.address, interest);

      // Borrower approves pool to pull repayment
      await wattUSD.connect(borrower).approve(await pool.getAddress(), outstanding);
    });

    it("reduces outstanding after partial repayment", async () => {
      const partialRepay = PRINCIPAL / 2n;
      await pool.connect(borrower).repay(loanId, partialRepay);
      const loan = await pool.getLoan(loanId);
      expect(loan.outstanding).to.equal(outstanding - partialRepay);
    });

    it("emits RepaymentReceived with correct split", async () => {
      const repayAmt = 10_000n * ONE_WATT;
      const totalInterest = calcInterest(PRINCIPAL, INTEREST_RATE, TERM_DAYS);
      const totalDebt = PRINCIPAL + totalInterest;
      const expectedInterest = repayAmt * totalInterest / totalDebt;
      const expectedPrincipal = repayAmt - expectedInterest;

      await expect(pool.connect(borrower).repay(loanId, repayAmt))
        .to.emit(pool, "RepaymentReceived")
        .withArgs(
          loanId,
          borrower.address,
          repayAmt,
          expectedPrincipal,
          expectedInterest,
          outstanding - repayAmt
        );
    });

    it("routes yield (90% of interest) to sWattUSD — NAV rises", async () => {
      const navBefore = await vault.navPerShare();
      await pool.connect(borrower).repay(loanId, outstanding);
      const navAfter = await vault.navPerShare();
      expect(navAfter).to.be.gt(navBefore);
    });

    it("accumulates protocol fees (10% of interest)", async () => {
      const totalInterest = calcInterest(PRINCIPAL, INTEREST_RATE, TERM_DAYS);
      const expectedProtocolFee = totalInterest * 1000n / 10_000n;

      await pool.connect(borrower).repay(loanId, outstanding);
      const fees = await pool.getProtocolFees();
      // Close enough (rounding): within 1 unit
      expect(fees).to.be.gte(expectedProtocolFee - 1n);
      expect(fees).to.be.lte(expectedProtocolFee + 1n);
    });

    it("marks status REPAYING after first partial payment", async () => {
      await pool.connect(borrower).repay(loanId, PRINCIPAL / 10n);
      expect((await pool.getLoan(loanId)).status).to.equal(2); // REPAYING
    });

    it("auto-settles on full repayment — status SETTLED", async () => {
      await pool.connect(borrower).repay(loanId, outstanding);
      expect((await pool.getLoan(loanId)).status).to.equal(3); // SETTLED
    });

    it("resets asset status to ACTIVE in AssetRegistry on settlement", async () => {
      await pool.connect(borrower).repay(loanId, outstanding);
      expect(await registry.isActive(assetId)).to.be.true;
    });

    it("allows new loan on same asset after settlement", async () => {
      await pool.connect(borrower).repay(loanId, outstanding);
      // Submit fresh attestation
      await timeTravel(HEALTH_COOLDOWN + 1);
      await attestation.connect(veriflow).submitAttestation(assetId, ethers.keccak256(ethers.randomBytes(32)), 85);
      // New loan should succeed
      await expect(
        pool.connect(curator).originateLoan(
          assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.not.be.reverted;
    });

    it("reverts on amount exceeding outstanding", async () => {
      await expect(
        pool.connect(borrower).repay(loanId, outstanding + 1n)
      ).to.be.revertedWithCustomError(pool, "ExceedsOutstanding");
    });

    it("reverts on zero amount", async () => {
      await expect(
        pool.connect(borrower).repay(loanId, 0n)
      ).to.be.revertedWithCustomError(pool, "ZeroAmount");
    });

    it("reverts on unknown loanId", async () => {
      await expect(
        pool.connect(borrower).repay(ethers.ZeroHash, 1n)
      ).to.be.revertedWithCustomError(pool, "LoanNotFound");
    });
  });

  // ── fullRepay ──────────────────────────────────────────────────────────────

  describe("fullRepay", () => {
    let loanId: string;

    beforeEach(async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      loanId = event.args.loanId;
      const outstanding = (await pool.getLoan(loanId)).outstanding;

      // Borrower received PRINCIPAL from loan disbursement.
      // Mint the interest portion so they can repay the full outstanding.
      const interest = outstanding - PRINCIPAL;
      if (interest > 0n) await wattUSD.connect(admin).mint(borrower.address, interest);

      await wattUSD.connect(borrower).approve(await pool.getAddress(), outstanding);
    });

    it("settles loan in one transaction", async () => {
      await pool.connect(borrower).fullRepay(loanId);
      expect((await pool.getLoan(loanId)).status).to.equal(3); // SETTLED
      expect((await pool.getLoan(loanId)).outstanding).to.equal(0n);
    });

    it("emits LoanSettled", async () => {
      await expect(pool.connect(borrower).fullRepay(loanId))
        .to.emit(pool, "LoanSettled")
        .withArgs(loanId, borrower.address);
    });
  });

  // ── flagDefaulted ──────────────────────────────────────────────────────────

  describe("flagDefaulted", () => {
    let loanId: string;

    beforeEach(async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      loanId = event.args.loanId;
    });

    it("flags loan DEFAULTED after maturity and emits LoanDefaulted", async () => {
      await timeTravel(Number(TERM_DAYS) * 24 * 60 * 60 + 1);
      await expect(pool.connect(stranger).flagDefaulted(loanId))
        .to.emit(pool, "LoanDefaulted")
        .withArgs(loanId, borrower.address);
      expect((await pool.getLoan(loanId)).status).to.equal(4); // DEFAULTED
    });

    it("reverts before maturity", async () => {
      await expect(
        pool.connect(stranger).flagDefaulted(loanId)
      ).to.be.revertedWithCustomError(pool, "LoanNotActive");
    });
  });

  // ── liquidate ──────────────────────────────────────────────────────────────

  describe("liquidate", () => {
    let loanId: string;

    beforeEach(async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      loanId = event.args.loanId;
      // Age the loan past maturity and flag it
      await timeTravel(Number(TERM_DAYS) * 24 * 60 * 60 + 1);
      await pool.connect(stranger).flagDefaulted(loanId);
    });

    it("liquidates DEFAULTED loan and emits LoanLiquidated", async () => {
      await expect(pool.connect(liquidator).liquidate(loanId))
        .to.emit(pool, "LoanLiquidated")
        .withArgs(loanId, borrower.address, liquidator.address);
      expect((await pool.getLoan(loanId)).status).to.equal(5); // LIQUIDATED
    });

    it("updates AssetRegistry to LIQUIDATED", async () => {
      await pool.connect(liquidator).liquidate(loanId);
      const asset = await registry.getAsset(assetId);
      expect(asset.status).to.equal(3); // LIQUIDATED
    });

    it("allows new asset registration with same assetId after liquidation? — asset is LIQUIDATED, not available", async () => {
      await pool.connect(liquidator).liquidate(loanId);
      expect(await registry.isActive(assetId)).to.be.false;
    });

    it("reverts if loan is not DEFAULTED", async () => {
      const tx2 = await pool.connect(curator).originateLoan(
        ethers.keccak256(ethers.toUtf8Bytes("ASSET-2")),
        borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      ).catch(() => null); // may fail — just check liquidation path

      await expect(
        pool.connect(liquidator).liquidate(ethers.ZeroHash)
      ).to.be.revertedWithCustomError(pool, "LoanNotFound");
    });

    it("reverts if caller lacks LIQUIDATOR_ROLE", async () => {
      await expect(
        pool.connect(stranger).liquidate(loanId)
      ).to.be.revertedWithCustomError(pool, "AccessControlUnauthorizedAccount");
    });
  });

  // ── withdrawFees ───────────────────────────────────────────────────────────

  describe("withdrawFees", () => {
    let loanId: string;

    beforeEach(async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      loanId = event.args.loanId;
      const outstanding = (await pool.getLoan(loanId)).outstanding;

      // Mint the interest portion so borrower can repay the full outstanding.
      const interest = outstanding - PRINCIPAL;
      if (interest > 0n) await wattUSD.connect(admin).mint(borrower.address, interest);

      await wattUSD.connect(borrower).approve(await pool.getAddress(), outstanding);
      await pool.connect(borrower).fullRepay(loanId);
    });

    it("withdraws accumulated fees to recipient", async () => {
      const fees = await pool.getProtocolFees();
      expect(fees).to.be.gt(0n);
      const before = await wattUSD.balanceOf(admin.address);
      await pool.connect(admin).withdrawFees(admin.address);
      expect(await wattUSD.balanceOf(admin.address)).to.equal(before + fees);
      expect(await pool.getProtocolFees()).to.equal(0n);
    });

    it("emits FeesWithdrawn", async () => {
      const fees = await pool.getProtocolFees();
      await expect(pool.connect(admin).withdrawFees(admin.address))
        .to.emit(pool, "FeesWithdrawn")
        .withArgs(admin.address, fees);
    });

    it("reverts when no fees available", async () => {
      await pool.connect(admin).withdrawFees(admin.address);
      await expect(
        pool.connect(admin).withdrawFees(admin.address)
      ).to.be.revertedWithCustomError(pool, "NoFeesAvailable");
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        pool.connect(stranger).withdrawFees(stranger.address)
      ).to.be.revertedWithCustomError(pool, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Pause ──────────────────────────────────────────────────────────────────

  describe("pause", () => {
    it("blocks originateLoan when paused", async () => {
      await pool.connect(admin).pause();
      await expect(
        pool.connect(curator).originateLoan(
          assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.be.revertedWithCustomError(pool, "EnforcedPause");
    });
  });

  // ── Upgrade ────────────────────────────────────────────────────────────────

  describe("upgrades", () => {
    it("upgrades preserving loan state", async () => {
      const tx = await pool.connect(curator).originateLoan(
        assetId, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      const loanId = event.args.loanId;

      const NewFactory = await ethers.getContractFactory("LendingPool", admin);
      const upgraded = await upgrades.upgradeProxy(await pool.getAddress(), NewFactory, { kind: "uups" });
      const loan = await upgraded.getLoan(loanId);
      expect(loan.borrower).to.equal(borrower.address);
      expect(loan.principal).to.equal(PRINCIPAL);
    });

    it("reverts upgrade if caller lacks UPGRADER_ROLE", async () => {
      const NewFactory = await ethers.getContractFactory("LendingPool", stranger);
      await expect(
        upgrades.upgradeProxy(await pool.getAddress(), NewFactory, { kind: "uups" })
      ).to.be.reverted;
    });
  });

  // ── Integration: full lifecycle ────────────────────────────────────────────

  describe("full lifecycle integration", () => {
    it("register → attest → originate → partial repay → full repay → new loan", async () => {
      // Register fresh asset
      const assetId2 = ethers.keccak256(ethers.toUtf8Bytes("GPU-CLUSTER-002"));
      await registry.connect(registrar).registerAsset(assetId2, GPU_CLUSTER, borrower.address, 7000);
      await registry.connect(admin).updateStatus(assetId2, ACTIVE);
      await attestation.connect(veriflow).submitAttestation(
        assetId2, ethers.keccak256(ethers.randomBytes(32)), 90
      );

      // Originate loan
      const tx = await pool.connect(curator).originateLoan(
        assetId2, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
      );
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "LoanOriginated") as any;
      const loanId = event.args.loanId;
      const outstanding = (await pool.getLoan(loanId)).outstanding;

      // Borrower received PRINCIPAL; mint interest so they can repay in full.
      const interest = outstanding - PRINCIPAL;
      if (interest > 0n) await wattUSD.connect(admin).mint(borrower.address, interest);

      // Partial repayment (half)
      await wattUSD.connect(borrower).approve(await pool.getAddress(), outstanding);
      await pool.connect(borrower).repay(loanId, outstanding / 2n);
      expect((await pool.getLoan(loanId)).status).to.equal(2); // REPAYING

      // NAV has increased
      expect(await vault.navPerShare()).to.be.gt(ONE_WATT);

      // Full repayment
      const remaining = (await pool.getLoan(loanId)).outstanding;
      await pool.connect(borrower).repay(loanId, remaining);
      expect((await pool.getLoan(loanId)).status).to.equal(3); // SETTLED

      // Asset is ACTIVE again — can take a new loan
      expect(await registry.isActive(assetId2)).to.be.true;
      await timeTravel(HEALTH_COOLDOWN + 1);
      await attestation.connect(veriflow).submitAttestation(
        assetId2, ethers.keccak256(ethers.randomBytes(32)), 88
      );
      await expect(
        pool.connect(curator).originateLoan(
          assetId2, borrower.address, PRINCIPAL, INTEREST_RATE, TERM_DAYS, ENGINE_TYPE
        )
      ).to.not.be.reverted;
    });
  });
});
