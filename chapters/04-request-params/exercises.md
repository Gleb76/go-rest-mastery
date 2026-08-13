# Задачи — Глава 04

## ⭐ 1. Query echo — GET /search?q=go&limit=10
## ⭐ 2. Pagination params — page, per_page с defaults (1, 20)
## ⭐ 3. Multi-value — GET /tags?tag=go&tag=api → `{"tags":["go","api"]}`
## ⭐ 4. Path param validation — /books/{id} только числа, иначе 400
## ⭐ 5. Header echo — GET /headers/Authorization
## ⭐⭐ 6. Sort params — ?sort_by=title&order=asc|desc
## ⭐⭐ 7. Filter builder — ?status=active&role=admin
## ⭐⭐ 8. Content-Type check — POST /data только application/json
## ⭐⭐⭐ 9. Query validation helper — pkg с ParsePagination(r)
## ⭐⭐⭐ 10. Books search — GET /books?q=code&author=Martin&year=2008
