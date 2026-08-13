# Задачи — Глава 05

## ⭐ 1. POST /books — decode JSON, return created book with id
## ⭐ 2. GET /books/:id — encode struct to JSON
## ⭐ 3. omitempty — optional description field
## ⭐ 4. Invalid JSON → 400 with error message
## ⭐ 5. Empty body POST → 400
## ⭐⭐ 6. Request/Response DTOs — не expose internal fields
## ⭐⭐ 7. time.Time formatting — RFC3339 in JSON
## ⭐⭐ 8. Nullable fields — `*string` for optional author
## ⭐⭐ 9. json.Decoder DisallowUnknownFields()
## ⭐⭐⭐ 10. List response wrapper — `{"data": [...], "count": N}`
## ⭐⭐⭐ 11. Streaming — json.Encoder for large list (1000 books)
## ⭐⭐⭐ 12. Custom UnmarshalJSON for enum status field
