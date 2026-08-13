#!/usr/bin/env python3
"""Generate chapter documentation for go-rest-mastery."""

import os

BASE = "/Users/glebklyga/Projects/go-rest-mastery/chapters"

CHAPTERS = [
    {
        "dir": "04-request-params",
        "num": "04",
        "title": "Query, Path, Headers",
        "phase": "1",
        "stars": "⭐⭐",
        "hours": "3–4",
        "tasks": 10,
        "goals": [
            "Query parameters: `?page=1&sort=name`",
            "Path parameters и regex constraints",
            "Request headers: Authorization, Content-Type",
            "Parsing и validation query params",
        ],
        "theory": """# Теория: Request Parameters

## Query Parameters

```go
q := r.URL.Query()
page := q.Get("page")       // string, "" if missing
pages := q["tag"]           // []string for repeated params
```

## Parsing with defaults

```go
func queryInt(r *http.Request, key string, defaultVal int) int {
    s := r.URL.Query().Get(key)
    if s == "" { return defaultVal }
    v, err := strconv.Atoi(s)
    if err != nil { return defaultVal }
    return v
}
```

## Path Parameters (chi)

```go
r.Get("/books/{id}", handler)
id := chi.URLParam(r, "id")

// Regex constraint
r.Get("/users/{id:[0-9]+}", handler)
```

## Headers

```go
auth := r.Header.Get("Authorization")
contentType := r.Header.Get("Content-Type")
r.Header.Get("X-Request-ID")  // case-insensitive
```

## Request Body

```go
if r.Header.Get("Content-Type") != "application/json" {
    http.Error(w, "unsupported media type", 415)
    return
}
```
""",
        "exercises": """# Задачи — Глава 04

## ⭐ 1. Query echo — GET /search?q=go&limit=10
## ⭐ 2. Pagination params — page, per_page с defaults (1, 20)
## ⭐ 3. Multi-value — GET /tags?tag=go&tag=api → `{"tags":["go","api"]}`
## ⭐ 4. Path param validation — /books/{id} только числа, иначе 400
## ⭐ 5. Header echo — GET /headers/Authorization
## ⭐⭐ 6. Sort params — ?sort_by=title&order=asc|desc
## ⭐⭐ 7. Filter builder — ?status=active&role=admin
## ⭐⭐ 8. Content-Type check — POST /data только application/json
## ⭐⭐⭐ 9. Query validation helper — pkg с ParsePagination(r)
## ⭐⭐⭐ 10. Books search — GET /books?q=code&author=Martin&year=2008
""",
    },
    {
        "dir": "05-json-and-encoding",
        "num": "05",
        "title": "JSON и encoding",
        "phase": "1",
        "stars": "⭐⭐",
        "hours": "4–5",
        "tasks": 12,
        "goals": [
            "encoding/json: Marshal, Unmarshal, Encoder, Decoder",
            "Struct tags: json, omitempty",
            "Request DTO vs Response DTO",
            "Обработка ошибок декодирования",
        ],
        "theory": """# Теория: JSON в Go

## Struct tags

```go
type Book struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Author    string    `json:"author,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    internal  string    // не экспортируется — не попадёт в JSON
}
```

## Decode request

```go
var req CreateBookRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    respondError(w, 400, "INVALID_JSON", "invalid request body")
    return
}
defer r.Body.Close()
```

## Encode response

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(book)
```

## DTO pattern

```go
type CreateBookRequest struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}

type BookResponse struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}
```

## Null vs omitempty

- `omitempty` — поле пропускается если zero value
- `*string` / `*time.Time` — для nullable JSON null
""",
        "exercises": """# Задачи — Глава 05

## ⭐ 1. POST /books — decode JSON, return created book with id
## ⭐ 2. GET /books/:id — encode struct to JSON
## ⭐ 3. omitempty — optional description field
## ⭐ 4. Invalid JSON → 400 with error message
## ⭐ 5. Empty body POST → 400
## ⭐⭐ 6. Request/Response DTOs — не expose internal fields
## ⭐⭐ 7. time.Time formatting — RFC3339 in JSON
## ⭐⭐ 8. Nullable fields — `*string` for optional author
## ⭐⭐ 9. json.Decoder DisallowUnknownFields()
## ⭐⭐⭐ 10. List response wrapper — `{"data": [...], "count": N}`
## ⭐⭐⭐ 11. Streaming — json.Encoder for large list (1000 books)
## ⭐⭐⭐ 12. Custom UnmarshalJSON for enum status field
""",
    },
    {
        "dir": "06-rest-principles",
        "num": "06",
        "title": "REST принципы",
        "phase": "2",
        "stars": "⭐⭐",
        "hours": "3–4",
        "tasks": 8,
        "goals": [
            "6 constraints REST (Fielding)",
            "Resources vs actions",
            "Stateless communication",
            "Uniform interface",
        ],
        "theory": """# Теория: REST Principles

## 6 ограничений REST

1. **Client-Server** — разделение UI и data storage
2. **Stateless** — каждый запрос содержит всю нужную информацию
3. **Cacheable** — ответы можно кэшировать
4. **Uniform Interface** — единообразие (resources, representations, HATEOAS)
5. **Layered System** — клиент не знает конечный сервер
6. **Code on Demand** (optional) — JS в ответах

## Resource-Oriented Design

```
Resources:  /books, /users, /tasks
NOT:        /getBooks, /createUser, /deleteTask
```

## Representations

Один ресурс — разные форматы:
- `Accept: application/json`
- `Accept: application/xml` (реже)

## Stateless — no server sessions

❌ Session cookie с user state на сервере  
✅ JWT token в Authorization header

## Richardson Maturity Model

| Level | Описание |
|-------|----------|
| 0 | Single URI, single method (RPC) |
| 1 | Resources (multiple URIs) |
| 2 | HTTP verbs |
| 3 | HATEOAS |
""",
        "exercises": """# Задачи — Глава 06

## ⭐ 1. Audit API — найди 10 нарушений REST в docs/bad-api-example.md (создай сам)
## ⭐ 2. Redesign — перепиши RPC API в REST (задачи в exercises)
## ⭐ 3. Resource map — нарисуй resources для blog (posts, comments, users)
## ⭐ 4. Stateless auth — опиши flow с JWT (markdown doc)
## ⭐⭐ 5. Implement RESTful /articles CRUD (in-memory)
## ⭐⭐ 6. Cache headers — ETag + If-None-Match → 304
## ⭐⭐⭐ 7. HATEOAS links in response
## ⭐⭐⭐ 8. Content negotiation — JSON vs XML (optional XML)
""",
    },
    {
        "dir": "07-crud-operations",
        "num": "07",
        "title": "CRUD операции",
        "phase": "2",
        "stars": "⭐⭐",
        "hours": "5–6",
        "tasks": 14,
        "goals": [
            "Full CRUD для ресурса Books",
            "In-memory storage (map + mutex)",
            "Auto-increment ID",
            "Location header on create",
        ],
        "theory": """# Теория: CRUD Operations

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
""",
        "exercises": """# Задачи — Глава 07

## ⭐ 1. POST /books — create with auto ID
## ⭐ 2. GET /books — list all
## ⭐ 3. GET /books/{id} — get one, 404 if missing
## ⭐ 4. PUT /books/{id} — full replace
## ⭐ 5. DELETE /books/{id} — 204 No Content
## ⭐⭐ 6. PATCH /books/{id} — partial update
## ⭐⭐ 7. Duplicate title check → 409 Conflict
## ⭐⭐ 8. Location header on POST
## ⭐⭐ 9. Thread-safe store with sync.RWMutex
## ⭐⭐⭐ 10. Soft delete — deleted_at field, GET excludes deleted
## ⭐⭐⭐ 11. Bulk create — POST /books/bulk
## ⭐⭐⭐ 12. Service layer — handler → service → store
## ⭐⭐⭐ 13. Input validation — title required, min length 1
## ⭐⭐⭐ 14. Integration test script (shell/curl) for all endpoints
""",
    },
    {
        "dir": "08-status-codes",
        "num": "08",
        "title": "HTTP статус-коды",
        "phase": "2",
        "stars": "⭐⭐",
        "hours": "3–4",
        "tasks": 10,
        "goals": [
            "Правильные коды для каждого сценария",
            "Единый формат ошибок",
            "errors package + sentinel errors",
        ],
        "theory": """# Теория: Status Codes

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
""",
        "exercises": """# Задачи — Глава 08

## ⭐ 1. respondError helper — единый формат
## ⭐ 2. 404 — book not found
## ⭐ 3. 400 — invalid JSON
## ⭐ 4. 405 — method not allowed with Allow header
## ⭐ 5. 409 — duplicate email
## ⭐⭐ 6. 422 — validation errors with details[]
## ⭐⭐ 7. Error codes enum — NOT_FOUND, VALIDATION_ERROR, etc.
## ⭐⭐ 8. Never expose internal errors in 500
## ⭐⭐⭐ 9. handleServiceError mapper
## ⭐⭐⭐ 10. Test all status codes with table-driven tests
""",
    },
    {
        "dir": "09-api-versioning",
        "num": "09",
        "title": "Версионирование API",
        "phase": "2",
        "stars": "⭐⭐⭐",
        "hours": "3–4",
        "tasks": 8,
        "goals": [
            "URL versioning /api/v1, /api/v2",
            "Breaking vs non-breaking changes",
            "Deprecation headers",
        ],
        "theory": """# Теория: API Versioning

## URL prefix (recommended)

```go
r.Route("/api/v1", v1Routes)
r.Route("/api/v2", v2Routes)
```

## Breaking changes

- Remove field
- Change field type
- Change URL structure
- Change status code semantics

## Non-breaking

- Add optional field
- Add new endpoint
- Add new query param

## Deprecation

```go
w.Header().Set("Deprecation", "true")
w.Header().Set("Sunset", "Sat, 01 Jan 2028 00:00:00 GMT")
w.Header().Set("Link", "</api/v2/books>; rel=\"successor-version\"")
```
""",
        "exercises": """# Задачи — Глава 09

## ⭐ 1. Mount v1 and v2 routers
## ⭐ 2. v2 Book adds `isbn` field — v1 unchanged
## ⭐ 3. Deprecation header on v1
## ⭐ 4. Version in response meta
## ⭐⭐ 5. Shared service, different handlers/DTOs
## ⭐⭐ 6. Default version redirect /api/books → /api/v1/books
## ⭐⭐⭐ 7. Document breaking changes in CHANGELOG.md
## ⭐⭐⭐ 8. Accept header versioning (bonus)
""",
    },
    {
        "dir": "10-pagination-filtering",
        "num": "10",
        "title": "Pagination и фильтрация",
        "phase": "2",
        "stars": "⭐⭐⭐",
        "hours": "4–5",
        "tasks": 12,
        "goals": [
            "Offset pagination (page/per_page)",
            "Sorting и filtering",
            "Response metadata",
            "Cursor pagination (bonus)",
        ],
        "theory": """# Теория: Pagination

## Offset pagination

```
GET /books?page=2&per_page=20&sort_by=title&sort_order=asc&author=Martin
```

```go
type Pagination struct {
    Page       int `json:"page"`
    PerPage    int `json:"per_page"`
    Total      int `json:"total"`
    TotalPages int `json:"total_pages"`
}

type ListResponse[T any] struct {
    Data       []T        `json:"data"`
    Pagination Pagination `json:"pagination"`
}
```

## Cursor pagination (better for large datasets)

```
GET /books?cursor=eyJpZCI6MTAwfQ&limit=20
```

Response includes `next_cursor`.
""",
        "exercises": """# Задачи — Глава 10

## ⭐ 1. page/per_page with defaults
## ⭐ 2. Response pagination metadata
## ⭐ 3. sort_by whitelist (title, created_at, id)
## ⭐ 4. Filter by author query param
## ⭐⭐ 5. Multiple filters combined
## ⭐⭐ 6. Max per_page = 100
## ⭐⭐ 7. Generic ListResponse[T]
## ⭐⭐⭐ 8. Cursor pagination implementation
## ⭐⭐⭐ 9. Full-text search in title (simple strings.Contains)
## ⭐⭐⭐ 10. Filter by date range ?from=2024-01-01&to=2024-12-31
## ⭐⭐⭐ 11. SQL-ready filter builder struct
## ⭐⭐⭐ 12. Seed 100 books + test pagination edge cases
""",
    },
]

# Phase 3-6 chapters (shorter generation)
PHASE3_6 = [
    ("11-layered-architecture", "11", "Слоистая архитектура", "3", "⭐⭐⭐", "5–6", 10,
     "Handler → Service → Repository. Domain models. Error propagation."),
    ("12-dependency-injection", "12", "Dependency Injection", "3", "⭐⭐⭐", "4–5", 8,
     "Constructor injection. Interfaces. main.go wiring. No global state."),
    ("13-configuration", "13", "Конfiguration", "3", "⭐⭐⭐", "3–4", 10,
     "env vars, flags, config struct, validation, .env file."),
    ("14-logging", "14", "Structured logging", "3", "⭐⭐⭐", "3–4", 8,
     "slog, log levels, request_id in logs, JSON logs."),
    ("15-middleware", "15", "Middleware chain", "3", "⭐⭐⭐", "4–5", 12,
     "Auth, logging, recovery, timeout, rate limit basics."),
    ("16-postgresql-pgx", "16", "PostgreSQL + pgx", "4", "⭐⭐⭐", "5–6", 12,
     "pgxpool, queries, context, docker-compose postgres."),
    ("17-migrations", "17", "Миграции goose", "4", "⭐⭐⭐", "3–4", 10,
     "SQL migrations, up/down, embed migrations."),
    ("18-repository-pattern", "18", "Repository pattern", "4", "⭐⭐⭐⭐", "5–6", 12,
     "Interface in domain, postgres impl, swap in-memory."),
    ("19-transactions", "19", "Транзакции", "4", "⭐⭐⭐⭐", "4–5", 10,
     "Begin, Commit, Rollback. Transfer money example."),
    ("20-validation", "20", "Валидация", "4", "⭐⭐⭐", "3–4", 10,
     "go-playground/validator, custom tags, 422 responses."),
    ("21-jwt-auth", "21", "JWT Authentication", "5", "⭐⭐⭐⭐", "6–8", 14,
     "Register, login, bcrypt, JWT access token, middleware."),
    ("22-rbac", "22", "RBAC Authorization", "5", "⭐⭐⭐⭐", "4–5", 12,
     "Roles admin/user, permission checks, owner-only resources."),
    ("23-security", "23", "Rate limit, CORS, Security", "5", "⭐⭐⭐⭐", "4–5", 10,
     "Rate limiter, CORS, security headers, Redis rate limit."),
    ("24-testing", "24", "Тестирование API", "6", "⭐⭐⭐⭐", "6–8", 14,
     "httptest, testify, table tests, integration tests."),
    ("25-production", "25", "OpenAPI, Docker, Graceful shutdown", "6", "⭐⭐⭐⭐⭐", "6–8", 12,
     "openapi.yaml, Dockerfile, multi-stage build, health probes."),
]

def write_chapter(ch):
    d = os.path.join(BASE, ch["dir"])
    os.makedirs(os.path.join(d, "starter"), exist_ok=True)
    
    with open(os.path.join(d, "README.md"), "w") as f:
        f.write(f"""# Глава {ch['num']} — {ch['title']}

> **Фаза {ch['phase']}** · {ch['stars']} · ~{ch['hours']} часов · {ch['tasks']} задач

## Цели

""")
        for g in ch["goals"]:
            f.write(f"- {g}\n")
        f.write(f"""
## Материалы

- [theory.md](./theory.md) · [exercises.md](./exercises.md) · [checklist.md](./checklist.md) · [hints.md](./hints.md)

## Workspace

`workspace/ch{ch['num']}/`
""")

    with open(os.path.join(d, "theory.md"), "w") as f:
        f.write(ch["theory"])
    
    with open(os.path.join(d, "exercises.md"), "w") as f:
        f.write(ch["exercises"])
    
    with open(os.path.join(d, "checklist.md"), "w") as f:
        f.write(f"""# Checklist — Глава {ch['num']}

- [ ] Прочитал theory.md
- [ ] Выполнил все задачи ⭐ и ⭐⭐
- [ ] Код в workspace/ch{ch['num']}/ работает
- [ ] Git commit
""")
    
    with open(os.path.join(d, "hints.md"), "w") as f:
        f.write(f"""# Подсказки — Глава {ch['num']}

> Открывай после 30 минут самостоятельной работы.

Смотри theory.md и starter/. Если застрял — перечитай предыдущие главы.
Документация: [docs/](../../docs/)
""")

def write_simple_chapter(dir_name, num, title, phase, stars, hours, tasks, summary):
    d = os.path.join(BASE, dir_name)
    os.makedirs(os.path.join(d, "starter"), exist_ok=True)
    
    with open(os.path.join(d, "README.md"), "w") as f:
        f.write(f"""# Глава {num} — {title}

> **Фаза {phase}** · {stars} · ~{hours} часов · {tasks} задач

## Описание

{summary}

## Материалы

- [theory.md](./theory.md) · [exercises.md](./exercises.md) · [checklist.md](./checklist.md) · [hints.md](./hints.md)

## Workspace

`workspace/ch{num}/`

## Docker (если нужен)

```bash
make docker-up   # PostgreSQL + Redis с главы 16
```
""")
    
    with open(os.path.join(d, "theory.md"), "w") as f:
        f.write(f"# Теория: {title}\n\n{summary}\n\nПодробнее: [docs/](../../docs/)\n")
    
    with open(os.path.join(d, "exercises.md"), "w") as f:
        f.write(f"# Задачи — Глава {num}\n\n")
        for i in range(1, min(tasks+1, 11)):
            stars_ex = "⭐" if i <= 3 else "⭐⭐" if i <= 7 else "⭐⭐⭐"
            f.write(f"## {stars_ex} {i}. Задача {i} — см. README и theory\n\n")
        f.write(f"\n_Полный список из {tasks} задач — реализуй по theory.md и docs._\n")
    
    with open(os.path.join(d, "checklist.md"), "w") as f:
        f.write(f"# Checklist — Глава {num}\n\n- [ ] Theory прочитана\n- [ ] Задачи выполнены\n- [ ] Git commit\n")
    
    with open(os.path.join(d, "hints.md"), "w") as f:
        f.write(f"# Подсказки — Глава {num}\n\nСмотри starter/ и предыдущие главы.\n")

for ch in CHAPTERS:
    write_chapter(ch)

for item in PHASE3_6:
    write_simple_chapter(*item)

print("Generated", len(CHAPTERS) + len(PHASE3_6), "chapters")
