package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zrajesh/olx-backend/internal/config"
)

func main() {
	cfg := config.MustLoad()
  
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

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