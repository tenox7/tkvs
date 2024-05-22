package tkvs_test

import (
	"errors"
	"log"
	"testing"

	"github.com/tenox7/tkvs"
)

func TestAll(t *testing.T) {
	myErr := errors.New("key not found")
	tk := tkvs.New("test.json", myErr)

	val, err := tk.Get(nil, "foo")
	if err != nil && err != myErr {
		log.Fatal(err)
	}
	log.Printf("val=%q", string(val))

	err = tk.Put(nil, "foo", []byte("bar"))
	if err != nil {
		log.Fatal(err)
	}

	val, err = tk.Get(nil, "foo")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("val=%q", string(val))

	err = tk.Put(nil, "baz", []byte("bug"))
	if err != nil {
		log.Fatal(err)
	}

	keys := tk.Keys()
	log.Printf("keys=%v", keys)

	err = tk.Delete(nil, "fuq")
	if err != nil {
		log.Fatal(err)
	}

	keys = tk.Keys()
	log.Printf("keys=%v", keys)
}
