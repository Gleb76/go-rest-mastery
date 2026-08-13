# Задачи — Глава 02

## ⭐ 1. Health check
`GET /health` → `{"status":"ok","timestamp":"..."}`

## ⭐ 2. Readiness
`GET /ready` → 200 если сервер готов, 503 если `ready=false` (переключай флаг через `POST /admin/ready`)

## ⭐ 3. Server info
`GET /info` → version, go_version, uptime_seconds

## ⭐ 4. Method routing
Один path `/resource` — разные handlers для GET, POST, DELETE (405 для остальных)

## ⭐⭐ 5. Request ID
Каждый ответ содержит header `X-Request-ID` (UUID). Генерируй в middleware-подобной обёртке.

## ⭐⭐ 6. Timeout handler
`GET /slow?seconds=5` — спит N секунд. Настрой `WriteTimeout: 3s` — клиент должен получить timeout.

## ⭐⭐ 7. Structured handlers
Вынеси handlers в `internal/handler/`. Main только wiring.

## ⭐⭐ 8. http.Server
Замени `ListenAndServe` на `http.Server` с ReadTimeout/WriteTimeout.

## ⭐⭐⭐ 9. Graceful shutdown
Ctrl+C → сервер завершает текущие запросы, логирует shutdown.

## ⭐⭐⭐ 10. API prefix
Все routes под `/api/v1/`. `/health` остаётся на root.

## Проверка
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/info
```
