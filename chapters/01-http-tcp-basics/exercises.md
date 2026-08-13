# Задачи — Глава 01

Выполняй в `workspace/ch01/`. Создай `go mod init` и пиши с нуля.

---

## ⭐ 1. Hello HTTP

Создай сервер на `:8080`:
- `GET /` → текст `Go REST Mastery — Chapter 01`
- `GET /hello` → `Hello, HTTP!`

**Проверка:**
```bash
curl http://localhost:8080/
curl http://localhost:8080/hello
```

---

## ⭐ 2. Echo endpoint

`POST /echo` — возвращает тело запроса как есть.

```bash
curl -X POST http://localhost:8080/echo -d "test body"
# → test body
```

Подсказка: `io.ReadAll(r.Body)`

---

## ⭐ 3. Method checker

`ANY /method` — возвращает JSON:
```json
{"method": "GET", "path": "/method"}
```

```bash
curl http://localhost:8080/method
curl -X DELETE http://localhost:8080/method
```

---

## ⭐⭐ 4. Headers inspector

`GET /headers` — возвращает JSON со всеми request headers.

```bash
curl -H "X-Custom: test123" http://localhost:8080/headers
```

---

## ⭐⭐ 5. Status playground

| Endpoint | Status |
|----------|--------|
| GET /status/200 | 200 |
| GET /status/404 | 404 |
| GET /status/500 | 500 |

Используй `w.WriteHeader()`.

---

## ⭐⭐ 6. Content-Type

`GET /json` — верни JSON `{"status":"ok"}` с правильным `Content-Type: application/json`.

`GET /html` — верни `<h1>Hello</h1>` с `Content-Type: text/html`.

---

## ⭐⭐⭐ 7. Query echo

`GET /search?q=golang&page=1` → JSON:
```json
{"query": "golang", "page": "1"}
```

Используй `r.URL.Query()`.

---

## ⭐⭐⭐ 8. 404 handler

Любой неизвестный путь → `404` + JSON:
```json
{"error": "not found", "path": "/unknown"}
```

Подсказка: handler для `/` не ловит всё — используй отдельную регистрацию или DefaultServeMux.

---

## Бонус ⭐⭐⭐⭐

TCP-сервер без HTTP (чистый `net`):
- Слушай `:9000`
- При подключении отправь `220 Welcome\r\n`
- Прочитай одну строку от клиента и отправь `250 OK: <строка>\r\n`

```bash
nc localhost 9000
```
