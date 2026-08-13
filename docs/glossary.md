# Glossary — Словарь терминов

| Термин | Определение |
|--------|-------------|
| **REST** | Representational State Transfer — архитектурный стиль для web API |
| **Resource** | Сущность, доступная через URL (book, user, task) |
| **Endpoint** | URL + HTTP метод (GET /books) |
| **Handler** | Функция, обрабатывающая HTTP запрос |
| **Middleware** | Функция-обёртка вокруг handler (logging, auth) |
| **Router** | Компонент, сопоставляющий URL+method с handler |
| **CRUD** | Create, Read, Update, Delete |
| **DTO** | Data Transfer Object — структура для API request/response |
| **Entity** | Доменная модель (может отличаться от DTO) |
| **Repository** | Слой доступа к данным (абстракция над БД) |
| **Service** | Слой бизнес-логики |
| **Migration** | SQL-скрипт изменения схемы БД |
| **JWT** | JSON Web Token — формат токена для auth |
| **RBAC** | Role-Based Access Control — права по ролям |
| **Idempotent** | Повторный запрос даёт тот же результат |
| **Safe** | Метод не изменяет состояние сервера (GET) |
| **Pagination** | Разбиение больших списков на страницы |
| **Graceful shutdown** | Корректное завершение: дождаться текущих запросов |
| **Connection pool** | Пул переиспользуемых соединений с БД |
| **Transaction** | Атомарная группа SQL операций |
| **CORS** | Cross-Origin Resource Sharing — политика браузера |
| **Rate limiting** | Ограничение количества запросов |
| **OpenAPI** | Спецификация описания REST API (Swagger) |
| **Health check** | Endpoint для проверки живости сервиса |
| **Structured logging** | Логи в JSON с полями (level, request_id, ...) |
