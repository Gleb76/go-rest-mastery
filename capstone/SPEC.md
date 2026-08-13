# TaskFlow API — Спецификация

## Описание

TaskFlow — REST API для управления проектами и задачами (аналог упрощённого Jira/Trello).

## Base URL

```
http://localhost:8080/api/v1
```

## Resources

### Users

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | /auth/register | ❌ | Регистрация |
| POST | /auth/login | ❌ | Логин → JWT |
| GET | /users/me | ✅ | Текущий пользователь |
| PATCH | /users/me | ✅ | Обновить профиль |

**User model:**
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "John Doe",
  "role": "user",
  "created_at": "2026-01-01T00:00:00Z"
}
```

### Projects

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | /projects | ✅ | Список проектов пользователя |
| POST | /projects | ✅ | Создать проект |
| GET | /projects/{id} | ✅ | Получить проект |
| PATCH | /projects/{id} | ✅ | Обновить (owner/admin) |
| DELETE | /projects/{id} | ✅ | Удалить (owner/admin) |

**Project model:**
```json
{
  "id": "uuid",
  "name": "Backend API",
  "description": "REST API project",
  "owner_id": "uuid",
  "created_at": "...",
  "updated_at": "..."
}
```

### Tasks

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | /projects/{projectID}/tasks | ✅ | Задачи проекта |
| POST | /projects/{projectID}/tasks | ✅ | Создать задачу |
| GET | /tasks/{id} | ✅ | Получить задачу |
| PATCH | /tasks/{id} | ✅ | Обновить |
| DELETE | /tasks/{id} | ✅ | Удалить |

**Query params для GET /projects/{id}/tasks:**
- `page`, `per_page` — pagination
- `status` — todo | in_progress | done
- `assignee_id` — фильтр по исполнителю
- `sort_by` — created_at | priority | title
- `sort_order` — asc | desc

**Task model:**
```json
{
  "id": "uuid",
  "project_id": "uuid",
  "title": "Implement auth",
  "description": "JWT + bcrypt",
  "status": "in_progress",
  "priority": 1,
  "assignee_id": "uuid",
  "due_date": "2026-02-01T00:00:00Z",
  "created_at": "...",
  "updated_at": "..."
}
```

### Comments

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | /tasks/{taskID}/comments | ✅ | Комментарии |
| POST | /tasks/{taskID}/comments | ✅ | Добавить |
| DELETE | /comments/{id} | ✅ | Удалить (author/admin) |

## Auth

```
Authorization: Bearer <jwt_token>
```

JWT claims: `sub` (user_id), `role`, `exp`, `iat`

## Roles

| Role | Permissions |
|------|-------------|
| admin | Всё |
| user | CRUD своих projects, tasks в своих projects |
| viewer | Только read |

## Error format

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input",
    "details": [{"field": "title", "message": "required"}]
  }
}
```

## Database schema

```sql
-- users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- projects
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- tasks
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'todo',
    priority INT NOT NULL DEFAULT 0,
    assignee_id UUID REFERENCES users(id),
    due_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- comments
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    author_id UUID NOT NULL REFERENCES users(id),
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

## Non-functional requirements

- [ ] Graceful shutdown
- [ ] Structured JSON logs (slog)
- [ ] Request ID in logs
- [ ] Health: GET /health
- [ ] Ready: GET /ready (checks DB)
- [ ] Rate limit: 100 req/min per IP
- [ ] CORS for localhost:3000
- [ ] Dockerfile multi-stage
- [ ] docker-compose with app + postgres + redis
- [ ] OpenAPI 3.0 spec
- [ ] Integration tests for auth + CRUD flow

## Milestones

### Milestone 1 — MVP (40 задач)
In-memory или PostgreSQL. Projects + Tasks CRUD без auth.

### Milestone 2 — Auth (10 задач)
Register, login, protected routes.

### Milestone 3 — Full (20 задач)
Comments, RBAC, pagination, validation.

### Milestone 4 — Production (10 задач)
Tests, Docker, OpenAPI, CI-ready README.

## Пример flow

```bash
# Register
curl -X POST localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"dev@test.com","password":"secret123","name":"Developer"}'

# Login
TOKEN=$(curl -s -X POST localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"dev@test.com","password":"secret123"}' | jq -r .token)

# Create project
curl -X POST localhost:8080/api/v1/projects \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"My Project","description":"First project"}'

# Create task
curl -X POST localhost:8080/api/v1/projects/PROJECT_ID/tasks \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Setup repo","status":"todo","priority":1}'
```
