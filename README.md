# TKVS - Trivial Key Value store in a single Json file

Implementation of the [Cache Interface](https://pkg.go.dev/golang.org/x/crypto/acme/autocert#Cache) for [Go acme/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert) in `chroot` environments. However it can be used for anything.

## Usage with ACME

```go
import (
	"github.com/tenox7/tkvs"
	"golang.org/x/crypto/acme/autocert"
)

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
cache.Put(ctx, "myKey", []byte("myValue"))
cache.Get(ctx, "myKey")
cache.Delete(ctx, "myKey")
```
