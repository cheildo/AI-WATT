package collector

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const uptimeBaselineSecs = 7 * 24 * 3600 // 7-day window for uptime percentage

// collectIPMI reads fan speed via ipmitool (best-effort — not all hardware has IPMI).
func collectIPMI(m *Metrics) {
	out, err := exec.Command("ipmitool", "sdr", "type", "Fan").Output()
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "RPM") {
			parts := strings.Split(line, "|")
			if len(parts) >= 2 {
				if rpm, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					m.FanSpeedRPM = rpm
					return
				}
			}
		}
	}
}

// collectUptime reads /proc/uptime and expresses uptime as a percentage over
// a 7-day baseline. A system running continuously for 7+ days reports 100%.
func collectUptime(m *Metrics) {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		m.UptimePct = 100 // assume up if /proc/uptime unavailable (non-Linux)
		return
	}
	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		m.UptimePct = 100
		return
	}
	uptimeSecs, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		m.UptimePct = 100
		return
	}
	pct := uptimeSecs / uptimeBaselineSecs * 100.0
	if pct > 100.0 {
		pct = 100.0
	}
	m.UptimePct = pct
}
