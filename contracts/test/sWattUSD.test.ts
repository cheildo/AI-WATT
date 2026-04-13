import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { sWattUSD, MockERC20 } from "../typechain-types";

const DECIMALS = 6n;
const ONE_WATT = 10n ** DECIMALS;
const HUNDRED_K_WATT = 100_000n * ONE_WATT;

describe("sWattUSD", () => {
  let vault: sWattUSD;
  let watt: MockERC20;

  let admin: HardhatEthersSigner;
  let alice: HardhatEthersSigner;
  let bob: HardhatEthersSigner;
  let yieldDistributor: HardhatEthersSigner;
  let fakeQueue: HardhatEthersSigner;

  const YIELD_DISTRIBUTOR_ROLE = ethers.keccak256(
    ethers.toUtf8Bytes("YIELD_DISTRIBUTOR_ROLE")
  );
  const PAUSER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("PAUSER_ROLE"));
  const ADMIN_ROLE = ethers.keccak256(ethers.toUtf8Bytes("ADMIN_ROLE"));
  const UPGRADER_ROLE = ethers.keccak256(ethers.toUtf8Bytes("UPGRADER_ROLE"));

  beforeEach(async () => {
    [admin, alice, bob, yieldDistributor, fakeQueue] =
      await ethers.getSigners();

    // Deploy mock WATT token
    const MockERC20Factory = await ethers.getContractFactory("MockERC20");
    watt = (await MockERC20Factory.deploy("WattUSD", "WATT", 6)) as MockERC20;

    // Deploy sWattUSD proxy
    const sWattFactory = await ethers.getContractFactory("sWattUSD");
    vault = (await upgrades.deployProxy(
      sWattFactory,
      [admin.address, await watt.getAddress()],
      { initializer: "initialize", kind: "uups" }
    )) as unknown as sWattUSD;

    // Grant yield distributor role
    await vault
      .connect(admin)
      .grantRole(YIELD_DISTRIBUTOR_ROLE, yieldDistributor.address);

    // Mint WATT to participants
    for (const signer of [admin, alice, bob, yieldDistributor]) {
      await watt.mint(signer.address, 10_000_000n * ONE_WATT);
    }
  });

  // ── Deployment ─────────────────────────────────────────────────────────────

  describe("deployment", () => {
    it("sets ERC-20 name and symbol", async () => {
      expect(await vault.name()).to.equal("Staked WattUSD");
      expect(await vault.symbol()).to.equal("sWATT");
    });

    it("sets underlying asset to WATT", async () => {
      expect(await vault.asset()).to.equal(await watt.getAddress());
    });

    it("reports 6 decimals", async () => {
      expect(await vault.decimals()).to.equal(6n);
    });

    it("sets default wevThreshold to 100,000 WATT", async () => {
      expect(await vault.wevThreshold()).to.equal(HUNDRED_K_WATT);
    });

    it("sets wevQueue to address(0)", async () => {
      expect(await vault.wevQueue()).to.equal(ethers.ZeroAddress);
    });

    it("grants all roles to admin", async () => {
      expect(
        await vault.hasRole(ethers.ZeroHash, admin.address)
      ).to.be.true;
      expect(await vault.hasRole(ADMIN_ROLE, admin.address)).to.be.true;
      expect(await vault.hasRole(PAUSER_ROLE, admin.address)).to.be.true;
      expect(await vault.hasRole(UPGRADER_ROLE, admin.address)).to.be.true;
    });

    it("reverts if admin is zero address", async () => {
      const Factory = await ethers.getContractFactory("sWattUSD");
      await expect(
        upgrades.deployProxy(
          Factory,
          [ethers.ZeroAddress, await watt.getAddress()],
          { initializer: "initialize", kind: "uups" }
        )
      ).to.be.revertedWithCustomError(vault, "ZeroAddress");
    });

    it("reverts if watt is zero address", async () => {
      const Factory = await ethers.getContractFactory("sWattUSD");
      await expect(
        upgrades.deployProxy(
          Factory,
          [admin.address, ethers.ZeroAddress],
          { initializer: "initialize", kind: "uups" }
        )
      ).to.be.revertedWithCustomError(vault, "ZeroAddress");
    });
  });

  // ── Deposit ────────────────────────────────────────────────────────────────

  describe("deposit", () => {
    it("mints shares 1:1 on first deposit (empty vault)", async () => {
      const depositAmount = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), depositAmount);
      await vault.connect(alice).deposit(depositAmount, alice.address);

      // First deposit: shares == assets (1:1 before any yield)
      expect(await vault.balanceOf(alice.address)).to.equal(depositAmount);
      expect(await vault.totalAssets()).to.equal(depositAmount);
    });

    it("emits Deposit event", async () => {
      const amount = 500n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), amount);
      await expect(
        vault.connect(alice).deposit(amount, alice.address)
      ).to.emit(vault, "Deposit");
    });

    it("transfers WATT from depositor to vault", async () => {
      const amount = 1_000n * ONE_WATT;
      const before = await watt.balanceOf(alice.address);
      await watt.connect(alice).approve(await vault.getAddress(), amount);
      await vault.connect(alice).deposit(amount, alice.address);
      expect(await watt.balanceOf(alice.address)).to.equal(before - amount);
      expect(await watt.balanceOf(await vault.getAddress())).to.equal(amount);
    });

    it("second depositor gets fewer shares after yield injection", async () => {
      // Alice deposits first
      const aliceDeposit = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), aliceDeposit);
      await vault.connect(alice).deposit(aliceDeposit, alice.address);

      // Inject yield — doubles total assets
      const yieldAmount = 1_000n * ONE_WATT;
      await watt
        .connect(yieldDistributor)
        .approve(await vault.getAddress(), yieldAmount);
      await vault.connect(yieldDistributor).receiveYield(yieldAmount);

      // Bob deposits same WATT — should get half the shares Alice got
      const bobDeposit = 1_000n * ONE_WATT;
      await watt.connect(bob).approve(await vault.getAddress(), bobDeposit);
      await vault.connect(bob).deposit(bobDeposit, bob.address);

      const aliceShares = await vault.balanceOf(alice.address);
      const bobShares = await vault.balanceOf(bob.address);
      // Bob gets ~half the shares Alice got
      expect(bobShares).to.be.approximately(aliceShares / 2n, ONE_WATT);
    });
  });

  // ── NAV per share ──────────────────────────────────────────────────────────

  describe("navPerShare", () => {
    it("returns 1 WATT per sWATT initially (empty vault)", async () => {
      // No deposits yet — ERC-4626 virtual shares: convertToAssets(10^6) = 10^6
      expect(await vault.navPerShare()).to.equal(ONE_WATT);
    });

    it("increases after yield injection", async () => {
      const depositAmount = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), depositAmount);
      await vault.connect(alice).deposit(depositAmount, alice.address);

      const navBefore = await vault.navPerShare();

      const yieldAmount = 200n * ONE_WATT; // 20% yield
      await watt
        .connect(yieldDistributor)
        .approve(await vault.getAddress(), yieldAmount);
      await vault.connect(yieldDistributor).receiveYield(yieldAmount);

      const navAfter = await vault.navPerShare();
      expect(navAfter).to.be.gt(navBefore);
      // NAV should be ~1.2 WATT per sWATT
      expect(navAfter).to.be.approximately(
        (12n * ONE_WATT) / 10n,
        1000n // 0.001 WATT tolerance for rounding
      );
    });

    it("does not change on deposit (only yield moves NAV)", async () => {
      const deposit = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), deposit);
      await vault.connect(alice).deposit(deposit, alice.address);
      const navAfterAlice = await vault.navPerShare();

      await watt.connect(bob).approve(await vault.getAddress(), deposit);
      await vault.connect(bob).deposit(deposit, bob.address);
      const navAfterBob = await vault.navPerShare();

      // NAV per share is unchanged by deposits (both get proportional shares)
      expect(navAfterBob).to.equal(navAfterAlice);
    });
  });

  // ── receiveYield ───────────────────────────────────────────────────────────

  describe("receiveYield", () => {
    it("increases totalAssets by the yield amount", async () => {
      const deposit = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), deposit);
      await vault.connect(alice).deposit(deposit, alice.address);

      const before = await vault.totalAssets();
      const yieldAmount = 100n * ONE_WATT;
      await watt
        .connect(yieldDistributor)
        .approve(await vault.getAddress(), yieldAmount);
      await vault.connect(yieldDistributor).receiveYield(yieldAmount);

      expect(await vault.totalAssets()).to.equal(before + yieldAmount);
    });

    it("emits YieldReceived event", async () => {
      const deposit = 500n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), deposit);
      await vault.connect(alice).deposit(deposit, alice.address);

      const yieldAmount = 50n * ONE_WATT;
      await watt
        .connect(yieldDistributor)
        .approve(await vault.getAddress(), yieldAmount);
      await expect(
        vault.connect(yieldDistributor).receiveYield(yieldAmount)
      )
        .to.emit(vault, "YieldReceived")
        .withArgs(yieldAmount, deposit + yieldAmount);
    });

    it("reverts if caller lacks YIELD_DISTRIBUTOR_ROLE", async () => {
      const yieldAmount = 100n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), yieldAmount);
      await expect(
        vault.connect(alice).receiveYield(yieldAmount)
      ).to.be.revertedWithCustomError(vault, "AccessControlUnauthorizedAccount");
    });

    it("reverts on zero amount", async () => {
      await expect(
        vault.connect(yieldDistributor).receiveYield(0n)
      ).to.be.revertedWithCustomError(vault, "ZeroAmount");
    });
  });

  // ── Withdraw / Redeem ──────────────────────────────────────────────────────

  describe("withdraw and redeem", () => {
    beforeEach(async () => {
      const deposit = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), deposit);
      await vault.connect(alice).deposit(deposit, alice.address);
    });

    it("allows full withdraw when no queue is set", async () => {
      const shares = await vault.balanceOf(alice.address);
      const assets = await vault.convertToAssets(shares);

      await vault.connect(alice).redeem(shares, alice.address, alice.address);
      expect(await vault.balanceOf(alice.address)).to.equal(0n);
      expect(await watt.balanceOf(alice.address)).to.be.gte(
        10_000_000n * ONE_WATT - 1n // back to near-starting balance
      );
    });

    it("emits Withdraw event on redeem", async () => {
      const shares = await vault.balanceOf(alice.address);
      await expect(
        vault.connect(alice).redeem(shares, alice.address, alice.address)
      ).to.emit(vault, "Withdraw");
    });

    it("redeems correct WATT amount after yield injection", async () => {
      const yield1 = 500n * ONE_WATT; // 50% yield on 1000 WATT
      await watt
        .connect(yieldDistributor)
        .approve(await vault.getAddress(), yield1);
      await vault.connect(yieldDistributor).receiveYield(yield1);

      const shares = await vault.balanceOf(alice.address);
      const wattBefore = await watt.balanceOf(alice.address);
      await vault.connect(alice).redeem(shares, alice.address, alice.address);
      const wattAfter = await watt.balanceOf(alice.address);

      // Alice should get ~1500 WATT back (deposit + 50% yield). ERC-4626 rounds down by 1 wei.
      expect(wattAfter - wattBefore).to.be.gte(1_500n * ONE_WATT - 1n);
      expect(wattAfter - wattBefore).to.be.lte(1_500n * ONE_WATT);
    });
  });

  // ── WEV threshold guard ────────────────────────────────────────────────────

  describe("WEV threshold guard", () => {
    const bigDeposit = 500_000n * ONE_WATT; // 500,000 WATT

    beforeEach(async () => {
      await watt.connect(alice).approve(await vault.getAddress(), bigDeposit);
      await vault.connect(alice).deposit(bigDeposit, alice.address);
      // Activate the queue guard
      await vault.connect(admin).setWEVQueue(fakeQueue.address);
    });

    it("maxWithdraw is capped at wevThreshold when queue is active", async () => {
      const max = await vault.maxWithdraw(alice.address);
      expect(max).to.equal(HUNDRED_K_WATT);
    });

    it("maxRedeem is capped at share equivalent of wevThreshold", async () => {
      const maxR = await vault.maxRedeem(alice.address);
      const thresholdShares = await vault.convertToShares(HUNDRED_K_WATT);
      expect(maxR).to.equal(thresholdShares);
    });

    it("reverts withdraw above threshold (ERC-4626 maxWithdraw check fires first)", async () => {
      // maxWithdraw() is capped at wevThreshold, so ERC-4626's standard guard fires
      // before _withdraw is reached. The descriptive LargeRedemptionUseWEVQueue error
      // is a defense-in-depth fallback for direct _withdraw calls bypassing the standard path.
      const largeAmount = HUNDRED_K_WATT + ONE_WATT;
      await expect(
        vault.connect(alice).withdraw(largeAmount, alice.address, alice.address)
      ).to.be.revertedWithCustomError(vault, "ERC4626ExceededMaxWithdraw");
    });

    it("allows withdrawal at or below threshold when queue is active", async () => {
      const atThreshold = HUNDRED_K_WATT;
      const wattBefore = await watt.balanceOf(alice.address);
      await vault
        .connect(alice)
        .withdraw(atThreshold, alice.address, alice.address);
      expect(await watt.balanceOf(alice.address)).to.equal(
        wattBefore + atThreshold
      );
    });

    it("allows full withdrawal when queue is cleared (address(0))", async () => {
      await vault.connect(admin).setWEVQueue(ethers.ZeroAddress);
      const shares = await vault.balanceOf(alice.address);
      // Should not revert
      await expect(
        vault.connect(alice).redeem(shares, alice.address, alice.address)
      ).to.not.be.reverted;
    });
  });

  // ── setWEVQueue ────────────────────────────────────────────────────────────

  describe("setWEVQueue", () => {
    it("updates wevQueue and emits WEVQueueUpdated", async () => {
      await expect(vault.connect(admin).setWEVQueue(fakeQueue.address))
        .to.emit(vault, "WEVQueueUpdated")
        .withArgs(ethers.ZeroAddress, fakeQueue.address);
      expect(await vault.wevQueue()).to.equal(fakeQueue.address);
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        vault.connect(alice).setWEVQueue(fakeQueue.address)
      ).to.be.revertedWithCustomError(vault, "AccessControlUnauthorizedAccount");
    });

    it("allows setting queue to address(0) to disable guard", async () => {
      await vault.connect(admin).setWEVQueue(fakeQueue.address);
      await vault.connect(admin).setWEVQueue(ethers.ZeroAddress);
      expect(await vault.wevQueue()).to.equal(ethers.ZeroAddress);
    });
  });

  // ── setWEVThreshold ────────────────────────────────────────────────────────

  describe("setWEVThreshold", () => {
    it("updates threshold and emits WEVThresholdUpdated", async () => {
      const newThreshold = 50_000n * ONE_WATT;
      await expect(vault.connect(admin).setWEVThreshold(newThreshold))
        .to.emit(vault, "WEVThresholdUpdated")
        .withArgs(HUNDRED_K_WATT, newThreshold);
      expect(await vault.wevThreshold()).to.equal(newThreshold);
    });

    it("reverts on zero threshold", async () => {
      await expect(
        vault.connect(admin).setWEVThreshold(0n)
      ).to.be.revertedWithCustomError(vault, "ZeroAmount");
    });

    it("reverts if caller lacks ADMIN_ROLE", async () => {
      await expect(
        vault.connect(alice).setWEVThreshold(50_000n * ONE_WATT)
      ).to.be.revertedWithCustomError(vault, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Pause ──────────────────────────────────────────────────────────────────

  describe("pause / unpause", () => {
    it("blocks deposits when paused", async () => {
      await vault.connect(admin).pause();
      const amount = 100n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), amount);
      await expect(
        vault.connect(alice).deposit(amount, alice.address)
      ).to.be.revertedWithCustomError(vault, "EnforcedPause");
    });

    it("blocks withdrawals when paused", async () => {
      const amount = 100n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), amount);
      await vault.connect(alice).deposit(amount, alice.address);

      await vault.connect(admin).pause();
      // maxWithdraw returns 0 when paused, so ERC-4626's standard guard fires first
      await expect(
        vault.connect(alice).withdraw(amount, alice.address, alice.address)
      ).to.be.revertedWithCustomError(vault, "ERC4626ExceededMaxWithdraw");
    });

    it("maxWithdraw returns 0 when paused", async () => {
      const amount = 100n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), amount);
      await vault.connect(alice).deposit(amount, alice.address);

      await vault.connect(admin).pause();
      expect(await vault.maxWithdraw(alice.address)).to.equal(0n);
      expect(await vault.maxRedeem(alice.address)).to.equal(0n);
    });

    it("resumes after unpause", async () => {
      const amount = 100n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), amount);
      await vault.connect(alice).deposit(amount, alice.address);

      await vault.connect(admin).pause();
      await vault.connect(admin).unpause();

      const shares = await vault.balanceOf(alice.address);
      await expect(
        vault.connect(alice).redeem(shares, alice.address, alice.address)
      ).to.not.be.reverted;
    });

    it("reverts pause if caller lacks PAUSER_ROLE", async () => {
      await expect(
        vault.connect(alice).pause()
      ).to.be.revertedWithCustomError(vault, "AccessControlUnauthorizedAccount");
    });
  });

  // ── Upgrades ───────────────────────────────────────────────────────────────

  describe("upgrades", () => {
    it("upgrades to new implementation preserving state", async () => {
      const deposit = 1_000n * ONE_WATT;
      await watt.connect(alice).approve(await vault.getAddress(), deposit);
      await vault.connect(alice).deposit(deposit, alice.address);

      const yieldAmount = 100n * ONE_WATT;
      await watt
        .connect(yieldDistributor)
        .approve(await vault.getAddress(), yieldAmount);
      await vault.connect(yieldDistributor).receiveYield(yieldAmount);

      const navBefore = await vault.navPerShare();
      const sharesBefore = await vault.balanceOf(alice.address);

      // Upgrade to new implementation (same contract = storage layout check)
      const NewFactory = await ethers.getContractFactory("sWattUSD", admin);
      const upgraded = await upgrades.upgradeProxy(
        await vault.getAddress(),
        NewFactory,
        { kind: "uups" }
      );

      expect(await upgraded.navPerShare()).to.equal(navBefore);
      expect(await upgraded.balanceOf(alice.address)).to.equal(sharesBefore);
      expect(await upgraded.wevThreshold()).to.equal(HUNDRED_K_WATT);
    });

    it("reverts upgrade if caller lacks UPGRADER_ROLE", async () => {
      const NewFactory = await ethers.getContractFactory("sWattUSD", alice);
      await expect(
        upgrades.upgradeProxy(await vault.getAddress(), NewFactory, {
          kind: "uups",
        })
      ).to.be.reverted;
    });
  });
});
