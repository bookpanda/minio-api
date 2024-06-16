package metrics

import (
	"fmt"

	"github.com/bookpanda/minio-api/constants"
	"github.com/prometheus/client_golang/prometheus"
)

type RequestsMetrics interface {
	AddRequest(domain constants.Domain, method constants.Method, statusCode int)
	GetCounterVec() *prometheus.CounterVec
}

type requestsMetricsImpl struct {
	requestsCounter *prometheus.CounterVec
}

func NewRequestsMetrics(requestsCounter *prometheus.CounterVec) RequestsMetrics {
	return &requestsMetricsImpl{
		requestsCounter: requestsCounter,
	}
}

func (m *requestsMetricsImpl) AddRequest(domain constants.Domain, method constants.Method, statusCode int) {
	m.requestsCounter.WithLabelValues(domain.String(), method.String(), fmt.Sprint(statusCode)).Inc()
}

func (m *requestsMetricsImpl) GetCounterVec() *prometheus.CounterVec {
	return m.requestsCounter
}
