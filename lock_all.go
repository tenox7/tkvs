//go:build !plan9
// +build !plan9

package tkvs

import (
	"os"
	"sync"

	"github.com/danjacques/gofslock/fslock"
)

type TKVS struct {
	file   *os.File
	misErr error
	lock   *fslock.Handle
	kvs    KeyVal
	sync.Mutex
}

func (t *TKVS) fsLock(path string) error {
	l, err := fslock.Lock(path)
	if err != nil {
		return err
	}
	t.lock = &l
	return nil
}
