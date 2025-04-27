package handler

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	// Each client is a channel that receives messages
	ch chan string
}

var (
	sseClients   = make(map[*Client]bool) // Active clients
	sseClientsMu sync.Mutex               // Mutex to protect map
)

// Handle client connections for SSE
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// Set required headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow cross-origin (for testing)

	// Create new client
	client := &Client{
		ch: make(chan string),
	}

	// Register client
	sseClientsMu.Lock()
	sseClients[client] = true
	sseClientsMu.Unlock()

	// Remove client when done
	defer func() {
		sseClientsMu.Lock()
		delete(sseClients, client)
		sseClientsMu.Unlock()
	}()

	// Flusher
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Listen to client channel and send events
	for {
		select {
		case msg := <-client.ch:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-r.Context().Done():
			// Client closed connection
			return
		}
	}
}

// Endpoint to simulate sending a new notification
func SendHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("New Order received at %s", time.Now().Format(time.RFC3339))

	sseClientsMu.Lock()
	for client := range sseClients {
		client.ch <- message
	}
	sseClientsMu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message broadcasted"))
}
