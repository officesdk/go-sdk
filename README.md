# How to User

## register

r.Engine is `github.com/gin-gonic/gin`

callback is `github.com/officesdk/go-sdk/callback`

```go
callback.RegisterRouter(r.Engine)
callback.RegisterHandler(&icallback.Handler{})
```

you should implement `icallback.Handler`

> var _ callback.Handler = (*Handler)(nil)
