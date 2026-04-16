// veriflow-agent — standalone Go binary deployed on borrower hardware as a systemd service.
// Reads GPU/system metrics and POSTs signed telemetry to the AI WATT backend.
//
// Build: CGO_ENABLED=0 go build -o dist/veriflow-agent-linux-amd64 ./cmd
package main

import (
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/neurowatt/veriflow-agent/internal/collector"
	"github.com/neurowatt/veriflow-agent/internal/reporter"
)

type config struct {
	AssetID        string
	HMACKey        string
	BackendURL     string
	ReportInterval time.Duration
}

func loadConfig() config {
	mustEnv := func(key string) string {
		v := os.Getenv(key)
		if v == "" {
			// Use a temporary stdlib logger — zap not yet initialised.
			panic("required env var not set: " + key)
		}
		return v
	}
	interval := 5 * time.Minute
	if s := os.Getenv("REPORT_INTERVAL"); s != "" {
		if secs, err := strconv.Atoi(s); err == nil && secs > 0 {
			interval = time.Duration(secs) * time.Second
		}
	}
	return config{
		AssetID:        mustEnv("ASSET_ID"),
		HMACKey:        mustEnv("HMAC_KEY"),
		BackendURL:     mustEnv("BACKEND_URL"),
		ReportInterval: interval,
	}
}

func main() {
	cfg := loadConfig()

	log, _ := zap.NewProduction()
	defer log.Sync()

	c := collector.New()
	r := reporter.New(cfg.AssetID, []byte(cfg.HMACKey), cfg.BackendURL, log)

	log.Info("veriflow-agent starting",
		zap.String("asset_id", cfg.AssetID),
		zap.String("backend_url", cfg.BackendURL),
		zap.Duration("report_interval", cfg.ReportInterval),
	)

	for {
		metrics, err := c.Collect()
		if err != nil {
			log.Error("collect failed", zap.Error(err))
		} else if err := r.Report(metrics); err != nil {
			log.Error("report failed", zap.Error(err))
		} else {
			log.Info("telemetry sent", zap.String("asset_id", cfg.AssetID))
		}
		time.Sleep(cfg.ReportInterval)
	}
}
