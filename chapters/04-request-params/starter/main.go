package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		json.NewEncoder(w).Encode(map[string]any{
			"query": q.Get("q"),
			"page":  q.Get("page"),
			"limit": q.Get("limit"),
			"tags":  q["tag"],
		})
	})

	r.Get("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, err := strconv.Atoi(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "id must be integer"})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"id": id})
	})

	log.Println("Chapter 04 starter on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
