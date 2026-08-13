# Теория: JWT Authentication

## Flow

```
1. POST /auth/register → create user (bcrypt password)
2. POST /auth/login → verify password → return JWT
3. Client sends: Authorization: Bearer <token>
4. Middleware validates JWT → sets user in context
5. Handler reads user from context
```

## bcrypt

```go
import "golang.org/x/crypto/bcrypt"

hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
err := bcrypt.CompareHashAndPassword(hash, []byte(password))
```

## JWT

```go
import "github.com/golang-jwt/jwt/v5"

claims := jwt.MapClaims{
    "sub": userID,
    "exp": time.Now().Add(24 * time.Hour).Unix(),
    "iat": time.Now().Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
signed, _ := token.SignedString([]byte(secret))
```

## Auth Middleware

```go
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            header := r.Header.Get("Authorization")
            if !strings.HasPrefix(header, "Bearer ") {
                respondError(w, 401, "UNAUTHORIZED", "missing token")
                return
            }
            tokenStr := strings.TrimPrefix(header, "Bearer ")
            token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
                return []byte(secret), nil
            })
            if err != nil || !token.Valid {
                respondError(w, 401, "UNAUTHORIZED", "invalid token")
                return
            }
            claims := token.Claims.(jwt.MapClaims)
            userID := claims["sub"].(string)
            ctx := context.WithValue(r.Context(), userIDKey, userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

## Security

- Password: `json:"-"` — never serialize
- Secret in env var, not in code
- HTTPS in production
- Short access token TTL (15min–24h)
