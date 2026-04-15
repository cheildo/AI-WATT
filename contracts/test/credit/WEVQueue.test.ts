import { expect } from "chai";
import { ethers, upgrades, network } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  WEVQueue,
  sWattUSD,
  WattUSD,
} from "../../typechain-types";

// ── Constants ───────────────────────────────────────────────────────────────

const ONE_WATT       = 1_000_000n;           // 1 WATT (6 decimals)
const WEV_THRESHOLD  = 100_000n * ONE_WATT;  // 100,000 WATT (sWattUSD default)
const DEPOSIT_AMOUNT = 500_000n * ONE_WATT;  // 500,000 WATT — enough for large redeems
const REDEEM_AMOUNT  = 200_000n * ONE_WATT;  // 200,000 sWATT (above threshold)

const PRIORITY_FEE_BPS = 50n;  // 0.5%

// Role keccak hashes
const role = (name: string) => ethers.keccak256(ethers.toUtf8Bytes(name));

async function timeTravel(seconds: number) {
  await network.provider.send("evm_increaseTime", [seconds]);
  await network.provider.send("evm_mine");
}

describe("WEVQueue", () => {
  let queue: WEVQueue;
  let vault: sWattUSD;
  let wattUSD: WattUSD;

  let admin: HardhatEthersSigner;
  let processor: HardhatEthersSigner;
  let user: HardhatEthersSigner;
  let user2: HardhatEthersSigner;
  let stranger: HardhatEthersSigner;

  // Shares the user holds after depositing
  let userShares: bigint;

  // ── Setup ──────────────────────────────────────────────────────────────────

  beforeEach(async () => {
    [admin, processor, user, user2, stranger] = await ethers.getSigners();

    // Deploy WattUSD
    const WattFactory = await ethers.getContractFactory("WattUSD");
    wattUSD = (await upgrades.deployProxy(WattFactory, [admin.address], {
      initializer: "initialize", kind: "uups",
    })) as unknown as WattUSD;

    // Deploy sWattUSD
    const sWattFactory = await ethers.getContractFactory("sWattUSD");
    vault = (await upgrades.deployProxy(
      sWattFactory,
      [admin.address, await wattUSD.getAddress()],
      { initializer: "initialize", kind: "uups" }
    )) as unknown as sWattUSD;

    // Deploy WEVQueue
    const QueueFactory = await ethers.getContractFactory("WEVQueue");
    queue = (await upgrades.deployProxy(
      QueueFactory,
      [admin.address, await vault.getAddress(), await wattUSD.getAddress()],
      { initializer: "initialize", kind: "uups" }
    )) as unknown as WEVQueue;

    // Grant MINTER_ROLE on WattUSD to admin for test minting
    await wattUSD.connect(admin).grantRole(role("MINTER_ROLE"), admin.address);

    // Seed sWattUSD with 1 WATT (inflation protection)
    await wattUSD.connect(admin).mint(admin.address, ONE_WATT);
    await wattUSD.connect(admin).approve(await vault.getAddress(), ONE_WATT);
    await vault.connect(admin).deposit(ONE_WATT, admin.address);

    // Wire WEVQueue into sWattUSD
    await vault.connect(admin).setWEVQueue(await queue.getAddress());

    // Grant PROCESSOR_ROLE to processor
    await queue.connect(admin).grantRole(role("PROCESSOR_ROLE"), processor.address);

    // Fund user with WATT and deposit into sWattUSD → get sWATT shares
    await wattUSD.connect(admin).mint(user.address, DEPOSIT_AMOUNT);
    await wattUSD.connect(user).approve(await vault.getAddress(), DEPOSIT_AMOUNT);
    await vault.connect(user).deposit(DEPOSIT_AMOUNT, user.address);
    userShares = await vault.balanceOf(user.address);

    // Fund user2 with a smaller WATT amount + shares
    const user2Deposit = 10_000n * ONE_WATT;
    await wattUSD.connect(admin).mint(user2.address, user2Deposit);
    await wattUSD.connect(user2).approve(await vault.getAddress(), user2Deposit);
    await vault.connect(user2).deposit(user2Deposit, user2.address);
  });

  // ── Deployment ─────────────────────────────────────────────────────────────

  describe("deployment", () => {
    it("wires sWattUSD and wattUSD addresses", async () => {
      expect(await queue.sWattUSD()).to.equal(await vault.getAddress());
      expect(await queue.wattUSD()).to.equal(await wattUSD.getAddress());
    });

    it("constants are correct", async () => {
      expect(await queue.PRIORITY_FEE_BPS()).to.equal(50);
      expect(await queue.STANDARD_WAIT()).to.equal(30n * 24n * 60n * 60n);
      expect(await queue.PRIORITY_WAIT()).to.equal(3n * 24n * 60n * 60n);
    });

    it("reverts if any address is zero on initialize", async () => {
      const Factory = await ethers.getContractFactory("WEVQueue");
      await expect(
        upgrades.deployProxy(Factory, [
          ethers.ZeroAddress,
          await vault.getAddress(),
          await wattUSD.getAddress(),
        ], { initializer: "initialize", kind: "uups" })
      ).to.be.revertedWithCustomError(queue, "ZeroAddress");
    });

    it("sWattUSD wevQueue points to this contract", async () => {
      expect(await vault.wevQueue()).to.equal(await queue.getAddress());
    });
  });

  // ── requestRedeem ──────────────────────────────────────────────────────────

  describe("requestRedeem", () => {
    beforeEach(async () => {
      await vault.connect(user).approve(await queue.getAddress(), userShares);
    });

    it("pulls sWATT from user and creates QUEUED request", async () => {
      const sharesBefore = await vault.balanceOf(user.address);
      await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      expect(await vault.balanceOf(user.address)).to.equal(sharesBefore - REDEEM_AMOUNT);
      expect(await vault.balanceOf(await queue.getAddress())).to.be.gte(REDEEM_AMOUNT);
    });

    it("emits RedemptionRequested with isPriority=false", async () => {
      await expect(queue.connect(user).requestRedeem(REDEEM_AMOUNT))
        .to.emit(queue, "RedemptionRequested")
        .withArgs(
          // requestId is any bytes32
          (id: string) => id !== ethers.ZeroHash,
          user.address,
          REDEEM_AMOUNT,
          false
        );
    });

    it("creates request with correct fields", async () => {
      const tx = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any;
      const requestId = event.args.requestId;

      const req = await queue.getRequest(requestId);
      expect(req.user).to.equal(user.address);
      expect(req.sWattAmount).to.equal(REDEEM_AMOUNT);
      expect(req.priorityFee).to.equal(0n);
      expect(req.status).to.equal(0); // QUEUED
    });

    it("increments queue depth", async () => {
      expect(await queue.getQueueDepth()).to.equal(0n);
      await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      expect(await queue.getQueueDepth()).to.equal(1n);
    });

    it("tracks request under user address", async () => {
      await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const ids = await queue.getUserRequests(user.address);
      expect(ids.length).to.equal(1);
    });

    it("reverts on zero amount", async () => {
      await expect(
        queue.connect(user).requestRedeem(0n)
      ).to.be.revertedWithCustomError(queue, "ZeroAmount");
    });

    it("reverts when paused", async () => {
      await queue.connect(admin).pause();
      await expect(
        queue.connect(user).requestRedeem(REDEEM_AMOUNT)
      ).to.be.revertedWithCustomError(queue, "EnforcedPause");
    });
  });

  // ── requestPriorityRedeem ─────────────────────────────────────────────────

  describe("requestPriorityRedeem", () => {
    let minFee: bigint;

    beforeEach(async () => {
      minFee = REDEEM_AMOUNT * PRIORITY_FEE_BPS / 10_000n;
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      // Mint WATT for priority fee
      await wattUSD.connect(admin).mint(user.address, minFee * 2n);
      await wattUSD.connect(user).approve(await queue.getAddress(), minFee * 2n);
    });

    it("creates priority request and pulls sWATT + WATT fee", async () => {
      const sharesBefore = await vault.balanceOf(user.address);
      const wattBefore   = await wattUSD.balanceOf(user.address);

      await queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee);

      expect(await vault.balanceOf(user.address)).to.equal(sharesBefore - REDEEM_AMOUNT);
      expect(await wattUSD.balanceOf(user.address)).to.equal(wattBefore - minFee);
    });

    it("emits RedemptionRequested with isPriority=true", async () => {
      await expect(queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee))
        .to.emit(queue, "RedemptionRequested")
        .withArgs(
          (id: string) => id !== ethers.ZeroHash,
          user.address,
          REDEEM_AMOUNT,
          true
        );
    });

    it("stores priority fee in request", async () => {
      const tx = await queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee);
      const receipt = await tx.wait();
      const event = receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any;
      const req = await queue.getRequest(event.args.requestId);
      expect(req.priorityFee).to.equal(minFee);
    });

    it("accepts fee above minimum", async () => {
      await expect(
        queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee * 2n)
      ).to.not.be.reverted;
    });

    it("reverts when fee is below minimum", async () => {
      await expect(
        queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee - 1n)
      ).to.be.revertedWithCustomError(queue, "InsufficientPriorityFee");
    });

    it("reverts on zero amount", async () => {
      await expect(
        queue.connect(user).requestPriorityRedeem(0n, 0n)
      ).to.be.revertedWithCustomError(queue, "ZeroAmount");
    });
  });

  // ── cancelRequest ─────────────────────────────────────────────────────────

  describe("cancelRequest", () => {
    let standardId: string;
    let priorityId: string;
    let minFee: bigint;

    beforeEach(async () => {
      minFee = REDEEM_AMOUNT * PRIORITY_FEE_BPS / 10_000n;
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      await wattUSD.connect(admin).mint(user.address, minFee);
      await wattUSD.connect(user).approve(await queue.getAddress(), minFee);

      // Standard request
      let tx = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      let receipt = await tx.wait();
      standardId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      // Priority request
      tx = await queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee);
      receipt = await tx.wait();
      priorityId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;
    });

    it("returns sWATT and marks standard request CANCELLED", async () => {
      const sharesBefore = await vault.balanceOf(user.address);
      await queue.connect(user).cancelRequest(standardId);
      expect(await vault.balanceOf(user.address)).to.equal(sharesBefore + REDEEM_AMOUNT);
      expect((await queue.getRequest(standardId)).status).to.equal(3); // CANCELLED
    });

    it("returns sWATT and WATT fee on priority cancel", async () => {
      const sharesBefore = await vault.balanceOf(user.address);
      const wattBefore   = await wattUSD.balanceOf(user.address);
      await queue.connect(user).cancelRequest(priorityId);
      expect(await vault.balanceOf(user.address)).to.equal(sharesBefore + REDEEM_AMOUNT);
      expect(await wattUSD.balanceOf(user.address)).to.equal(wattBefore + minFee);
    });

    it("emits RedemptionCancelled", async () => {
      await expect(queue.connect(user).cancelRequest(standardId))
        .to.emit(queue, "RedemptionCancelled")
        .withArgs(standardId, user.address);
    });

    it("decrements queue depth", async () => {
      expect(await queue.getQueueDepth()).to.equal(2n);
      await queue.connect(user).cancelRequest(standardId);
      expect(await queue.getQueueDepth()).to.equal(1n);
    });

    it("reverts if requestId not found", async () => {
      await expect(
        queue.connect(user).cancelRequest(ethers.ZeroHash)
      ).to.be.revertedWithCustomError(queue, "RequestNotFound");
    });

    it("reverts if caller is not the request owner", async () => {
      await expect(
        queue.connect(stranger).cancelRequest(standardId)
      ).to.be.revertedWithCustomError(queue, "NotRequestOwner");
    });

    it("reverts if request is not QUEUED (already cancelled)", async () => {
      await queue.connect(user).cancelRequest(standardId);
      await expect(
        queue.connect(user).cancelRequest(standardId)
      ).to.be.revertedWithCustomError(queue, "NotQueued");
    });
  });

  // ── processBatch ──────────────────────────────────────────────────────────

  describe("processBatch", () => {
    let requestId: string;

    beforeEach(async () => {
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      const tx = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const receipt = await tx.wait();
      requestId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;
    });

    it("fulfills request: burns sWATT and delivers WATT to user", async () => {
      const wattBefore = await wattUSD.balanceOf(user.address);
      await queue.connect(processor).processBatch([requestId]);
      const wattAfter = await wattUSD.balanceOf(user.address);
      expect(wattAfter).to.be.gt(wattBefore);
      expect((await queue.getRequest(requestId)).status).to.equal(2); // FULFILLED
    });

    it("emits RedemptionFulfilled and BatchProcessed", async () => {
      await expect(queue.connect(processor).processBatch([requestId]))
        .to.emit(queue, "RedemptionFulfilled")
        .and.to.emit(queue, "BatchProcessed")
        .withArgs(1);
    });

    it("decrements queue depth", async () => {
      expect(await queue.getQueueDepth()).to.equal(1n);
      await queue.connect(processor).processBatch([requestId]);
      expect(await queue.getQueueDepth()).to.equal(0n);
    });

    it("accumulates priority fee as protocol fees on fulfillment", async () => {
      // Submit a priority request
      const minFee = REDEEM_AMOUNT * PRIORITY_FEE_BPS / 10_000n;
      await wattUSD.connect(admin).mint(user.address, minFee);
      await wattUSD.connect(user).approve(await queue.getAddress(), minFee);
      const tx = await queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee);
      const receipt = await tx.wait();
      const prioId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      expect(await queue.getProtocolFees()).to.equal(0n);
      await queue.connect(processor).processBatch([prioId]);
      expect(await queue.getProtocolFees()).to.equal(minFee);
    });

    it("processes multiple requests in one batch", async () => {
      // Create second request from user2
      const user2Shares = await vault.balanceOf(user2.address);
      await vault.connect(user2).approve(await queue.getAddress(), user2Shares);
      const tx2 = await queue.connect(user2).requestRedeem(user2Shares);
      const r2 = await tx2.wait();
      const requestId2 = (r2?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      await expect(queue.connect(processor).processBatch([requestId, requestId2]))
        .to.emit(queue, "BatchProcessed")
        .withArgs(2);
      expect(await queue.getQueueDepth()).to.equal(0n);
    });

    it("reverts on empty array", async () => {
      await expect(
        queue.connect(processor).processBatch([])
      ).to.be.revertedWithCustomError(queue, "NothingToProcess");
    });

    it("reverts on unknown requestId", async () => {
      await expect(
        queue.connect(processor).processBatch([ethers.ZeroHash])
      ).to.be.revertedWithCustomError(queue, "RequestNotFound");
    });

    it("reverts if request is already FULFILLED", async () => {
      await queue.connect(processor).processBatch([requestId]);
      await expect(
        queue.connect(processor).processBatch([requestId])
      ).to.be.revertedWithCustomError(queue, "NotQueued");
    });

    it("reverts if caller lacks PROCESSOR_ROLE", async () => {
      await expect(
        queue.connect(stranger).processBatch([requestId])
      ).to.be.revertedWithCustomError(queue, "AccessControlUnauthorizedAccount");
    });
  });

  // ── withdrawFees ─────────────────────────────────────────────────────────

  describe("withdrawFees", () => {
    let requestId: string;
    let minFee: bigint;

    beforeEach(async () => {
      // Create and fulfill a priority request to generate fees
      minFee = REDEEM_AMOUNT * PRIORITY_FEE_BPS / 10_000n;
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      await wattUSD.connect(admin).mint(user.address, minFee);
      await wattUSD.connect(user).approve(await queue.getAddress(), minFee);

      const tx = await queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee);
      const receipt = await tx.wait();
      requestId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;
      await queue.connect(processor).processBatch([requestId]);
    });

    it("transfers accumulated fees to recipient", async () => {
      const fees = await queue.getProtocolFees();
      expect(fees).to.equal(minFee);
      const before = await wattUSD.balanceOf(admin.address);
      await queue.connect(admin).withdrawFees(admin.address);
      expect(await wattUSD.balanceOf(admin.address)).to.equal(before + fees);
      expect(await queue.getProtocolFees()).to.equal(0n);
    });

    it("emits FeesWithdrawn", async () => {
      const fees = await queue.getProtocolFees();
      await expect(queue.connect(admin).withdrawFees(admin.address))
        .to.emit(queue, "FeesWithdrawn")
        .withArgs(admin.address, fees);
    });

    it("reverts when no fees available", async () => {
      await queue.connect(admin).withdrawFees(admin.address);
      await expect(
        queue.connect(admin).withdrawFees(admin.address)
      ).to.be.revertedWithCustomError(queue, "NoFeesAvailable");
    });

    it("reverts on zero address recipient", async () => {
      await expect(
        queue.connect(admin).withdrawFees(ethers.ZeroAddress)
      ).to.be.revertedWithCustomError(queue, "ZeroAddress");
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        queue.connect(stranger).withdrawFees(stranger.address)
      ).to.be.revertedWithCustomError(queue, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Queue depth ───────────────────────────────────────────────────────────

  describe("getQueueDepth", () => {
    it("tracks depth across requests, cancels, and fulfillments", async () => {
      await vault.connect(user).approve(await queue.getAddress(), userShares);

      expect(await queue.getQueueDepth()).to.equal(0n);

      const tx1 = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const r1 = await tx1.wait();
      const id1 = (r1?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      const tx2 = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const r2 = await tx2.wait();
      const id2 = (r2?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      expect(await queue.getQueueDepth()).to.equal(2n);

      await queue.connect(user).cancelRequest(id1);
      expect(await queue.getQueueDepth()).to.equal(1n);

      await queue.connect(processor).processBatch([id2]);
      expect(await queue.getQueueDepth()).to.equal(0n);
    });
  });

  // ── nextProcessingTimestamp ───────────────────────────────────────────────

  describe("nextProcessingTimestamp", () => {
    it("returns approx block.timestamp + 30 days", async () => {
      const block = await ethers.provider.getBlock("latest");
      const now = BigInt(block!.timestamp);
      const ts = await queue.nextProcessingTimestamp();
      const thirtyDays = 30n * 24n * 60n * 60n;
      // Allow ±10s tolerance for block processing time
      expect(ts).to.be.gte(now + thirtyDays - 10n);
      expect(ts).to.be.lte(now + thirtyDays + 10n);
    });
  });

  // ── Pause ─────────────────────────────────────────────────────────────────

  describe("pause", () => {
    it("blocks requestRedeem when paused", async () => {
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      await queue.connect(admin).pause();
      await expect(
        queue.connect(user).requestRedeem(REDEEM_AMOUNT)
      ).to.be.revertedWithCustomError(queue, "EnforcedPause");
    });

    it("blocks requestPriorityRedeem when paused", async () => {
      const minFee = REDEEM_AMOUNT * PRIORITY_FEE_BPS / 10_000n;
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      await wattUSD.connect(admin).mint(user.address, minFee);
      await wattUSD.connect(user).approve(await queue.getAddress(), minFee);
      await queue.connect(admin).pause();
      await expect(
        queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee)
      ).to.be.revertedWithCustomError(queue, "EnforcedPause");
    });

    it("resumes after unpause", async () => {
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      await queue.connect(admin).pause();
      await queue.connect(admin).unpause();
      await expect(
        queue.connect(user).requestRedeem(REDEEM_AMOUNT)
      ).to.not.be.reverted;
    });
  });

  // ── Upgrade ───────────────────────────────────────────────────────────────

  describe("upgrades", () => {
    it("upgrades preserving request state", async () => {
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      const tx = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const receipt = await tx.wait();
      const requestId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      const NewFactory = await ethers.getContractFactory("WEVQueue", admin);
      const upgraded = await upgrades.upgradeProxy(await queue.getAddress(), NewFactory, { kind: "uups" });
      const req = await upgraded.getRequest(requestId);
      expect(req.user).to.equal(user.address);
      expect(req.sWattAmount).to.equal(REDEEM_AMOUNT);
    });

    it("reverts upgrade if caller lacks UPGRADER_ROLE", async () => {
      const NewFactory = await ethers.getContractFactory("WEVQueue", stranger);
      await expect(
        upgrades.upgradeProxy(await queue.getAddress(), NewFactory, { kind: "uups" })
      ).to.be.reverted;
    });
  });

  // ── Integration: sWattUSD WEV guard bypass ────────────────────────────────

  describe("integration: sWattUSD WEV guard", () => {
    it("regular user cannot withdraw above threshold via sWattUSD directly", async () => {
      // User has >100k WATT deposited. Direct withdraw above threshold should revert.
      await expect(
        vault.connect(user).withdraw(WEV_THRESHOLD + 1n, user.address, user.address)
      ).to.be.revertedWithCustomError(vault, "ERC4626ExceededMaxWithdraw");
    });

    it("WEVQueue bypasses the guard and redeems large amounts for users", async () => {
      // User queues a large redemption (above threshold)
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      const tx = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const receipt = await tx.wait();
      const requestId = (receipt?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      const wattBefore = await wattUSD.balanceOf(user.address);

      // Processor fulfills — WEVQueue redeems above threshold without reverting
      await expect(
        queue.connect(processor).processBatch([requestId])
      ).to.not.be.reverted;

      // User receives WATT
      expect(await wattUSD.balanceOf(user.address)).to.be.gt(wattBefore);
    });

    it("full lifecycle: queue → cancel standard, queue priority → fulfill → withdraw fees", async () => {
      const minFee = REDEEM_AMOUNT * PRIORITY_FEE_BPS / 10_000n;
      await vault.connect(user).approve(await queue.getAddress(), userShares);
      await wattUSD.connect(admin).mint(user.address, minFee);
      await wattUSD.connect(user).approve(await queue.getAddress(), minFee);

      // Standard request then cancel
      const tx1 = await queue.connect(user).requestRedeem(REDEEM_AMOUNT);
      const r1 = await tx1.wait();
      const id1 = (r1?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;
      await queue.connect(user).cancelRequest(id1);
      expect((await queue.getRequest(id1)).status).to.equal(3); // CANCELLED

      // Priority request then fulfill
      const tx2 = await queue.connect(user).requestPriorityRedeem(REDEEM_AMOUNT, minFee);
      const r2 = await tx2.wait();
      const id2 = (r2?.logs.find((l: any) => l.fragment?.name === "RedemptionRequested") as any).args.requestId;

      const wattBefore = await wattUSD.balanceOf(user.address);
      await queue.connect(processor).processBatch([id2]);
      expect(await wattUSD.balanceOf(user.address)).to.be.gt(wattBefore);
      expect((await queue.getRequest(id2)).status).to.equal(2); // FULFILLED

      // Admin withdraws fees
      const fees = await queue.getProtocolFees();
      expect(fees).to.equal(minFee);
      await queue.connect(admin).withdrawFees(admin.address);
      expect(await queue.getProtocolFees()).to.equal(0n);
    });
  });
});
