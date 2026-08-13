# Теория: chi router

## Установка

```go
import "github.com/go-chi/chi/v5"

r := chi.NewRouter()
r.Get("/books/{id}", getBook)
r.Post("/books", createBook)
http.ListenAndServe(":8080", r)
```

## URL Parameters

```go
id := chi.URLParam(r, "id")
```

## Route Groups

```go
r.Route("/api/v1", func(r chi.Router) {
    r.Get("/books", listBooks)
    r.Route("/books/{bookID}", func(r chi.Router) {
        r.Get("/", getBook)
        r.Put("/", updateBook)
        r.Delete("/", deleteBook)
    })
})
```

## Middleware

```go
r.Use(middleware.Logger)
r.Use(middleware.Recoverer)
r.Use(middleware.RequestID)
r.Use(middleware.Timeout(60 * time.Second))
```

## Sub-router mounting

```go
r.Mount("/admin", adminRouter())
```

## chi vs net/http ServeMux

| Feature | ServeMux 1.22 | chi |
|---------|---------------|-----|
| Path params | ✅ | ✅ |
| Middleware chain | ❌ | ✅ |
| Route groups | ❌ | ✅ |
| Sub-routers | ❌ | ✅ |
