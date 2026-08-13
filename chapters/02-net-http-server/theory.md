# Теория: net/http сервер

## Handler interface

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Любая функция с сигнатурой `func(w, r)` может быть handler через `http.HandlerFunc`.

## ServeMux (Go 1.22+)

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /health", healthHandler)
mux.HandleFunc("GET /users/{id}", getUserHandler)  // path params в 1.22+
http.ListenAndServe(":8080", mux)
```

## http.Server — production basics

```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
}
server.ListenAndServe()
```

## Graceful Shutdown

```go
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
server.Shutdown(ctx)
```

## Организация handlers

```go
// Плохо — всё в main
// Хорошо — отдельные функции/файлы
func healthHandler(w http.ResponseWriter, r *http.Request) {
    respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
```

## respondJSON helper

Создай один раз — используй во всех главах:

```go
func respondJSON(w http.ResponseWriter, status int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}
```
