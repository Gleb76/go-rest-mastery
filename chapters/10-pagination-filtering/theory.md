# Теория: Pagination

## Offset pagination

```
GET /books?page=2&per_page=20&sort_by=title&sort_order=asc&author=Martin
```

```go
type Pagination struct {
    Page       int `json:"page"`
    PerPage    int `json:"per_page"`
    Total      int `json:"total"`
    TotalPages int `json:"total_pages"`
}

type ListResponse[T any] struct {
    Data       []T        `json:"data"`
    Pagination Pagination `json:"pagination"`
}
```

## Cursor pagination (better for large datasets)

```
GET /books?cursor=eyJpZCI6MTAwfQ&limit=20
```

Response includes `next_cursor`.
