# Capstone — TaskFlow API

**Финальный проект** — полноценный REST API для управления задачами.

После 25 глав ты соберёшь всё воедино: PostgreSQL, JWT, RBAC, тесты, Docker, OpenAPI.

## Этапы

| Этап | Что реализовать | Главы |
|------|-----------------|-------|
| 1 — MVP | Users, Projects, Tasks CRUD | 07–18 |
| 2 — Auth | Register, Login, JWT | 21 |
| 3 — RBAC | Roles, permissions | 22 |
| 4 — Polish | Pagination, validation, errors | 08–10, 20 |
| 5 — Production | Tests, Docker, OpenAPI, graceful shutdown | 24–25 |

## Спецификация

Полная спецификация: [SPEC.md](./SPEC.md)

## Starter

```bash
make docker-up
cd capstone/starter
go run ./cmd/server
```

Стартовый код — skeleton. **Реализуй сам** в `workspace/capstone/`.

## Оценка

| Критерий | Баллы |
|----------|-------|
| CRUD всех ресурсов | 20 |
| PostgreSQL + migrations | 15 |
| JWT auth | 15 |
| RBAC | 10 |
| Pagination/filtering | 10 |
| Error handling | 10 |
| Tests (>70% handlers) | 10 |
| Docker + README | 10 |
| OpenAPI spec | 10 |

**100 баллов = mastery 🏆**

## Workspace

```
workspace/capstone/
├── cmd/server/main.go
├── internal/
├── migrations/
├── api/openapi.yaml
├── Dockerfile
├── docker-compose.yml
└── README.md
```

Не копируй starter/ — используй как reference.
