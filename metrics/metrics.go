package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metrics interface {
	Registry() *prometheus.Registry
}

type metricsImpl struct {
	registry *prometheus.Registry
}

func NewMetrics(registry *prometheus.Registry, requestsMetrics RequestsMetrics) Metrics {
	// requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
	// 	Name:    "http_request_duration_seconds",
	// 	Help:    "A histogram of the HTTP request durations in seconds.",
	// 	Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
	// })

	// requestsMetrics := prometheus.NewCounterVec(prometheus.CounterOpts{
	// 	Name: "api_requests_total",
	// 	Help: "Total number of API requests by domain, method and status code",
	// }, []string{"domain", "method", "status_code"})

	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		requestsMetrics.GetCounterVec(),
		// requestDurations,
	)

	return &metricsImpl{
		registry: registry,
	}
}

func (m *metricsImpl) Registry() *prometheus.Registry {
	return m.registry
}
