package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Chapter 03 — chi router"})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/books", listBooks)
		r.Get("/books/{id}", getBook)
		r.Post("/books", createBook)
	})

	log.Println("Chapter 03 starter on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]any{
		{"id": 1, "title": "The Go Programming Language"},
		{"id": 2, "title": "Clean Code"},
	})
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	json.NewEncoder(w).Encode(map[string]any{"id": id, "title": "Mock Book"})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}
