# Go REST Mastery

**Огромный практический курс по REST API на Go** — 25 глав, 200+ задач, финальный capstone-проект.

Ты пройдёшь путь от «что такое HTTP» до production-ready API с PostgreSQL, JWT, тестами, Docker и OpenAPI.

> Уже знаешь синтаксис Go? Отлично. Если нет — сначала пройди [go-syntax-concurrency](../go-syntax-concurrency/), затем возвращайся сюда.
>
> Нужен более широкий backend-трек (TCP, gRPC, Redis)? → [backend-go-mastery](../backend-go-mastery/)

---

## Быстрый старт

```bash
# 1. Клонируй / открой проект
cd ~/Projects/go-rest-mastery

# 2. Зависимости
make deps

# 3. Открой PROGRESS.md и отметь главу 01
# 4. Читай chapters/01-http-tcp-basics/README.md
# 5. Пиши код в workspace/ch01/
```

---

## Как проходить курс

1. **Строго по порядку** — каждая глава опирается на предыдущую.
2. Читай `README.md` главы → `theory.md` → выполняй задачи из `exercises.md`.
3. Пиши код в папке `workspace/` (создай `workspace/ch01/`, `workspace/ch02/` и т.д.).
4. Отмечай прогресс в [`PROGRESS.md`](./PROGRESS.md).
5. **30 минут без Google** — сначала попробуй сам, потом смотри подсказки в `hints.md`.
6. После каждой главы — `git commit` с описанием что сделал.

### Структура каждой главы

```
chapters/05-json-and-encoding/
├── README.md       ← начни отсюда (обзор + цели)
├── theory.md       ← теория с примерами
├── exercises.md    ← задачи (⭐ → ⭐⭐⭐⭐⭐)
├── checklist.md    ← самопроверка перед переходом
├── hints.md        ← подсказки (открывай после 30 мин)
└── starter/        ← стартовый код для экспериментов
```

---

## Стек технологий

| Компонент | Технология | С какой главы |
|-----------|------------|---------------|
| HTTP | `net/http` | 02 |
| Роутер | [chi](https://github.com/go-chi/chi) | 03 |
| JSON | `encoding/json` | 05 |
| Валидация | [validator](https://github.com/go-playground/validator) | 20 |
| PostgreSQL | [pgx](https://github.com/jackc/pgx) | 16 |
| Миграции | [goose](https://github.com/pressly/goose) | 17 |
| JWT | [golang-jwt](https://github.com/golang-jwt/jwt) | 21 |
| Redis | [go-redis](https://github.com/redis/go-redis) | 23 |
| Тесты | `testing`, `httptest`, testify | 24 |
| Контейнеры | Docker Compose | 16+ |

---

## Дорожная карта (25 глав)

### Фаза 1 — HTTP и сеть (главы 01–05)

| # | Глава | Задач | Сложность |
|---|-------|-------|-----------|
| 01 | [HTTP и TCP основы](./chapters/01-http-tcp-basics/) | 8 | ⭐ |
| 02 | [net/http сервер](./chapters/02-net-http-server/) | 10 | ⭐ |
| 03 | [Маршрутизация с chi](./chapters/03-routing-chi/) | 12 | ⭐⭐ |
| 04 | [Query, Path, Headers](./chapters/04-request-params/) | 10 | ⭐⭐ |
| 05 | [JSON и encoding](./chapters/05-json-and-encoding/) | 12 | ⭐⭐ |

### Фаза 2 — REST API Design (главы 06–10)

| # | Глава | Задач | Сложность |
|---|-------|-------|-----------|
| 06 | [REST принципы](./chapters/06-rest-principles/) | 8 | ⭐⭐ |
| 07 | [CRUD операции](./chapters/07-crud-operations/) | 14 | ⭐⭐ |
| 08 | [HTTP статус-коды](./chapters/08-status-codes/) | 10 | ⭐⭐ |
| 09 | [Версионирование API](./chapters/09-api-versioning/) | 8 | ⭐⭐⭐ |
| 10 | [Pagination и фильтрация](./chapters/10-pagination-filtering/) | 12 | ⭐⭐⭐ |

### Фаза 3 — Архитектура (главы 11–15)

| # | Глава | Задач | Сложность |
|---|-------|-------|-----------|
| 11 | [Слоистая архитектура](./chapters/11-layered-architecture/) | 10 | ⭐⭐⭐ |
| 12 | [Dependency Injection](./chapters/12-dependency-injection/) | 8 | ⭐⭐⭐ |
| 13 | [Конфигурация](./chapters/13-configuration/) | 10 | ⭐⭐⭐ |
| 14 | [Structured logging](./chapters/14-logging/) | 8 | ⭐⭐⭐ |
| 15 | [Middleware chain](./chapters/15-middleware/) | 12 | ⭐⭐⭐ |

### Фаза 4 — Данные (главы 16–20)

| # | Глава | Задач | Сложность |
|---|-------|-------|-----------|
| 16 | [PostgreSQL + pgx](./chapters/16-postgresql-pgx/) | 12 | ⭐⭐⭐ |
| 17 | [Миграции goose](./chapters/17-migrations/) | 10 | ⭐⭐⭐ |
| 18 | [Repository pattern](./chapters/18-repository-pattern/) | 12 | ⭐⭐⭐⭐ |
| 19 | [Транзакции](./chapters/19-transactions/) | 10 | ⭐⭐⭐⭐ |
| 20 | [Валидация входных данных](./chapters/20-validation/) | 10 | ⭐⭐⭐ |

### Фаза 5 — Безопасность (главы 21–23)

| # | Глава | Задач | Сложность |
|---|-------|-------|-----------|
| 21 | [JWT Authentication](./chapters/21-jwt-auth/) | 14 | ⭐⭐⭐⭐ |
| 22 | [RBAC Authorization](./chapters/22-rbac/) | 12 | ⭐⭐⭐⭐ |
| 23 | [Rate limit, CORS, Security](./chapters/23-security/) | 10 | ⭐⭐⭐⭐ |

### Фаза 6 — Production (главы 24–25)

| # | Глава | Задач | Сложность |
|---|-------|-------|-----------|
| 24 | [Тестирование API](./chapters/24-testing/) | 14 | ⭐⭐⭐⭐ |
| 25 | [OpenAPI, Docker, Graceful shutdown](./chapters/25-production/) | 12 | ⭐⭐⭐⭐⭐ |

### Capstone — TaskFlow API

| Проект | Описание | Задач |
|--------|----------|-------|
| [TaskFlow API](./capstone/) | Полноценный REST API для управления задачами | 40+ |

---

## Документация

| Документ | Описание |
|----------|----------|
| [Getting Started](./docs/getting-started.md) | Установка Go, IDE, первые шаги |
| [REST Conventions](./docs/rest-conventions.md) | Соглашения именования, URL design |
| [HTTP Status Codes](./docs/http-status-codes.md) | Когда какой код возвращать |
| [API Design Guide](./docs/api-design-guide.md) | Проектирование REST API |
| [Project Structure](./docs/project-structure.md) | Как организовать Go-проект |
| [Glossary](./docs/glossary.md) | Словарь терминов |

---

## Workspace — твой код

Все решения пишешь в `workspace/`:

```
workspace/
├── ch01/          # решения главы 01
├── ch02/
├── ...
├── ch25/
└── capstone/      # финальный проект
```

Каждая папка — отдельный Go-модуль или подпакет. Пример:

```bash
mkdir -p workspace/ch07
cd workspace/ch07
go mod init github.com/YOUR_USERNAME/ch07-books-api
# пиши код, запускай: go run ./cmd/server
```

---

## Оценка времени

| Фаза | Главы | Примерно часов |
|------|-------|----------------|
| 1 | 01–05 | 15–25 |
| 2 | 06–10 | 20–30 |
| 3 | 11–15 | 25–35 |
| 4 | 16–20 | 30–40 |
| 5 | 21–23 | 25–35 |
| 6 | 24–25 | 20–30 |
| Capstone | — | 40–60 |
| **Итого** | | **175–255 часов** |

Это **2–4 месяца** при 2–3 часах в день.

---

## Полезные команды

```bash
make deps          # зависимости
make docker-up     # PostgreSQL + Redis + Adminer
make run-ch01      # запустить starter главы 01
make test          # тесты в workspace/
```

Adminer (веб-UI для PostgreSQL): http://localhost:8081  
Логин: `restapi` / `restapi` / БД: `taskflow`

---

## Лицензия

MIT — учись, копируй, делись.
