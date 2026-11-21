package main

import (
	"log/slog"
	"net/http"

	"api-gateway/internal/server"
)

func main() {
	mux := http.NewServeMux()

	log := slog.Default()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("GET /check", server.CheckHealth)

	log.Info("Starting server on :8080")
	if err := s.ListenAndServe(); err != nil {
		log.Error("Failed to start server", slog.Any("error", err))
	}
}
