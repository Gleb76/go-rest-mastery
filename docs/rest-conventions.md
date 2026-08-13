# REST Conventions — Соглашения

## URL Design

### Ресурсы — существительные во множественном числе

```
✅ GET  /api/v1/books
✅ GET  /api/v1/books/42
✅ POST /api/v1/books

❌ GET  /api/v1/getBooks
❌ POST /api/v1/createBook
❌ GET  /api/v1/book/42        (единственное число — допустимо, но множественное предпочтительнее)
```

### Вложенные ресурсы

```
GET  /api/v1/users/1/tasks          # задачи пользователя
POST /api/v1/projects/5/tasks       # создать задачу в проекте
GET  /api/v1/tasks/10/comments      # комментарии к задаче
```

Максимум **2 уровня** вложенности. Глубже — используй query params:

```
GET /api/v1/comments?task_id=10     # вместо /tasks/10/projects/5/comments
```

### HTTP-методы → действия

| Метод | Действие | Идемпотентность | Safe |
|-------|----------|-----------------|------|
| GET | Получить | ✅ | ✅ |
| POST | Создать | ❌ | ❌ |
| PUT | Полная замена | ✅ | ❌ |
| PATCH | Частичное обновление | ❌* | ❌ |
| DELETE | Удалить | ✅ | ❌ |

*PATCH формально не идемпотентен, но на практике часто делают идемпотентным.

## Формат ответов

### Успех — один ресурс

```json
{
  "id": 42,
  "title": "Clean Code",
  "author": "Robert Martin",
  "created_at": "2026-01-15T10:30:00Z"
}
```

### Успех — коллекция с pagination

```json
{
  "data": [...],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 156,
    "total_pages": 8
  }
}
```

### Ошибка — единый формат

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input",
    "details": [
      {"field": "title", "message": "required"}
    ]
  }
}
```

## Именование полей

- **JSON:** `snake_case` (`created_at`, `user_id`)
- **Go struct tags:** `json:"created_at"`
- **Query params:** `snake_case` (`?page=1&per_page=20&sort_by=created_at`)

## Версионирование

Предпочитай **URL prefix**:

```
/api/v1/books
/api/v2/books
```

Header `Accept: application/vnd.myapi.v1+json` — для enterprise, но сложнее.

## Idempotency-Key (POST)

Для критичных операций (платежи, создание заказов):

```
POST /api/v1/orders
Idempotency-Key: 550e8400-e29b-41d4-a716-446655440000
```

Повторный запрос с тем же ключом → тот же результат, без дубликата.
