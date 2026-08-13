# Задачи — Глава 11

## ⭐ 1. Структура проекта
Создай `cmd/server`, `internal/handler`, `internal/service`, `internal/repository`, `internal/domain`

## ⭐ 2. Domain model
```go
type Book struct { ID int; Title, Author string; CreatedAt time.Time }
```

## ⭐ 3. Repository interface
```go
type BookRepository interface {
    Create(ctx context.Context, book *Book) error
    GetByID(ctx context.Context, id int) (*Book, error)
    List(ctx context.Context) ([]Book, error)
    Update(ctx context.Context, book *Book) error
    Delete(ctx context.Context, id int) error
}
```

## ⭐ 4. In-memory repository impl

## ⭐ 5. Service layer — business rules (title not empty)

## ⭐⭐ 6. Handler calls service, not repository

## ⭐⭐ 7. Domain errors: ErrNotFound, ErrValidation

## ⭐⭐ 8. Handler maps domain errors → HTTP status

## ⭐⭐⭐ 9. DTOs: CreateBookRequest, BookResponse in handler package

## ⭐⭐⭐ 10. Full CRUD through all layers

## Проверка
Handler не импортирует repository напрямую — только service.
