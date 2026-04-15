package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	api "github.com/neurowatt/aiwatt-backend/internal/api"
	"github.com/neurowatt/aiwatt-backend/internal/api/handler"
	"github.com/neurowatt/aiwatt-backend/internal/blockchain"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/configs"
	"github.com/neurowatt/aiwatt-backend/pkg/logger"
)

// @title           AI WATT API
// @version         1.0
// @description     Decentralized credit protocol for AI and automation assets — XDC Network
// @termsOfService  https://aiwatt.io/terms

// @contact.name   Neurowatt Engineering
// @contact.email  engineering@neurowatt.io

// @license.name  Proprietary
// @license.url   https://neurowatt.io

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cfg := configs.Load()

	logger.Init(cfg.Server.Env)
	defer logger.Sync()
	log := logger.L()

	// ── Database ──────────────────────────────────────────────────────────────
	db, err := gorm.Open(mysql.Open(cfg.Database.URL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get underlying sql.DB", zap.Error(err))
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// ── Redis ─────────────────────────────────────────────────────────────────
	redisOpts, err := redis.ParseURL(cfg.Redis.URL)
	if err != nil {
		log.Fatal("failed to parse REDIS_URL", zap.Error(err))
	}
	rdb := redis.NewClient(redisOpts)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("failed to connect to Redis", zap.Error(err))
	}

	// ── Blockchain ────────────────────────────────────────────────────────────
	addrs := blockchain.ContractAddresses{
		WattUSD:           common.HexToAddress(cfg.XDC.WattUSDAddress),
		SWattUSD:          common.HexToAddress(cfg.XDC.SWattUSDAddress),
		MintEngine:        common.HexToAddress(cfg.XDC.MintEngineAddress),
		AssetRegistry:     common.HexToAddress(cfg.XDC.AssetRegistryAddress),
		OCNFT:             common.HexToAddress(cfg.XDC.OCNFTAddress),
		HealthAttestation: common.HexToAddress(cfg.XDC.HealthAttestationAddress),
		LendingPool:       common.HexToAddress(cfg.XDC.LendingPoolAddress),
		WEVQueue:          common.HexToAddress(cfg.XDC.WEVQueueAddress),
	}
	bcClient, err := blockchain.NewBlockchainClient(cfg.XDC.RPCURL, addrs)
	if err != nil {
		log.Fatal("failed to connect to XDC node", zap.Error(err))
	}
	defer bcClient.Close()

	txManager, err := blockchain.NewTxManager(bcClient, cfg.Veriflow.SignerPrivateKey, cfg.XDC.ChainID, rdb, log)
	if err != nil {
		log.Fatal("failed to create TxManager", zap.Error(err))
	}

	// ── Repositories ──────────────────────────────────────────────────────────
	userRepo        := repository.NewUserRepository(db)
	loanRepo        := repository.NewLoanRepository(db)
	assetRepo       := repository.NewAssetRepository(db)
	telemetryRepo   := repository.NewTelemetryRepository(db)
	attestationRepo := repository.NewAttestationRepository(db)
	wevRepo         := repository.NewWEVRepository(db)
	eventRepo       := repository.NewEventRepository(db)

	// ── Event Indexer ─────────────────────────────────────────────────────────
	indexer, err := blockchain.NewEventIndexer(bcClient, eventRepo, rdb, log)
	if err != nil {
		log.Fatal("failed to create EventIndexer", zap.Error(err))
	}

	// ── Services ──────────────────────────────────────────────────────────────
	userSvc     := service.NewUserService(userRepo, cfg.JWT.Secret)
	loanSvc     := service.NewLoanService(loanRepo, assetRepo, bcClient)
	assetSvc    := service.NewAssetService(assetRepo, txManager, log)
	mintSvc     := service.NewMintService()
	yieldSvc    := service.NewYieldService()
	wevSvc      := service.NewWEVService(wevRepo)
	veriflowSvc := service.NewVeriflowService(telemetryRepo, attestationRepo, assetRepo)
	notifySvc   := service.NewNotifyService(log)

	_ = notifySvc // wired in Phase 9 when events trigger notifications

	// ── Handlers ─────────────────────────────────────────────────────────────
	allowedOrigins := getEnvOrDefault("ALLOWED_ORIGINS", "*")

	deps := api.RouterDeps{
		JWTSecret:      cfg.JWT.Secret,
		AllowedOrigins: allowedOrigins,
		Logger:         log,
		RedisClient:    rdb,
		UserHandler:    handler.NewUserHandler(userSvc, log),
		LoanHandler:    handler.NewLoanHandler(loanSvc, log),
		MintHandler:    handler.NewMintHandler(mintSvc, yieldSvc, log),
		AssetHandler:   handler.NewAssetHandler(assetSvc, log),
		VeriflowHandler: handler.NewVeriflowHandler(veriflowSvc, log),
		WEVHandler:     handler.NewWEVHandler(wevSvc, log),
	}

	router := api.NewRouter(deps)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// ── Graceful shutdown ─────────────────────────────────────────────────────
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start event indexer in background.
	go func() {
		if err := indexer.Start(ctx); err != nil && ctx.Err() == nil {
			log.Error("EventIndexer stopped unexpectedly", zap.Error(err))
		}
	}()

	go func() {
		log.Info("server starting", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed", zap.Error(err))
		}
	}()

	<-ctx.Done()
	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("graceful shutdown failed", zap.Error(err))
	}
	log.Info("server stopped")
}

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
