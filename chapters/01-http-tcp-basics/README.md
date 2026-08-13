# Глава 01 — HTTP и TCP основы

> **Фаза 1** · ⭐ · ~2–3 часа · 8 задач

## Цели обучения

После этой главы ты будешь понимать:
- Как данные передаются по сети (TCP → HTTP)
- Структуру HTTP-запроса и ответа
- Разницу между методами GET и POST на уровне протокола
- Как использовать `curl` для отладки

## Что делаем

1. Читаем теорию: [theory.md](./theory.md)
2. Запускаем starter и экспериментируем
3. Выполняем задачи: [exercises.md](./exercises.md)
4. Пишем решение в `workspace/ch01/`
5. Проверяем: [checklist.md](./checklist.md)

## Быстрый старт

```bash
# Запусти starter
make run-ch01

# В другом терминале
curl -v http://localhost:8080/
curl -v http://localhost:8080/hello
```

## Starter

Минимальный HTTP-сервер в [starter/main.go](./starter/main.go) — изучи построчно.

## Связь с REST

REST API — это HTTP + соглашения. Без понимания HTTP невозможно строить API. Эта глава — фундамент.

## Следующая глава

→ [Глава 02: net/http сервер](../02-net-http-server/)
