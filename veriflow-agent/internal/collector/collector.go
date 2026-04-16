package collector

import "fmt"

// Collector orchestrates GPU and system metric collection.
type Collector struct{}

// New constructs a Collector.
func New() *Collector {
	return &Collector{}
}

// Collect reads metrics from nvidia-smi, ipmitool, and /proc/uptime.
// ipmitool and /proc/uptime failures are best-effort and do not abort collection.
func (c *Collector) Collect() (*Metrics, error) {
	m := &Metrics{}
	if err := collectGPU(m); err != nil {
		return nil, fmt.Errorf("collector.Collect: gpu: %w", err)
	}
	collectIPMI(m)
	collectUptime(m)
	return m, nil
}
