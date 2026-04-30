package configs

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all application configuration parsed from environment variables.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	XDC      XDCConfig
	JWT      JWTConfig
	Veriflow VeriflowConfig
	Pinata   PinataConfig
	KYC      KYCConfig
	BitGo    BitGoConfig
}

type ServerConfig struct {
	Port           string
	Env            string
	AllowedOrigins string
}

type DatabaseConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

type XDCConfig struct {
	RPCURL                   string
	ChainID                  int64
	IndexerStartBlock        uint64
	WattUSDAddress           string
	SWattUSDAddress          string
	MintEngineAddress        string
	AssetRegistryAddress     string
	OCNFTAddress             string
	HealthAttestationAddress string
	LendingPoolAddress       string
	WEVQueueAddress          string
}

type JWTConfig struct {
	Secret string
}

type VeriflowConfig struct {
	// SignerPrivateKey is the 0x-prefixed hex key used to write attestations on-chain.
	// Empty string disables on-chain writes (safe for local dev / CI without a funded wallet).
	SignerPrivateKey string
}

type PinataConfig struct {
	APIKey string
}

type KYCConfig struct {
	ProviderAPIKey string
}

type BitGoConfig struct {
	APIKey string
}

// Load reads configuration from environment variables.
// Only DATABASE_URL, REDIS_URL, and JWT_SECRET are required (will panic if missing).
// All other fields have sensible defaults for local development.
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:           getEnv("PORT", "8080"),
			Env:            getEnv("APP_ENV", "development"),
			AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
		},
		Database: DatabaseConfig{
			// go-migrate uses the mysql:// URL scheme; go-sql-driver/mysql (GORM)
			// expects the bare DSN format user:pass@tcp(host)/db?params — strip prefix.
			URL: strings.TrimPrefix(mustEnv("DATABASE_URL"), "mysql://"),
		},
		Redis: RedisConfig{
			URL: mustEnv("REDIS_URL"),
		},
		XDC: XDCConfig{
			RPCURL:                   getEnv("XDC_RPC_URL", "https://erpc.apothem.network"),
			ChainID:                  getEnvInt64("XDC_CHAIN_ID", 51), // 51=Apothem, 50=mainnet
			IndexerStartBlock:        uint64(getEnvInt64("XDC_INDEXER_START_BLOCK", 0)),
			WattUSDAddress:           getEnv("WATT_USD_PROXY_ADDRESS", ""),
			SWattUSDAddress:          getEnv("SWATT_USD_PROXY_ADDRESS", ""),
			MintEngineAddress:        getEnv("MINT_ENGINE_PROXY_ADDRESS", ""),
			AssetRegistryAddress:     getEnv("ASSET_REGISTRY_PROXY_ADDRESS", ""),
			OCNFTAddress:             getEnv("OCNFT_PROXY_ADDRESS", ""),
			HealthAttestationAddress: getEnv("HEALTH_ATTESTATION_PROXY_ADDRESS", ""),
			LendingPoolAddress:       getEnv("LENDING_POOL_PROXY_ADDRESS", ""),
			WEVQueueAddress:          getEnv("WEV_QUEUE_PROXY_ADDRESS", ""),
		},
		JWT: JWTConfig{
			Secret: mustEnv("JWT_SECRET"),
		},
		Veriflow: VeriflowConfig{
			SignerPrivateKey: getEnv("VERIFLOW_SIGNER_PRIVATE_KEY", ""),
		},
		Pinata: PinataConfig{
			APIKey: getEnv("PINATA_API_KEY", ""),
		},
		KYC: KYCConfig{
			ProviderAPIKey: getEnv("KYC_PROVIDER_API_KEY", ""),
		},
		BitGo: BitGoConfig{
			APIKey: getEnv("BITGO_API_KEY", ""),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return fallback
	}
	return n
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("required environment variable not set: " + key)
	}
	return v
}
