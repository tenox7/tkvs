# TKVS - Trivial Key Value store in a single Json file

The primary use case is an implementation of the [Cache Interface](https://pkg.go.dev/golang.org/x/crypto/acme/autocert#Cache) for [Go acme/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert). However it can be used for anything. The motivation is backing store in a single Json file, opened on start, so can be used in chroot environments.

## Usage with ACME

```go
import "github.com/tenox7/tkvs"

acm := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.mysite.com"),
		Cache:      tkvs.NewJsonCache("/var/cache/acme-store.json", autocert.ErrCacheMiss),
}

syscall.Chroot(dir)
```

## General Usage

```go
import "github.com/tenox7/tkvs"

cache := tkvs.NewJsonCache("/var/cache/my-store", errors.New("key not found"))
cache.Put(ctx, "my-key", []byte("my-value"))
cache.Get(ctx, "my-key")
cache.Delete(ctx, "my-key")
```
