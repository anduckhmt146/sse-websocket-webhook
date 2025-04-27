package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketClients   = make(map[*websocket.Conn]bool) // Active WebSocket clients
	websocketClientsMu sync.Mutex                       // Mutex for safe access
	upgrader           = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections (you can customize this)
			return true
		},
	}
)

// Handle WebSocket connections
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer ws.Close()

	websocketClientsMu.Lock()
	websocketClients[ws] = true
	websocketClientsMu.Unlock()

	fmt.Println("New WebSocket client connected")

	for {
		// Listen for messages from client (not necessary for push-only)
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}

	// Remove client when done
	websocketClientsMu.Lock()
	delete(websocketClients, ws)
	websocketClientsMu.Unlock()
	fmt.Println("WebSocket client disconnected")
}

// Broadcast message to all clients
func SendWebSocketMessage(w http.ResponseWriter, r *http.Request) {
	message := "🔔 New message at server!"

	websocketClientsMu.Lock()
	for client := range websocketClients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error writing message:", err)
			client.Close()
			delete(websocketClients, client)
		}
	}
	websocketClientsMu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message broadcasted"))
}
