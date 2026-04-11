// veriflow-agent — standalone Go binary deployed on borrower hardware as a systemd service.
// Reads GPU/system metrics and POSTs signed telemetry to the AI WATT backend every 5 minutes.
//
// Build: CGO_ENABLED=0 go build -o veriflow-agent ./cmd
package main

import (
	"log"
	"os"
	"time"

	"github.com/neurowatt/veriflow-agent/internal/collector"
	"github.com/neurowatt/veriflow-agent/internal/reporter"
)

func main() {
	assetID := os.Getenv("ASSET_ID")
	if assetID == "" {
		log.Fatal("ASSET_ID environment variable is required")
	}
	hmacKey := os.Getenv("HMAC_KEY")
	if hmacKey == "" {
		log.Fatal("HMAC_KEY environment variable is required")
	}
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		log.Fatal("BACKEND_URL environment variable is required")
	}

	c := collector.New()
	r := reporter.New(assetID, []byte(hmacKey), backendURL)

	interval := 5 * time.Minute
	log.Printf("veriflow-agent starting: asset=%s interval=%s", assetID, interval)

	for {
		metrics, err := c.Collect()
		if err != nil {
			log.Printf("collect error: %v", err)
		} else if err := r.Report(metrics); err != nil {
			log.Printf("report error: %v", err)
		} else {
			log.Printf("telemetry sent: asset=%s", assetID)
		}
		time.Sleep(interval)
	}
}
