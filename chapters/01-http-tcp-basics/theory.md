# Теория: HTTP и TCP

## Модель OSI — упрощённо

```
Приложение (Go HTTP server)
        ↓
    HTTP (протокол)
        ↓
    TCP (надёжная доставка)
        ↓
    IP (маршрутизация)
        ↓
    Ethernet / Wi-Fi
```

**TCP** — гарантирует доставку байтов в правильном порядке.  
**HTTP** — формат сообщений поверх TCP (request/response).

## HTTP Request

```http
GET /api/v1/books HTTP/1.1
Host: localhost:8080
Accept: application/json
User-Agent: curl/8.0

```

Структура:
1. **Request line:** `METHOD PATH HTTP/VERSION`
2. **Headers:** key: value (метаданные)
3. **Blank line**
4. **Body** (опционально, для POST/PUT/PATCH)

## HTTP Response

```http
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 42

{"message":"Hello, REST!"}
```

1. **Status line:** `HTTP/VERSION CODE REASON`
2. **Headers**
3. **Blank line**
4. **Body**

## HTTP Methods (основные)

| Method | Назначение | Body | Safe | Idempotent |
|--------|-----------|------|------|------------|
| GET | Получить ресурс | ❌ | ✅ | ✅ |
| POST | Создать / действие | ✅ | ❌ | ❌ |
| PUT | Заменить ресурс | ✅ | ❌ | ✅ |
| PATCH | Частичное обновление | ✅ | ❌ | ❌ |
| DELETE | Удалить | ❌ | ❌ | ✅ |

## curl — твой лучший друг

```bash
# Verbose — показывает request и response headers
curl -v http://localhost:8080/

# Только headers ответа
curl -I http://localhost:8080/

# POST с телом
curl -X POST http://localhost:8080/echo \
  -H "Content-Type: text/plain" \
  -d "Hello HTTP"

# Сохранить response headers
curl -D - http://localhost:8080/
```

## Go: net/http — минимальный сервер

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Method: %s, Path: %s\n", r.Method, r.URL.Path)
})
http.ListenAndServe(":8080", nil)
```

`http.HandleFunc` регистрирует handler для пути.  
`http.ListenAndServe` слушает TCP порт и принимает соединения.

## Ключевые типы

- `http.ResponseWriter` — пишешь ответ (status, headers, body)
- `*http.Request` — входящий запрос (method, URL, headers, body)

## Порядок записи ответа

```go
w.Header().Set("Content-Type", "application/json")  // 1. headers
w.WriteHeader(http.StatusOK)                          // 2. status (если не 200)
w.Write([]byte(`{"ok":true}`))                        // 3. body
```

⚠️ `WriteHeader` можно вызвать только один раз. После `Write()` status = 200 автоматически.

## Дальше

→ [exercises.md](./exercises.md)
