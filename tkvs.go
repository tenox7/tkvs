// Trivial Key Value Store with Json Backend
package tkvs

import (
	"context"
	"encoding/json"
	"io"
	"log"
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

type KeyVal map[string][]byte

type Container struct {
	KeyVal `json:"keyval"`
}

func (j *TKVS) readJson() error {
	j.Lock()
	defer j.Unlock()
	j.file.Seek(0, 0)
	buf, err := io.ReadAll(j.file)
	if err != nil {
		return err
	}
	if len(buf) == 0 {
		j.kvs = make(KeyVal)
		return nil
	}
	c := Container{}
	if err = json.Unmarshal(buf, &c); err != nil {
		return err
	}
	j.kvs = c.KeyVal
	return nil
}

func (j *TKVS) writeJson() error {
	j.Lock()
	defer j.Unlock()
	c := Container{KeyVal: j.kvs}
	out, err := json.Marshal(&c)
	if err != nil {
		return err
	}
	j.file.Seek(0, 0)
	j.file.Truncate(0)
	_, err = j.file.Write(out)
	return err
}

func (j *TKVS) Get(_ context.Context, key string) ([]byte, error) {
	val, ok := j.kvs[key]
	if !ok {
		return nil, j.misErr
	}
	return val, nil
}

func (j *TKVS) Put(_ context.Context, key string, data []byte) error {
	j.kvs[key] = data
	return j.writeJson()
}

func (j *TKVS) Delete(_ context.Context, key string) error {
	delete(j.kvs, key)
	return j.writeJson()
}

func (j *TKVS) Keys() ([]string, error) {
	s := []string{}
	for n := range j.kvs {
		s = append(s, n)
	}
	return s, nil
}

func New(path string, misErr error) *TKVS {
	l, err := fslock.Lock(path)
	if err != nil {
		log.Fatalf("unable to lock %q: %v", path, err)
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("unable to open %q: %v", path, err)
	}
	k := &TKVS{file: f, misErr: misErr, lock: &l}
	err = k.readJson()
	if err != nil {
		log.Fatalf("unable to read %q: %v", path, err)
	}
	return k
}
