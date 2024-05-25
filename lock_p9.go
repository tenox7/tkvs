//go:build plan9
// +build plan9

package tkvs

import (
	"os"
	"sync"
)

type TKVS struct {
	file   *os.File
	misErr error
	kvs    KeyVal
	sync.Mutex
}

func (t *TKVS) fsLock(_ string) error { return nil }
