# Теория: PostgreSQL + pgx

## Docker

```bash
make docker-up
# postgres://restapi:restapi@localhost:5432/taskflow
```

## pgxpool

```go
import "github.com/jackc/pgx/v5/pgxpool"

pool, err := pgxpool.New(ctx, databaseURL)
defer pool.Close()

if err := pool.Ping(ctx); err != nil {
    log.Fatal(err)
}
```

## QueryRow — одна строка

```go
var book Book
err := pool.QueryRow(ctx,
    "SELECT id, title, author FROM books WHERE id = $1", id,
).Scan(&book.ID, &book.Title, &book.Author)

if errors.Is(err, pgx.ErrNoRows) {
    return nil, domain.ErrNotFound
}
```

## Query — много строк

```go
rows, err := pool.Query(ctx, "SELECT id, title, author FROM books ORDER BY id")
defer rows.Close()

books, err := pgx.CollectRows(rows, pgx.RowToStructByName[Book])
```

## Exec — INSERT/UPDATE/DELETE

```go
_, err := pool.Exec(ctx,
    "INSERT INTO books (title, author) VALUES ($1, $2)", title, author,
)
```

## Context

Всегда передавай `r.Context()` из handler — для cancellation и timeouts.

```go
ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
defer cancel()
```

## SQL injection

✅ Всегда placeholders `$1, $2`  
❌ Никогда string concatenation
