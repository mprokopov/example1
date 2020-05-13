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

type Simple struct {
	Name string
	Description string
	Url string
}

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "api_calls_total_counter",
	Help: "The total number of processed events",
})

func SimpleFactory (host string) Simple {
	return Simple{"Hello", "World", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
	simple := SimpleFactory(r.Host)

	jsonOutput, _ := json.Marshal(simple)

	counter.Inc() // add to Handler function

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started normally")
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Monitoring is started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
