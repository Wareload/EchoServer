package main

import (
	"fmt"
	"net/http"
	"sort"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Print request details in the response
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintln(w, "Headers:")

	// Collect and sort headers
	var headers []string
	for name := range r.Header {
		headers = append(headers, name)
	}
	sort.Strings(headers)

	// Print sorted headers
	for _, name := range headers {
		for _, value := range r.Header[name] {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}

	// Set response headers
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Custom-Header", "EchoServer")

	// Print response headers in the response
	fmt.Fprintln(w, "Response Headers:")

	// Collect and sort response headers
	headers = headers[:0] // Reuse the slice
	for name := range w.Header() {
		headers = append(headers, name)
	}
	sort.Strings(headers)

	// Print sorted response headers
	for _, name := range headers {
		for _, value := range w.Header()[name] {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Echo server response\n"))
}

func main() {
	http.HandleFunc("/", echoHandler)
	fmt.Println("Starting server on :2000")
	http.ListenAndServe(":2000", nil)
}
