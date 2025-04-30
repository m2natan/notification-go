package push

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	clients    = make(map[string]chan string) // Map to store channels for each client
	clientsMux sync.Mutex                     // Mutex to protect the clients map
)

// SSEHandler handles Server-Sent Events
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	recipient := r.URL.Query().Get("recipient")
	if recipient == "" {
		http.Error(w, "Recipient not specified", http.StatusBadRequest)
		return
	}

	// Set CORS and SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Ensure the connection supports flushing for SSE
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Create a channel to send messages to the specific recipient
	msgChannel := make(chan string)

	// Register the client
	clientsMux.Lock()
	clients[recipient] = msgChannel
	clientsMux.Unlock()

	// Ensure cleanup when client disconnects
	defer func() {
		clientsMux.Lock()
		delete(clients, recipient)
		clientsMux.Unlock()
		close(msgChannel)
	}()

	// Listen for messages for this client and send them
	for msg := range msgChannel {
		fmt.Fprintf(w, "data: %s\n\n", msg)
		flusher.Flush()
	}
}

// sendNotification sends a notification to a specific recipient
func SendNotification(recipient, message string) error {
	clientsMux.Lock()
	defer clientsMux.Unlock()

	if msgChannel, exists := clients[recipient]; exists {
		msgChannel <- message
		return nil
	} else {
		log.Printf("No client found for recipient: %s", recipient)
	}
	return fmt.Errorf("no client found for recipient: %s", recipient)
}
