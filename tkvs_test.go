package tkvs_test

import (
	"errors"
	"log"
	"testing"

	"github.com/tenox7/tkvs"
)

func TestAll(t *testing.T) {
	tk := tkvs.New("test.json", errors.New("key not found"))
	err := tk.Put(nil, "foo", []byte("bar"))
	if err != nil {
		log.Fatal(err)
	}

	val, err := tk.Get(nil, "foo")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("val=", string(val))

	keys, err := tk.Keys()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("keys=", keys)

	err = tk.Delete(nil, "foo")
	if err != nil {
		log.Fatal(err)
	}
}
