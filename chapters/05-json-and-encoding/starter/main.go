package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateBookRequest struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description *string `json:"description,omitempty"`
}

var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert Martin", CreatedAt: time.Now()},
}

func main() {
	r := chi.NewRouter()

	r.Get("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"data": books, "count": len(books)})
	})

	r.Post("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		var req CreateBookRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
			return
		}
		defer r.Body.Close()

		book := Book{
			ID:        len(books) + 1,
			Title:     req.Title,
			Author:    req.Author,
			Description: req.Description,
			CreatedAt: time.Now().UTC(),
		}
		books = append(books, book)

		w.Header().Set("Location", "/api/v1/books/"+strconv.Itoa(book.ID))
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	})

	log.Println("Chapter 05 starter on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
