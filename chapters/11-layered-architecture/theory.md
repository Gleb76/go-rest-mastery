# Теория: Слоистая архитектура

## Слои

```
HTTP Request
    ↓
Handler (HTTP) — decode request, encode response, status codes
    ↓
Service (Business) — validation, business rules, orchestration
    ↓
Repository (Data) — CRUD, SQL, caching
    ↓
Database
```

## Правила

1. **Зависимости только вниз** — handler → service → repo
2. **Domain в центре** — типы и интерфейсы в `internal/domain`
3. **Handler не знает SQL**
4. **Repository не знает HTTP**

## Domain package

```go
// internal/domain/book.go
package domain

type Book struct { ... }

type BookRepository interface { ... }

var ErrNotFound = errors.New("book not found")
```

## Service

```go
type BookService struct {
    repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
    return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, title, author string) (*domain.Book, error) {
    if title == "" {
        return nil, domain.ErrValidation
    }
    // ...
}
```

## Handler

```go
type BookHandler struct {
    service *service.BookService
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateBookRequest
    json.NewDecoder(r.Body).Decode(&req)
    book, err := h.service.CreateBook(r.Context(), req.Title, req.Author)
    if err != nil { handleError(w, err); return }
    respondJSON(w, 201, toResponse(book))
}
```

## main.go — wiring

```go
repo := memory.NewBookRepository()
svc := service.NewBookService(repo)
handler := handler.NewBookHandler(svc)
router := router.New(handler)
```
