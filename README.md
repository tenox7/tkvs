# TKVS - Trivial Key Value store in a single Json file

The primary use case is an implementation of the [Cache Interface](https://pkg.go.dev/golang.org/x/crypto/acme/autocert#Cache) for [Go acme/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert). However it can be used for anything. The motivation is backing store in a single Json file opened on start, so can be used in chroot environments.

