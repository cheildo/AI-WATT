package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/api/handler"
	"github.com/neurowatt/aiwatt-backend/internal/api/middleware"
)

// RouterDeps holds all handler dependencies injected at startup.
type RouterDeps struct {
	JWTSecret       string
	AllowedOrigins  string // comma-separated or "*"
	Logger          *zap.Logger
	RedisClient     *redis.Client
	UserHandler     *handler.UserHandler
	LoanHandler     *handler.LoanHandler
	MintHandler     *handler.MintHandler
	AssetHandler    *handler.AssetHandler
	VeriflowHandler *handler.VeriflowHandler
	WEVHandler      *handler.WEVHandler
}

// NewRouter builds and returns the Gin engine with all routes registered.
func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS(deps.AllowedOrigins))
	r.Use(middleware.RequestLogger(deps.Logger))

	v1 := r.Group("/api/v1")

	// Public routes
	auth := v1.Group("/auth")
	{
		auth.POST("/register", deps.UserHandler.Register)
		auth.POST("/login", deps.UserHandler.Login)
		auth.POST("/wallet-login", deps.UserHandler.WalletLogin)
	}

	// Veriflow telemetry ingestion (authenticated by HMAC, not JWT)
	v1.POST("/veriflow/telemetry", deps.VeriflowHandler.IngestTelemetry)

	// Protected routes
	protected := v1.Group("/")
	protected.Use(middleware.JWTAuth(deps.JWTSecret))
	protected.Use(middleware.RateLimit(deps.RedisClient, 60, 60*time.Second))
	{
		// Users
		users := protected.Group("/users")
		{
			users.GET("/:id", deps.UserHandler.GetByID)
			users.PATCH("/:id", deps.UserHandler.Update)
		}

		// Loans
		loans := protected.Group("/loans")
		{
			loans.GET("", deps.LoanHandler.List)
			loans.POST("", deps.LoanHandler.Create)
			loans.GET("/:id", deps.LoanHandler.GetByID)
			loans.PATCH("/:id", middleware.RequireRole("curator", "admin"), deps.LoanHandler.Update)
		}

		// Assets
		assets := protected.Group("/assets")
		{
			assets.GET("", deps.AssetHandler.List)
			assets.POST("", middleware.RequireRole("admin", "curator"), deps.AssetHandler.Register)
			assets.GET("/:id", deps.AssetHandler.GetByID)
			assets.GET("/:assetId/health", deps.AssetHandler.GetHealth)
			assets.PATCH("/:id/ltv", middleware.RequireRole("admin"), deps.AssetHandler.UpdateLTV)
		}

		// Mint / Redeem
		mint := protected.Group("/mint")
		{
			mint.POST("", deps.MintHandler.Mint)
			mint.POST("/redeem", deps.MintHandler.Redeem)
			mint.GET("/nav", deps.MintHandler.GetNAV)
		}

		// Vault stats
		protected.GET("/vault/stats", deps.MintHandler.GetVaultStats)

		// WEV redemption queue
		wev := protected.Group("/wev")
		{
			wev.POST("/redeem", deps.WEVHandler.RequestRedeem)
			wev.DELETE("/redeem/:requestId", deps.WEVHandler.CancelRedeem)
			wev.GET("/queue", deps.WEVHandler.GetQueue)
			wev.GET("/queue/me", deps.WEVHandler.GetMyQueue)
		}

		// Veriflow (health scores, attestations)
		veriflow := protected.Group("/veriflow")
		{
			veriflow.GET("/assets/:id/score", deps.VeriflowHandler.GetHealthScore)
			veriflow.GET("/assets/:id/attestation", deps.VeriflowHandler.GetAttestation)
		}
	}

	return r
}
