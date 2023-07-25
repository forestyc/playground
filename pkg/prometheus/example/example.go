package main

import (
	"github.com/forestyc/playground/pkg/prometheus"
	"time"
)

func main() {
	p := prometheus.Prometheus{
		Addr: ":12112",
		Path: "/metrics",
	}
	p.Start()
	// counter
	counter := prometheus.NewCounter("test_counter", "test counter", "label")
	counter.Inc("inc")
	counter.Add(3.0, "add")
	counter1 := prometheus.NewCounter("test_counter", "test counter", "label")
	counter1.Inc("inc")
	counter1.Add(3.0, "add")
	// gauge
	gauge := prometheus.NewGauge("test_gauge", "test gauge", "label")
	gauge.Inc("inc")
	gauge.Dec("dec")
	// histogram
	histogram := prometheus.NewHistogram("test_histogram", "test histogram", "label")
	histogram.Observe(3.14, "observe")
	// summary
	summary := prometheus.NewSummary("test_summary", "test summary", "label")
	summary.Observe(3.14, "observe")
	time.Sleep(time.Hour)
}
