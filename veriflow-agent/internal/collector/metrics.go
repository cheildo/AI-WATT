// Package collector reads GPU and system health metrics from nvidia-smi and ipmitool.
package collector

// Metrics holds the raw hardware telemetry values collected from the machine.
type Metrics struct {
	GPUUtilization  float64
	GPUTemperature  float64
	GPUMemoryUsedMB int64
	// GPUErrorRate is the ECC error rate fraction (ECCErrors / 10000).
	GPUErrorRate   float64
	ECCErrors      int64
	PowerDrawWatts float64
	FanSpeedRPM    int
	// UptimePct is the system uptime expressed as a percentage over a 7-day baseline.
	UptimePct float64
}
