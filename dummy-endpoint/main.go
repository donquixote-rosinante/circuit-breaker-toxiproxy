package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the desired status code
	statusCode := http.StatusOK

	// Set the response header
	w.Header().Set("Content-Type", "text/plain")

	// Set the status code for the response
	w.WriteHeader(statusCode)

	// Write the response content (string)
	responseString := "Hello, this is a custom status code example!"
	fmt.Fprintln(w, responseString)
}

func main() {
	http.HandleFunc("/", handler)

	// Define the server address and port
	serverAddr := "localhost:8080"

	fmt.Printf("Server is listening at %s\n", serverAddr)

	// Start the HTTP server
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
