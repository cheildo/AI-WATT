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

	"github.com/neurowatt/veriflow-agent/internal/collector"
)

// Reporter sends signed telemetry to the AI WATT backend API.
type Reporter struct {
	assetID    string
	hmacKey    []byte
	backendURL string
	client     *http.Client
}

// New constructs a Reporter.
func New(assetID string, hmacKey []byte, backendURL string) *Reporter {
	return &Reporter{
		assetID:    assetID,
		hmacKey:    hmacKey,
		backendURL: backendURL,
		client:     &http.Client{Timeout: 30 * time.Second},
	}
}

type payload struct {
	AssetID         string  `json:"asset_id"`
	Timestamp       int64   `json:"timestamp"`
	HMACSignature   string  `json:"hmac_signature"`
	GPUUtilization  float64 `json:"gpu_utilization"`
	GPUTemperature  float64 `json:"gpu_temperature"`
	GPUMemoryUsedMB int64   `json:"gpu_memory_used_mb"`
	GPUErrorRate    float64 `json:"gpu_error_rate"`
	PowerDrawWatts  float64 `json:"power_draw_watts"`
	FanSpeedRPM     int     `json:"fan_speed_rpm"`
}

// Report signs the collected metrics and POSTs them to the backend.
func (r *Reporter) Report(m *collector.Metrics) error {
	ts := time.Now().UTC().Unix()
	raw := fmt.Sprintf("%s:%d:%.2f:%.2f:%d:%.6f:%.2f:%d",
		r.assetID, ts,
		m.GPUUtilization, m.GPUTemperature, m.GPUMemoryUsedMB,
		m.GPUErrorRate, m.PowerDrawWatts, m.FanSpeedRPM,
	)
	mac := hmac.New(sha256.New, r.hmacKey)
	mac.Write([]byte(raw))
	sig := hex.EncodeToString(mac.Sum(nil))

	p := payload{
		AssetID:         r.assetID,
		Timestamp:       ts,
		HMACSignature:   sig,
		GPUUtilization:  m.GPUUtilization,
		GPUTemperature:  m.GPUTemperature,
		GPUMemoryUsedMB: m.GPUMemoryUsedMB,
		GPUErrorRate:    m.GPUErrorRate,
		PowerDrawWatts:  m.PowerDrawWatts,
		FanSpeedRPM:     m.FanSpeedRPM,
	}
	body, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("reporter.Report: marshal: %w", err)
	}
	resp, err := r.client.Post(r.backendURL+"/api/v1/veriflow/telemetry", "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("reporter.Report: post: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("reporter.Report: backend returned %d", resp.StatusCode)
	}
	return nil
}
