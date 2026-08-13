# Задачи — Глава 24

## ⭐ 1. Unit test для service layer (mock repository)

## ⭐ 2. Table-driven tests

## ⭐ 3. httptest.NewRecorder для handler

## ⭐ 4. Test GET /health → 200

## ⭐ 5. Test POST invalid JSON → 400

## ⭐⭐ 6. Mock BookRepository с testify/mock или manual mock

## ⭐⭐ 7. Test CRUD flow: create → get → update → delete

## ⭐⭐ 8. Test 404 not found

## ⭐⭐ 9. Test auth middleware — no token → 401

## ⭐⭐⭐ 10. Integration test с testcontainers или docker postgres

## ⭐⭐⭐ 11. Test helper: newTestServer(t)

## ⭐⭐⭐ 12. Coverage > 70% для handlers

## ⭐⭐⭐ 13. Benchmark handler ListBooks

## ⭐⭐⭐ 14. CI script: go test -race -cover ./...

## Пример httptest

```go
func TestHealth(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    rec := httptest.NewRecorder()
    handler.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)
}
```
