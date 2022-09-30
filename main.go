package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Hi!")

	handlerCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "metrics_example_handlers_requests_total",
		Help: "RPC latency distributions.",
	},
		[]string{"method", "handler", "code"},
	)
	prometheus.MustRegister(handlerCounter)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		handlerCounter.With(prometheus.Labels{
			"method":  r.Method,
			"handler": r.RequestURI,
			"code":    "200",
		}).Inc()

		fmt.Fprintf(w, "Hello every one!")
	})

	http.HandleFunc("/bue", func(w http.ResponseWriter, r *http.Request) {
		handlerCounter.With(prometheus.Labels{
			"method":  r.Method,
			"handler": r.RequestURI,
			"code":    "200",
		}).Inc()

		fmt.Fprintf(w, "Bue every one! See you soon!")
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":1234", nil)
}
