# Bad API Example — для аудита (Глава 06)

Найди все нарушения REST и исправь в своём redesign.

## Текущий API (ПЛОХОЙ)

```
POST /api/getAllUsers
POST /api/createUser
POST /api/deleteUserById?id=5
GET  /api/user/update/5?name=John
POST /api/doLogin
GET  /api/session/check
PUT  /api/users/5/delete
GET  /api/getUserTasks/5
POST /api/assignTaskToUser
```

## Проблемы для поиска

1. RPC-style URLs (глаголы в path)
2. POST для read operations
3. GET для update (side effects)
4. Query param для delete
5. Inconsistent naming (camelCase, mixed)
6. No versioning
7. Session-based state (not stateless)
8. No proper HTTP methods
9. No resource nesting conventions
10. No standard error format

## Твоё задание

Перепиши в RESTful API с `/api/v1/` prefix. Документируй в `workspace/ch06/REDESIGN.md`.
