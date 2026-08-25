package main

import (
	"ch04/internal/query"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func parseSort(r *http.Request) (sortBy, order string) {
	q := r.URL.Query()
	allowed := map[string]bool{
		"title": true, "created_at": true, "updated_at": true,
		"price": true, "rating": true, "id": true,
	}
	sortBy = "created_at"
	if v := q.Get("sort_by"); v != "" && allowed[v] {
		sortBy = v
	}
	order = "desc"
	if v := q.Get("order"); v != "" && (v == "asc" || v == "desc") {
		order = v
	}
	return
}

func parseFilters(r *http.Request) map[string]string {
	filters := make(map[string]string)
	exclude := map[string]bool{
		"page": true, "per_page": true, "sort_by": true,
		"order": true, "q": true, "tag": true,
	}
	for key, values := range r.URL.Query() {
		if !exclude[key] && len(values) > 0 {
			filters[key] = values[0]
		}
	}
	return filters
}

func main() {
	r := chi.NewRouter()

	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		pagination := query.ParsePagination(r)
		sortBy, order := parseSort(r)
		filters := parseFilters(r)

		json.NewEncoder(w).Encode(map[string]any{
			"query":    q.Get("q"),
			"page":     pagination.Page,
			"per_page": pagination.PerPage,
			"offset":   pagination.Offset,
			"tags":     q["tag"],
			"sort_by":  sortBy,
			"order":    order,
			"filters":  filters,
		})
	})

	r.Get("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, err := strconv.Atoi(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "ID must be an integer",
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"id": id,
		})
	})

	r.Get("/headers/Authorization", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Authorization header required",
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"Authorization": auth,
		})
	})

	r.Post("/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Content-Type must be application/json",
			})
			return
		}
		var data map[string]any
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid JSON body",
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Data received",
			"data":    data,
		})
	})

	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		searchQuery := q.Get("q")
		author := q.Get("author")
		year := q.Get("year")

		if year != "" {
			if _, err := strconv.Atoi(year); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "year must be an integer",
				})
				return
			}
		}

		json.NewEncoder(w).Encode(map[string]any{
			"query":   searchQuery,
			"author":  author,
			"year":    year,
			"message": "Search results will appear here",
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
