# Подсказки — Глава 01

> Открывай только после 30 минут самостоятельной работы!

## Задача 2 — Echo

```go
body, err := io.ReadAll(r.Body)
if err != nil {
    http.Error(w, "bad request", http.StatusBadRequest)
    return
}
defer r.Body.Close()
w.Write(body)
```

## Задача 4 — Headers

```go
headers := make(map[string]string)
for key, values := range r.Header {
    headers[key] = strings.Join(values, ", ")
}
json.NewEncoder(w).Encode(headers)
```

## Задача 7 — Query

```go
q := r.URL.Query()
page := q.Get("page")  // "" если нет
```

## Задача 8 — 404

Вариант 1 — NotFound handler:
```go
http.HandleFunc("/", handler)
// NotFoundHandler для DefaultServeMux не кастомизируется легко
// Лучше перейти на chi в главе 03
```

Вариант 2 — проверка в handler:
```go
if r.URL.Path != "/expected" {
    w.WriteHeader(http.StatusNotFound)
    // ...
}
```

## Бонус — TCP

```go
listener, _ := net.Listen("tcp", ":9000")
for {
    conn, _ := listener.Accept()
    go func(c net.Conn) {
        defer c.Close()
        fmt.Fprintf(c, "220 Welcome\r\n")
        buf := make([]byte, 1024)
        n, _ := c.Read(buf)
        fmt.Fprintf(c, "250 OK: %s\r\n", strings.TrimSpace(string(buf[:n])))
    }(conn)
}
```
