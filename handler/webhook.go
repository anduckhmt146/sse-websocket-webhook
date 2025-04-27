package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WebhookPayload struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received webhook: %+v\n", payload)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received"))
}
