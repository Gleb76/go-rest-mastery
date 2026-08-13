# Задачи — Глава 03

## ⭐ 1. Basic routes
chi router: GET /, GET /ping → pong

## ⭐ 2. Resource routes
CRUD routes для `/books/{id}` (handlers могут быть заглушками)

## ⭐ 3. Route groups
`/api/v1/*` — все API routes

## ⭐ 4. URL params
GET /users/{id} → `{"id": "<id>"}`

## ⭐⭐ 5. Middleware logger
Логируй: method, path, duration, status

## ⭐⭐ 6. NotFound
Кастомный 404 JSON handler

## ⭐⭐ 7. Method Not Allowed
405 JSON для неподдерживаемых методов

## ⭐⭐ 8. Mount sub-router
`/admin` sub-router с routes: GET /stats, POST /cache/clear

## ⭐⭐⭐ 9. Nested resources
`/users/{userID}/posts/{postID}`

## ⭐⭐⭐ 10. File organization
`internal/handler/book_handler.go`, `internal/router/router.go`

## ⭐⭐⭐ 11. chi middleware chain
RequestID + Logger + Recoverer + Timeout(30s)

## ⭐⭐⭐ 12. Books API skeleton
Полный routing для Books (List, Get, Create, Update, Delete) — handlers return mock JSON

```bash
curl http://localhost:8080/api/v1/books
curl http://localhost:8080/api/v1/books/1
```
