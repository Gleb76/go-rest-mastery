# Getting Started

## Требования

- **Go 1.22+** — проверь: `go version`
- **Git**
- **Docker Desktop** (с главы 16)
- **curl** или **httpie** для тестирования API
- Редактор: VS Code / Cursor с расширением Go

## Установка Go (macOS)

```bash
brew install go
go version   # должно быть go1.22+
```

## Настройка IDE

1. Установи расширение **Go** (golang.go)
2. `Cmd+Shift+P` → `Go: Install/Update Tools` → выбери все
3. Включи format on save

## Первый запуск

```bash
cd ~/Projects/go-rest-mastery
make deps

# Запусти starter главы 01
make run-ch01
# В другом терминале:
curl http://localhost:8080/
```

## Структура workspace

Создай папку для своих решений:

```bash
mkdir -p workspace/ch01
cd workspace/ch01
go mod init github.com/$(whoami)/ch01-http-basics
```

Пиши код **сам** — starter/ в главах только для экспериментов, не копируй слепо.

## Тестирование API

### curl

```bash
curl -v http://localhost:8080/api/v1/books
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code","author":"Robert Martin"}'
```

### httpie (рекомендуется)

```bash
brew install httpie
http GET localhost:8080/api/v1/books
http POST localhost:8080/api/v1/books title="Clean Code" author="Robert Martin"
```

## Docker (с главы 16)

```bash
make docker-up
# PostgreSQL: localhost:5432, user/pass/db: restapi/taskflow
# Adminer UI: http://localhost:8081
make docker-down
```

## Если застрял

1. Перечитай `theory.md` главы
2. Посмотри `hints.md` (после 30 минут самостоятельной работы)
3. Запусти `starter/` и поэкспериментируй
4. Проверь `checklist.md` — возможно пропустил шаг

## Полезные ссылки

- [Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [net/http docs](https://pkg.go.dev/net/http)
- [chi router](https://github.com/go-chi/chi)
