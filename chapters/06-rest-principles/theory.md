# Теория: REST Principles

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
