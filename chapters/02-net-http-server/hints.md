# Подсказки — Глава 02

## Uptime
```go
var startTime = time.Now()
uptime := time.Since(startTime).Seconds()
```

## Graceful shutdown
```go
go func() {
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()
<-quit
server.Shutdown(ctx)
```

## Method not allowed
```go
if r.Method != http.MethodGet {
    w.Header().Set("Allow", "GET")
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
}
```
