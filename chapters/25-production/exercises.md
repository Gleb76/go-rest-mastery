# Задачи — Глава 25

## ⭐ 1. openapi.yaml — document all endpoints

## ⭐ 2. Dockerfile multi-stage (build + alpine runtime)

## ⭐ 3. docker-compose: app + postgres + redis

## ⭐ 4. Graceful shutdown с context timeout

## ⭐ 5. GET /health и GET /ready (DB ping)

## ⭐⭐ 6. Environment config — 12-factor app

## ⭐⭐ 7. Non-root user in Docker

## ⭐⭐ 8. .dockerignore

## ⭐⭐⭐ 9. Swagger UI mount /docs (optional)

## ⭐⭐⭐ 10. Kubernetes-style liveness/readiness probes doc

## ⭐⭐⭐ 11. Production README: deploy, env vars, migrations

## ⭐⭐⭐ 12. Full capstone deploy: docker compose up → API works

## Dockerfile example

```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /server ./cmd/server

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
COPY --from=builder /server /server
EXPOSE 8080
USER nobody
ENTRYPOINT ["/server"]
```
