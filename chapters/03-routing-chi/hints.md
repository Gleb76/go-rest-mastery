# Подсказки — Глава 03

## Custom 404
```go
r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    respondJSON(w, 404, map[string]string{"error": "not found"})
})
```

## Middleware signature
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

## Wrap response writer for status
Используй `middleware.NewWrapResponseWriter(w)` из chi для capture status code.
