package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "server_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
)

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
	})
}
