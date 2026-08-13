# Теория: JSON в Go

## Struct tags

```go
type Book struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Author    string    `json:"author,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    internal  string    // не экспортируется — не попадёт в JSON
}
```

## Decode request

```go
var req CreateBookRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    respondError(w, 400, "INVALID_JSON", "invalid request body")
    return
}
defer r.Body.Close()
```

## Encode response

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(book)
```

## DTO pattern

```go
type CreateBookRequest struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}

type BookResponse struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}
```

## Null vs omitempty

- `omitempty` — поле пропускается если zero value
- `*string` / `*time.Time` — для nullable JSON null
