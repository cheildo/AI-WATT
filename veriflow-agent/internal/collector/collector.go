// collector reads GPU and system health metrics from nvidia-smi and ipmitool.
package collector

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Metrics holds the raw hardware telemetry values collected from the machine.
type Metrics struct {
	GPUUtilization  float64
	GPUTemperature  float64
	GPUMemoryUsedMB int64
	GPUErrorRate    float64
	PowerDrawWatts  float64
	FanSpeedRPM     int
}

// Collector reads hardware metrics from local tools.
type Collector struct{}

// New constructs a Collector.
func New() *Collector {
	return &Collector{}
}

// Collect reads metrics from nvidia-smi and ipmitool.
func (c *Collector) Collect() (*Metrics, error) {
	m := &Metrics{}
	if err := c.collectGPU(m); err != nil {
		return nil, fmt.Errorf("collector.Collect: gpu: %w", err)
	}
	// ipmitool collection is best-effort — non-GPU hardware may not have it
	_ = c.collectIPMI(m)
	return m, nil
}

// collectGPU runs nvidia-smi and parses key metrics.
func (c *Collector) collectGPU(m *Metrics) error {
	out, err := exec.Command(
		"nvidia-smi",
		"--query-gpu=utilization.gpu,temperature.gpu,memory.used,power.draw",
		"--format=csv,noheader,nounits",
	).Output()
	if err != nil {
		return fmt.Errorf("nvidia-smi: %w", err)
	}
	fields := strings.Split(strings.TrimSpace(string(out)), ",")
	if len(fields) < 4 {
		return fmt.Errorf("nvidia-smi: unexpected output: %q", string(out))
	}
	parse := func(s string) float64 {
		v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
		return v
	}
	m.GPUUtilization = parse(fields[0])
	m.GPUTemperature = parse(fields[1])
	m.GPUMemoryUsedMB = int64(parse(fields[2]))
	m.PowerDrawWatts = parse(fields[3])
	return nil
}

// collectIPMI reads fan speed via ipmitool (best-effort).
func (c *Collector) collectIPMI(m *Metrics) error {
	out, err := exec.Command("ipmitool", "sdr", "type", "Fan").Output()
	if err != nil {
		return fmt.Errorf("ipmitool: %w", err)
	}
	// Parse first fan RPM reading
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "RPM") {
			parts := strings.Split(line, "|")
			if len(parts) >= 2 {
				rpm, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					m.FanSpeedRPM = rpm
					break
				}
			}
		}
	}
	return nil
}
