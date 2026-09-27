package middleware

import (
	"strconv"
	"time"

	"cloud-assessment-tool/metrics"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("PROMETHEUS MIDDLEWARE HIT:", c.Request.URL.Path)
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		metrics.APIRequests.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			status,
		).Inc()

		metrics.APIResponseTime.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)
	}
}
