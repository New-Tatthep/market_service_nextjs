package product

import (
	"fmt"
	"net/http"
	"time"
)

// StreamUpdates is the HTTP handler for streaming updates via SSE
func StreamUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulating a stream of updates
	for {
		fmt.Fprintf(w, "data: %s\n\n", time.Now().String())

		// Flush the buffer
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		time.Sleep(1 * time.Second) // Adjust the interval as needed
	}
}
