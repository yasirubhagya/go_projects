package main

import (
	"fmt"
	"net/http"
	"time"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events (you can replace this with real data)
	for i := 0; ; i++ {
		select {
		case <-r.Context().Done():
			fmt.Println("connection closed from client")
			return
		default:
			fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d %v", i, time.Now()))
			time.Sleep(2 * time.Second)
			w.(http.Flusher).Flush()
		}
	}
}

func main() {
	http.HandleFunc("/events", eventsHandler)
	http.ListenAndServe(":8080", nil)
}
