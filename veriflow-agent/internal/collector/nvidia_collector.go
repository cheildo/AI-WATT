package collector

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// collectGPU runs nvidia-smi and parses GPU utilisation, temperature, memory,
// power draw, and uncorrected ECC error count.
func collectGPU(m *Metrics) error {
	out, err := exec.Command(
		"nvidia-smi",
		"--query-gpu=utilization.gpu,temperature.gpu,memory.used,power.draw,ecc.errors.uncorrected.total.volatile",
		"--format=csv,noheader,nounits",
	).Output()
	if err != nil {
		return fmt.Errorf("nvidia-smi: %w", err)
	}

	fields := strings.Split(strings.TrimSpace(string(out)), ",")
	if len(fields) < 5 {
		return fmt.Errorf("nvidia-smi: unexpected output: %q", string(out))
	}

	parseF := func(s string) float64 {
		v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
		return v
	}
	parseI := func(s string) int64 {
		v, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		return v
	}

	m.GPUUtilization  = parseF(fields[0])
	m.GPUTemperature  = parseF(fields[1])
	m.GPUMemoryUsedMB = int64(parseF(fields[2]))
	m.PowerDrawWatts  = parseF(fields[3])
	m.ECCErrors       = parseI(fields[4])
	// Normalise ECC count to a rate fraction for the backend scorer.
	// 1 error → 0.0001 (0.01%), 10 errors → 0.001 (0.1%).
	m.GPUErrorRate = float64(m.ECCErrors) / 10000.0

	return nil
}
