package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zrajesh/olx-backend/internal/config"
	"github.com/zrajesh/olx-backend/internal/db"
	"github.com/zrajesh/olx-backend/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	_, dbErr := db.Connect(cfg.DatabaseURL)
	if dbErr != nil {
		log.Fatalf("Database connection failed: %v", dbErr)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	srv := http.Server{
		Addr:   ":" + cfg.Port,
		Handler: mux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout: time.Second * 60,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}