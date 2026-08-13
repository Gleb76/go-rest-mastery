# Теория: Status Codes

См. также [docs/http-status-codes.md](../../docs/http-status-codes.md)

## Error response type

```go
type APIError struct {
    Error struct {
        Code    string            `json:"code"`
        Message string            `json:"message"`
        Details []FieldError      `json:"details,omitempty"`
    } `json:"error"`
}

type FieldError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}
```

## Map domain errors → HTTP

```go
func handleServiceError(w http.ResponseWriter, err error) {
    switch {
    case errors.Is(err, ErrNotFound):
        respondError(w, 404, "NOT_FOUND", err.Error())
    case errors.Is(err, ErrConflict):
        respondError(w, 409, "CONFLICT", err.Error())
    default:
        respondError(w, 500, "INTERNAL", "internal server error")
    }
}
```
