# HTTP Status Codes — Шпаргалка

## 2xx — Успех

| Код | Когда использовать | Пример |
|-----|-------------------|--------|
| **200 OK** | GET, PUT, PATCH успешны | `GET /books/1` |
| **201 Created** | POST создал ресурс | `POST /books` → `Location: /books/42` |
| **204 No Content** | DELETE успешен, тело не нужно | `DELETE /books/42` |

## 4xx — Ошибка клиента

| Код | Когда использовать | Пример |
|-----|-------------------|--------|
| **400 Bad Request** | Невалидный JSON, синтаксис | `{ "title": }` |
| **401 Unauthorized** | Нет токена / токен невалиден | Запрос без `Authorization` |
| **403 Forbidden** | Есть auth, но нет прав | User пытается удалить чужой пост |
| **404 Not Found** | Ресурс не существует | `GET /books/99999` |
| **405 Method Not Allowed** | Метод не поддерживается | `PATCH /health` |
| **409 Conflict** | Конфликт состояния | Дубликат email при регистрации |
| **422 Unprocessable Entity** | Семантически неверные данные | email без `@`, отрицательная цена |
| **429 Too Many Requests** | Rate limit | >100 req/min |

## 5xx — Ошибка сервера

| Код | Когда использовать |
|-----|-------------------|
| **500 Internal Server Error** | Неожиданная ошибка, баг |
| **502 Bad Gateway** | Upstream недоступен |
| **503 Service Unavailable** | Сервер на maintenance |

## Частые ошибки новичков

```
❌ 200 + {"error": "not found"}     → используй 404
❌ 404 для "email уже занят"        → используй 409
❌ 500 для "title is required"      → используй 422
❌ 401 когда нет прав на ресурс     → используй 403
❌ 200 для POST create              → используй 201
```

## Go — helper функции

```go
func respondJSON(w http.ResponseWriter, status int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, code, message string) {
    respondJSON(w, status, map[string]any{
        "error": map[string]string{
            "code":    code,
            "message": message,
        },
    })
}
```
