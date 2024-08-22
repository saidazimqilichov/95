package main

import (
	"log"
	"net/http"
	"95/internal/config"
	"95/internal/handler"
	"95/internal/storage"
)

func main() {
	cfg := config.New()

	db, err := storage.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}
	defer db.Close()

	webhookHandler := handler.NewWebhookHandler(db)

	http.HandleFunc("/webhook", webhookHandler.Handle)

	log.Printf("Server started on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
