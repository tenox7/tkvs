# TKVS - Trivial Key Value store in a single Json file

Implementation of the [Cache Interface](https://pkg.go.dev/golang.org/x/crypto/acme/autocert#Cache) for [Go acme/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert) in `chroot` environments. However it can be used for anything. The key/value store is realized in a single Json file opened on startup, thus available even if caller invokes `chroot` after.

## Usage with ACME

```go
import (
	"github.com/tenox7/tkvs"
	"golang.org/x/crypto/acme/autocert"
)

acm := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.mysite.com"),
		Cache:      tkvs.New("/var/cache/acme-store.json", autocert.ErrCacheMiss),
}

syscall.Chroot(dir)
```

## General Usage

```go
import "github.com/tenox7/tkvs"

kvs := tkvs.New("/var/cache/mystore.json", errors.New("key not found"))
kvs.Put(ctx, "myKey", []byte("myValue"))
kvs.Get(ctx, "myKey")
kvs.Delete(ctx, "myKey")
```
