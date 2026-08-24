package main

import (
	"ch02/internal/handler"
	"ch02/internal/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Регистрация всех обработчиков
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/ready", handler.Ready)
	mux.HandleFunc("/admin/ready", handler.AdminReady)
	mux.HandleFunc("/info", handler.Info)
	mux.HandleFunc("/resource", handler.Resource)
	mux.HandleFunc("/slow", handler.Slow)

	// Применяем middleware
	handlerWithMiddleware := middleware.RequestIDMiddleware(mux)

	// Настраиваем http.Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handlerWithMiddleware,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server starting on :8080")
	log.Fatal(server.ListenAndServe())
}
