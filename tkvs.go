// Trivial Key Value Store with Json Backend
package tkvs

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type KVS struct {
	file   *os.File
	misErr error
	sync.Mutex
}

type KeyVal map[string][]byte

type Container struct {
	KeyVal `json:"keyval"`
}

func (j *KVS) readJson() (KeyVal, error) {
	j.Lock()
	defer j.Unlock()
	j.file.Seek(0, 0)
	buf, err := io.ReadAll(j.file)
	if err != nil {
		return nil, err
	}
	if len(buf) == 0 {
		log.Print("making new container")
		return make(KeyVal), nil
	}
	c := Container{}
	if err = json.Unmarshal(buf, &c); err != nil {
		return nil, err
	}
	return c.KeyVal, nil
}

func (j *KVS) writeJson(kv *KeyVal) error {
	j.Lock()
	defer j.Unlock()
	c := Container{KeyVal: *kv}
	out, err := json.Marshal(c)
	if err != nil {
		return err
	}
	j.file.Seek(0, 0)
	j.file.Truncate(0)
	_, err = j.file.Write(out)
	return err
}

func (j *KVS) Get(_ context.Context, key string) ([]byte, error) {
	log.Printf("get %q", key)
	kv, err := j.readJson()
	if err != nil {
		return nil, err
	}
	val, ok := kv[key]
	if !ok {
		return nil, j.misErr
	}
	return val, nil
}

func (j *KVS) Put(_ context.Context, key string, data []byte) error {
	log.Printf("put %q", key)
	kv, err := j.readJson()
	if err != nil {
		return err
	}
	kv[key] = data
	return j.writeJson(&kv)
}

func (j *KVS) Delete(_ context.Context, key string) error {
	log.Printf("del %q", key)
	kv, err := j.readJson()
	if err != nil {
		return err
	}
	delete(kv, key)
	return j.writeJson(&kv)
}

func New(path string, misErr error) *KVS {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	return &KVS{file: f, misErr: misErr}
}
