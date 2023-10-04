package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func export() {
	for {
		customMetric.Inc()
		time.Sleep(3 * time.Second)
	}
}

var (
	customMetric = promauto.NewCounter(prometheus.CounterOpts{
		Name: "MyApps_Custom_metrics_pod",
		Help: "custom metrics",
	})
)

func main() {
	go export()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
}
