# Теория: API Versioning

## URL prefix (recommended)

```go
r.Route("/api/v1", v1Routes)
r.Route("/api/v2", v2Routes)
```

## Breaking changes

- Remove field
- Change field type
- Change URL structure
- Change status code semantics

## Non-breaking

- Add optional field
- Add new endpoint
- Add new query param

## Deprecation

```go
w.Header().Set("Deprecation", "true")
w.Header().Set("Sunset", "Sat, 01 Jan 2028 00:00:00 GMT")
w.Header().Set("Link", "</api/v2/books>; rel="successor-version"")
```
