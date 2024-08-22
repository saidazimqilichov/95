package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"95/internal/storage"
)

type WebhookHandler struct {
	db storage.Database
}

func NewWebhookHandler(db storage.Database) *WebhookHandler {
	return &WebhookHandler{db: db}
}

func (h *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var event map[string]interface{}
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Error processing JSON", http.StatusBadRequest)
		return
	}

	if event["repository"] == nil {
		http.Error(w, "This is not a repository event", http.StatusBadRequest)
		return
	}

	log.Printf("Received event: %v", event)

	if err := h.db.SaveEvent(event); err != nil {
		log.Printf("Error saving event: %v", err)
		http.Error(w, "Error saving event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
