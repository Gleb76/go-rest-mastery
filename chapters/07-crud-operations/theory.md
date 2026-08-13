# Теория: CRUD Operations

## CRUD mapping

| Operation | HTTP | Path | Status |
|-----------|------|------|--------|
| Create | POST | /books | 201 |
| Read all | GET | /books | 200 |
| Read one | GET | /books/{id} | 200/404 |
| Update | PUT | /books/{id} | 200/404 |
| Partial update | PATCH | /books/{id} | 200/404 |
| Delete | GET | /books/{id} | 204/404 |

## In-memory store

```go
type BookStore struct {
    mu    sync.RWMutex
    books map[int]Book
    nextID int
}
```

## POST — Location header

```go
w.Header().Set("Location", fmt.Sprintf("/api/v1/books/%d", book.ID))
w.WriteHeader(http.StatusCreated)
```

## PUT vs PATCH

- PUT — полная замена (missing fields → zero value)
- PATCH — частичное обновление (только переданные поля)
