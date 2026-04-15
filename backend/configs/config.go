package configs

import (
	"os"
)

// Config holds all application configuration parsed from environment variables.
type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	Redis      RedisConfig
	XDC        XDCConfig
	JWT        JWTConfig
	Veriflow   VeriflowConfig
	Pinata     PinataConfig
	KYC        KYCConfig
	BitGo      BitGoConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

type XDCConfig struct {
	RPCURL            string
	ChainID           int64
	WattUSDAddress    string
	SWattUSDAddress   string
	MintEngineAddress string
	AssetRegistryAddress    string
	OCNFTAddress            string
	HealthAttestationAddress string
	LendingPoolAddress      string
	WEVQueueAddress         string
}

type JWTConfig struct {
	Secret string
}

type VeriflowConfig struct {
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
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("APP_ENV", "development"),
		},
		Database: DatabaseConfig{
			URL: mustEnv("DATABASE_URL"),
		},
		Redis: RedisConfig{
			URL: mustEnv("REDIS_URL"),
		},
		XDC: XDCConfig{
			RPCURL:                   getEnv("XDC_RPC_URL", "https://erpc.apothem.network"),
			ChainID:                  51, // Apothem testnet; override with XDC_CHAIN_ID for mainnet (50)
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
			SignerPrivateKey: mustEnv("VERIFLOW_SIGNER_PRIVATE_KEY"),
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

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("required environment variable not set: " + key)
	}
	return v
}

