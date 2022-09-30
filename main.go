package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Hi!")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello every one!")
	})
	http.HandleFunc("/bue", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bue every one! See you soon!")
	})
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":1234", nil)
}
