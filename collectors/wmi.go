package collectors

import "github.com/prometheus/client_golang/prometheus"

const (
	wmiNamespace = "wmi"
)

var Factories = make(map[string]func() (Collector, error))

// Interface a collector has to implement.
type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update(ch chan<- prometheus.Metric) (err error)
}
