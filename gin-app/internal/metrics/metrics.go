package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RqCounter, RqSuccCounter     prometheus.Counter
	UserCounter, UserSuccCounter prometheus.Counter
)

func InitRestMetrics() {
	RqCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rq_send",
		Help: "The total number of request to send",
	})
	RqSuccCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rq_success_send",
		Help: "The total success number of request to send",
	})
}
func InitWorkerMetrics() {
	UserCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "new_user",
		Help: "The total number user",
	})
	UserSuccCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "create_new_user",
		Help: "The total success create user",
	})
}
