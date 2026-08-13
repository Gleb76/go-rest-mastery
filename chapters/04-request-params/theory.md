# Теория: Request Parameters

## Query Parameters

```go
q := r.URL.Query()
page := q.Get("page")       // string, "" if missing
pages := q["tag"]           // []string for repeated params
```

## Parsing with defaults

```go
func queryInt(r *http.Request, key string, defaultVal int) int {
    s := r.URL.Query().Get(key)
    if s == "" { return defaultVal }
    v, err := strconv.Atoi(s)
    if err != nil { return defaultVal }
    return v
}
```

## Path Parameters (chi)

```go
r.Get("/books/{id}", handler)
id := chi.URLParam(r, "id")

// Regex constraint
r.Get("/users/{id:[0-9]+}", handler)
```

## Headers

```go
auth := r.Header.Get("Authorization")
contentType := r.Header.Get("Content-Type")
r.Header.Get("X-Request-ID")  // case-insensitive
```

## Request Body

```go
if r.Header.Get("Content-Type") != "application/json" {
    http.Error(w, "unsupported media type", 415)
    return
}
```
