# Задачи — Глава 16

> Требует Docker: `make docker-up`

## ⭐ 1. Подключение pgxpool
```go
pool, err := pgxpool.New(ctx, "postgres://restapi:restapi@localhost:5432/taskflow?sslmode=disable")
```

## ⭐ 2. Health check с ping к БД

## ⭐ 3. CREATE TABLE books (вручную или migration)

## ⭐ 4. INSERT book через Exec

## ⭐ 5. SELECT by id через QueryRow

## ⭐⭐ 6. List books с Query

## ⭐⭐ 7. UPDATE и DELETE

## ⭐⭐ 8. Context timeout на queries

## ⭐⭐ 9. pgx.CollectRows для scan

## ⭐⭐⭐ 10. Connection pool config (MaxConns, MinConns)

## ⭐⭐⭐ 11. Prepared statements

## ⭐⭐⭐ 12. Repository postgres impl для Books

## Проверка
```bash
make docker-up
psql postgres://restapi:restapi@localhost:5432/taskflow -c "SELECT 1"
```
