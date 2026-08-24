package main

import (
	"ch03/internal/handler"
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
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.NotFound(notFound)
	r.MethodNotAllowed(notAllowed)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "API v1 работает",
				"version": "1.0",
			})
		})

		r.Get("/ping", handler.HandlePong)
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", handler.GetUser)

			r.Route("/{id}/posts", func(r chi.Router) {
				r.Post("/", handler.CreatePost)
				r.Get("/", handler.GetUserPosts)
				r.Get("/{postID}", handler.GetUserPostId)
			})
		})
		r.Route("/books", func(r chi.Router) {
			r.Get("/", handler.ListBooks)
			r.Post("/", handler.CreateBook)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handler.GetBook)
				r.Put("/", handler.UpdateBook)
				r.Patch("/", handler.PatchBook)
				r.Delete("/", handler.DeleteBook)
			})
		})
	})

	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":   "Endpoint not found",
		"path":    r.URL.Path,
		"method":  r.Method,
		"message": "Доступные пути: /api/v1/ping, /api/v1/books, /api/v1/books/{id}, /api/v1/users/{id}",
	})
}

func customLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(ww, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s - %d (%v)",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			ww.statusCode,
			duration,
		)
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func notAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{
		"not allowed": "That's method not Allowed",
		"path":        r.URL.Path,
		"method":      r.Method,
		"message":     "Use different method",
	})
}

// # Задачи — Глава 03

// ## ⭐ 1. Basic routes
// chi router: GET /, GET /ping → pong

// ## ⭐ 2. Resource routes
// CRUD routes для `/books/{id}` (handlers могут быть заглушками)

// ## ⭐ 3. Route groups
// `/api/v1/*` — все API routes

// ## ⭐ 4. URL params
// GET /users/{id} → `{"id": "<id>"}`

// ## ⭐⭐ 5. Middleware logger
// Логируй: method, path, duration, status

// ## ⭐⭐ 6. NotFound
// Кастомный 404 JSON handler

// ## ⭐⭐ 7. Method Not Allowed
// 405 JSON для неподдерживаемых методов

// ## ⭐⭐ 8. Mount sub-router
// `/admin` sub-router с routes: GET /stats, POST /cache/clear

// ## ⭐⭐⭐ 9. Nested resources
// `/users/{userID}/posts/{postID}`

// ## ⭐⭐⭐ 10. File organization
// `internal/handler/book_handler.go`, `internal/router/router.go`

// ## ⭐⭐⭐ 11. chi middleware chain
// RequestID + Logger + Recoverer + Timeout(30s)

// ## ⭐⭐⭐ 12. Books API skeleton
// Полный routing для Books (List, Get, Create, Update, Delete) — handlers return mock JSON

// ```bash
// curl http://localhost:8080/api/v1/books
// curl http://localhost:8080/api/v1/books/1
// ```
