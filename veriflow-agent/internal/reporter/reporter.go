// reporter signs telemetry payloads with HMAC-SHA256 and POSTs them to the backend.
package reporter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/neurowatt/veriflow-agent/internal/collector"
)

// Reporter sends signed telemetry to the AI WATT backend API.
type Reporter struct {
	assetID    string
	hmacKey    []byte
	backendURL string
	client     *http.Client
	logger     *zap.Logger
}

// New constructs a Reporter.
func New(assetID string, hmacKey []byte, backendURL string, logger *zap.Logger) *Reporter {
	return &Reporter{
		assetID:    assetID,
		hmacKey:    hmacKey,
		backendURL: backendURL,
		client:     &http.Client{Timeout: 30 * time.Second},
		logger:     logger,
	}
}

// unsignedPayload holds all fields that are included in the HMAC digest.
// The hmac_signature field is absent so the backend can reproduce the same JSON.
type unsignedPayload struct {
	AssetID         string  `json:"asset_id"`
	Timestamp       int64   `json:"timestamp"`
	GPUUtilization  float64 `json:"gpu_utilization"`
	GPUTemperature  float64 `json:"gpu_temperature"`
	GPUMemoryUsedMB int64   `json:"gpu_memory_used_mb"`
	GPUErrorRate    float64 `json:"gpu_error_rate"`
	ECCErrors       int64   `json:"ecc_errors"`
	PowerDrawWatts  float64 `json:"power_draw_watts"`
	FanSpeedRPM     int     `json:"fan_speed_rpm"`
	UptimePct       float64 `json:"uptime_pct"`
}

// signedPayload adds the hmac_signature field to the unsigned payload for transport.
type signedPayload struct {
	unsignedPayload
	HMACSignature string `json:"hmac_signature"`
}

// Report signs the collected metrics and POSTs them to the backend with up to 3 retries.
func (r *Reporter) Report(m *collector.Metrics) error {
	ts := time.Now().UTC().Unix()

	unsigned := unsignedPayload{
		AssetID:         r.assetID,
		Timestamp:       ts,
		GPUUtilization:  m.GPUUtilization,
		GPUTemperature:  m.GPUTemperature,
		GPUMemoryUsedMB: m.GPUMemoryUsedMB,
		GPUErrorRate:    m.GPUErrorRate,
		ECCErrors:       m.ECCErrors,
		PowerDrawWatts:  m.PowerDrawWatts,
		FanSpeedRPM:     m.FanSpeedRPM,
		UptimePct:       m.UptimePct,
	}

	// Sign the JSON bytes of the unsigned payload — must match backend verification.
	unsignedBytes, err := json.Marshal(unsigned)
	if err != nil {
		return fmt.Errorf("reporter.Report: marshal unsigned: %w", err)
	}
	mac := hmac.New(sha256.New, r.hmacKey)
	mac.Write(unsignedBytes)
	sig := hex.EncodeToString(mac.Sum(nil))

	body, err := json.Marshal(signedPayload{unsignedPayload: unsigned, HMACSignature: sig})
	if err != nil {
		return fmt.Errorf("reporter.Report: marshal signed: %w", err)
	}

	return r.postWithRetry(body)
}

// postWithRetry attempts the POST up to 3 times with exponential backoff.
func (r *Reporter) postWithRetry(body []byte) error {
	url := r.backendURL + "/api/v1/veriflow/telemetry"
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(1<<attempt) * time.Second // 2s, 4s
			r.logger.Warn("retrying telemetry POST",
				zap.Int("attempt", attempt+1),
				zap.Duration("backoff", backoff),
				zap.Error(lastErr),
			)
			time.Sleep(backoff)
		}

		resp, err := r.client.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			lastErr = fmt.Errorf("http post: %w", err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}
		lastErr = fmt.Errorf("backend returned %d", resp.StatusCode)
	}
	return fmt.Errorf("reporter.postWithRetry: all attempts failed: %w", lastErr)
}
