package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

func main() {

	breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name: "MyCircuitBreaker",
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
		MaxRequests: 2,               // max request to downstream if in half-open state
		Timeout:     5 * time.Second, // period time to retry
	})

	circuitBreakerRequest := func() error { // per 1 service 1 endpoint
		_, err := breaker.Execute(makeRequest)
		return err
	}

	// Set up an HTTP server to monitor circuit breaker status
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		state := breaker.State()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Circuit Breaker State: %v\n", state)
	})

	// Start the HTTP server in a separate goroutine
	go func() {
		fmt.Println("Starting HTTP server on :9000")
		if err := http.ListenAndServe(":9000", nil); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()

	// Execute the circuit breaker request
	i := 0
	for {
		i++
		err := circuitBreakerRequest()
		if err != nil {
			fmt.Printf("Request %v failed. err: %v \n", i, err)
		} else {
			fmt.Printf("Request %v succeeded \n", i)
		}
		time.Sleep(time.Second)
	}
}

// Request to downstream
func makeRequest() (interface{}, error) {
	downstreamURL := "http://localhost:20000"
	url := downstreamURL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil, nil
}
