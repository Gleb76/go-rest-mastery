# Глава 02 — net/http сервер

> **Фаза 1** · ⭐ · ~3–4 часа · 10 задач

## Цели

- `http.ServeMux` vs `http.DefaultServeMux`
- Handler functions и Handler interface
- Организация нескольких endpoints
- `http.Server` с таймаутами
- Graceful shutdown (базовый)

## Материалы

- [theory.md](./theory.md)
- [exercises.md](./exercises.md)
- [checklist.md](./checklist.md)
- [hints.md](./hints.md)

## Starter

```bash
make run-ch02
curl http://localhost:8080/health
```

## Workspace

`workspace/ch02/` — реализуй полноценный multi-endpoint сервер.

## Следующая

→ [Глава 03: chi router](../03-routing-chi/)
