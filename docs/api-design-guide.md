# API Design Guide

## Проектирование REST API — пошагово

### Шаг 1: Определи ресурсы

Для TaskFlow (capstone):

| Ресурс | Описание |
|--------|----------|
| users | Пользователи |
| projects | Проекты |
| tasks | Задачи |
| comments | Комментарии к задачам |
| tags | Метки |

### Шаг 2: Нарисуй endpoints

```
POST   /api/v1/auth/register
POST   /api/v1/auth/login
GET    /api/v1/users/me
PATCH  /api/v1/users/me

GET    /api/v1/projects
POST   /api/v1/projects
GET    /api/v1/projects/:id
PATCH  /api/v1/projects/:id
DELETE /api/v1/projects/:id

GET    /api/v1/projects/:id/tasks
POST   /api/v1/projects/:id/tasks
GET    /api/v1/tasks/:id
PATCH  /api/v1/tasks/:id
DELETE /api/v1/tasks/:id

GET    /api/v1/tasks/:id/comments
POST   /api/v1/tasks/:id/comments
```

### Шаг 3: Определи модели данных

```go
type Task struct {
    ID          uuid.UUID  `json:"id"`
    ProjectID   uuid.UUID  `json:"project_id"`
    Title       string     `json:"title"`
    Description *string    `json:"description,omitempty"`
    Status      TaskStatus `json:"status"`
    Priority    int        `json:"priority"`
    AssigneeID  *uuid.UUID `json:"assignee_id,omitempty"`
    DueDate     *time.Time `json:"due_date,omitempty"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskStatus string

const (
    TaskStatusTodo       TaskStatus = "todo"
    TaskStatusInProgress TaskStatus = "in_progress"
    TaskStatusDone       TaskStatus = "done"
)
```

### Шаг 4: Спроектируй ошибки

Единый формат для всего API — см. [rest-conventions.md](./rest-conventions.md).

### Шаг 5: Pagination с первого дня

Любой `GET` коллекции должен поддерживать:

```
GET /api/v1/tasks?page=1&per_page=20&sort_by=created_at&sort_order=desc
GET /api/v1/tasks?status=todo&assignee_id=xxx
```

### Шаг 6: Документируй

OpenAPI 3.0 spec — с главы 25. Начинай вести `openapi.yaml` с главы 07.

## Anti-patterns

| Anti-pattern | Проблема | Решение |
|--------------|----------|---------|
| RPC-style URLs | `/createTask`, `/deleteUser` | REST ресурсы + HTTP методы |
| Fat controllers | Вся логика в handler | Handler → Service → Repository |
| Returning passwords | `"password": "hash"` в JSON | `json:"-"` на поле |
| No pagination | `GET /tasks` → 100k записей | page + per_page |
| Inconsistent naming | `userId` vs `user_id` | snake_case в JSON |

## Checklist перед деплоем

- [ ] Все endpoints задокументированы
- [ ] Единый формат ошибок
- [ ] Auth на protected routes
- [ ] Input validation
- [ ] Rate limiting
- [ ] Health check: `GET /health`
- [ ] Graceful shutdown
- [ ] Structured logs с request_id
