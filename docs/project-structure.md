# Project Structure — Структура Go REST API

## Рекомендуемая структура (с главы 11)

```
taskflow-api/
├── cmd/
│   └── server/
│       └── main.go              # точка входа, wiring
├── internal/
│   ├── config/
│   │   └── config.go            # env, flags
│   ├── domain/
│   │   ├── user.go              # сущности + интерфейсы
│   │   ├── task.go
│   │   └── errors.go            # domain errors
│   ├── handler/
│   │   ├── user_handler.go      # HTTP handlers
│   │   ├── task_handler.go
│   │   └── response.go          # JSON helpers
│   ├── service/
│   │   ├── user_service.go      # бизнес-логика
│   │   └── task_service.go
│   ├── repository/
│   │   ├── postgres/
│   │   │   ├── user_repo.go
│   │   │   └── task_repo.go
│   │   └── repository.go        # интерфейсы
│   └── middleware/
│       ├── auth.go
│       ├── logging.go
│       └── recovery.go
├── migrations/
│   ├── 001_create_users.sql
│   └── 002_create_tasks.sql
├── api/
│   └── openapi.yaml
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum
```

## Правила

### `cmd/` — точки входа

Один бинарник = одна папка в `cmd/`. `main.go` только собирает зависимости и запускает сервер.

### `internal/` — приватный код

Go запрещает импорт `internal/` извне модуля. Вся логика приложения здесь.

### Слои и зависимости

```
handler → service → repository → database
   ↓         ↓          ↓
 domain    domain     domain
```

- **Handler** знает про HTTP (`http.Request`, status codes)
- **Service** знает про бизнес-правила, не знает про HTTP
- **Repository** знает про SQL, не знает про бизнес-правила
- **Domain** — чистые типы и интерфейсы, без зависимостей

### `pkg/` — только если переиспользуешь

Для учебного проекта `pkg/` не нужен. Если библиотека нужна нескольким сервисам — тогда `pkg/httputil/`.

## Плохие структуры

```
❌ models/, controllers/, routes/     # PHP-style
❌ everything in main.go               # до главы 02 ок, дальше — нет
❌ utils/ с 50 функциями               # god package
❌ handler вызывает SQL напрямую       # пропуск service layer
```

## Naming conventions

| Что | Convention | Пример |
|-----|------------|--------|
| Package | lowercase, singular | `handler`, `service` |
| File | snake_case | `user_handler.go` |
| Interface | noun или -er | `UserRepository`, `TaskStore` |
| Handler func | HTTP method + resource | `CreateTask`, `ListTasks` |
| Error vars | Err prefix | `ErrNotFound`, `ErrConflict` |
