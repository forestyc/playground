package main

import (
	"github.com/forestyc/playground/pkg/core/http"
	prometheus2 "github.com/forestyc/playground/pkg/core/prometheus"
	"time"
)

func main() {
	httpServer := http.NewServer(":12112", http.WithPrometheus("/metrics"))
	httpServer.Serve()
	// counter
	counter := prometheus2.NewCounter("test_counter", "test counter", "label")
	counter.Inc("inc")
	counter.Add(3.0, "add")

	counter2 := prometheus2.NewCounter("test_counter2", "test counter", "label")
	counter2.Inc("inc")
	counter2.Add(3.0, "add")
	// gauge
	gauge := prometheus2.NewGauge("test_gauge", "test gauge", "label")
	gauge.Inc("inc")
	gauge.Dec("dec")
	// histogram
	histogram := prometheus2.NewHistogram("test_histogram", "test histogram", "label")
	histogram.Observe(3.14, "observe")
	// summary
	summary := prometheus2.NewSummary("test_summary", "test summary", "label")
	summary.Observe(3.14, "observe")
	time.Sleep(1 * time.Minute)
	counter.Unregister()
	counter2.Unregister()
	gauge.Unregister()
	histogram.Unregister()
	summary.Unregister()
	httpServer.Close()
}
