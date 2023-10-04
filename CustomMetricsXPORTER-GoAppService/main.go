package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
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
		Name: "MyApps_Custom_metrics_Service",
		Help: "custom metrics",
	})
)

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	reg.MustRegister(customMetric)
	promHandaler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	go export()
	http.Handle("/metrics", promHandaler)
	http.ListenAndServe(":3000", nil)
}
