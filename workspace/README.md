# Workspace — твои решения

Здесь ты пишешь **свой код** по мере прохождения глав.

## Как организовать

```bash
# Для каждой главы создавай отдельную папку:
mkdir -p workspace/ch07
cd workspace/ch07
go mod init github.com/YOUR_USERNAME/ch07-crud-books
```

## Рекомендуемая структура (с главы 11)

```
workspace/ch11/
├── cmd/server/main.go
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
├── go.mod
└── README.md          # что реализовал, как запускать
```

## Git

Коммить после каждой главы:

```bash
git add workspace/ch07/
git commit -m "ch07: CRUD API для books с in-memory storage"
```

## Не копируй starter/

Папки `chapters/*/starter/` — для экспериментов и подсказок.  
Твоя реализация живёт здесь, в `workspace/`.
