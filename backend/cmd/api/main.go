package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	api "github.com/neurowatt/aiwatt-backend/internal/api"
	"github.com/neurowatt/aiwatt-backend/internal/api/handler"
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

	// ── Services (DI wiring — wire real implementations here) ────────────────
	userSvc := service.NewUserService()
	loanSvc := service.NewLoanService()
	mintSvc := service.NewMintService()
	assetSvc := service.NewAssetService()
	veriflowSvc := service.NewVeriflowService()

	// ── Handlers ─────────────────────────────────────────────────────────────
	deps := api.RouterDeps{
		JWTSecret:       cfg.JWT.Secret,
		Logger:          log,
		UserHandler:     handler.NewUserHandler(userSvc, log),
		LoanHandler:     handler.NewLoanHandler(loanSvc, log),
		MintHandler:     handler.NewMintHandler(mintSvc, log),
		AssetHandler:    handler.NewAssetHandler(assetSvc, log),
		VeriflowHandler: handler.NewVeriflowHandler(veriflowSvc, log),
	}

	router := api.NewRouter(deps)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// ── Graceful shutdown ─────────────────────────────────────────────────────
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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
