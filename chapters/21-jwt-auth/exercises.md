# Задачи — Глава 21

## ⭐ 1. User model с password hash (bcrypt)

## ⭐ 2. POST /auth/register — email, password, name

## ⭐ 3. POST /auth/login — return JWT access token

## ⭐ 4. JWT claims: sub (user_id), exp, iat

## ⭐ 5. Auth middleware — parse Bearer token

## ⭐⭐ 6. GET /users/me — protected route

## ⭐⭐ 7. Invalid token → 401

## ⭐⭐ 8. Missing token → 401

## ⭐⭐ 9. Password never in JSON response

## ⭐⭐⭐ 10. Refresh token (optional)

## ⭐⭐⭐ 11. Token expiry check

## ⭐⭐⭐ 12. Context user — store user in request context

## ⭐⭐⭐ 13. Logout / token blacklist (Redis)

## ⭐⭐⭐ 14. Integration test: register → login → me

## Проверка
```bash
curl -X POST localhost:8080/api/v1/auth/register -d '{"email":"a@b.com","password":"secret123","name":"Test"}'
curl -X POST localhost:8080/api/v1/auth/login -d '{"email":"a@b.com","password":"secret123"}'
curl -H "Authorization: Bearer TOKEN" localhost:8080/api/v1/users/me
```
