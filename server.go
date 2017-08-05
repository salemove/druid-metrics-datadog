package main

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Server struct {
	metricPublisher *DatadogPublisher
}

func (s *Server) healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func (s *Server) metricsHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// Jason does not know how to parse a root level array. Wrapping it in an object.
	wrappedBody := fmt.Sprintf("{\"metrics\": %s}", body)

	payload, err := jason.NewObjectFromBytes([]byte(wrappedBody))
	if err != nil {
		panic(err)
	}

	metrics, err := payload.GetObjectArray("metrics")
	if err != nil {
		panic(err)
	}

	for _, metric := range metrics {
		s.metricPublisher.Publish(metric)
	}

	log.Println("Processed request")
	w.WriteHeader(202)
}

func main() {
	datadogAddr := os.Getenv("DATADOG_ADDRESS")
	if datadogAddr == "" {
		datadogAddr = "127.0.0.1:8125"
	}

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "8424"
	}

	metricPublisher, err := CreateDatadogPublisher(datadogAddr)
	if err != nil {
		panic(err)
	}
	server := Server{metricPublisher: metricPublisher}

	http.HandleFunc("/healthz", server.healthzHandler)
	http.HandleFunc("/metrics", server.metricsHandler)

	log.Println("Listening on " + listenPort)
	http.ListenAndServe(":"+listenPort, nil)
}
