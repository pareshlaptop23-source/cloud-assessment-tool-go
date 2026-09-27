package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	APIRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cloudtool_api_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method", "route", "status"},
	)

	APIResponseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "cloudtool_api_response_time_seconds",
			Help: "API response time in seconds",
		},
		[]string{"method", "route"},
	)
)

func Init() {
	prometheus.MustRegister(APIRequests)
	prometheus.MustRegister(APIResponseTime)
}
