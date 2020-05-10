package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "api_calls_total",
	Help: "The total number of processed events",
})

type Simple struct {
	Name string
	Description string
	Url string
}

func SimpleFactory (host string) Simple {
	return Simple{"Hello", "World", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
	simple := SimpleFactory(r.Host)

	jsonOutput, _ := json.Marshal(simple)

	counter.Inc() // inc counter

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started normally")
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Monitoring endpoint added")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
