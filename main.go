package main

import (
	"anduckhmt146/sse-webhook/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/sse", handler.SSEHandler)
	http.HandleFunc("/sse/send", handler.SendHandler)
	http.HandleFunc("/webhook", handler.WebhookHandler)
	http.HandleFunc("/ws", handler.WebSocketHandler)
	http.HandleFunc("/ws/send", handler.SendWebSocketMessage)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
