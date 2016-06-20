package statestore

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadWrite(t *testing.T) {
	dir, err := ioutil.TempDir("", "statestore-test")
	if err != nil {
		t.Fatalf("fail to create temporary directory: %s", err)
	}
	defer os.RemoveAll(dir)
	f, err := ioutil.TempFile(dir, "test")
	if err != nil {
		t.Fatalf("fail to create temp file: %s", err)
	}
	defer os.Remove(f.Name())
	st := NewFileStateStore(f.Name())
	type SomeType struct {
		A, B, C, D string
	}
	some := SomeType{"a", "b", "c", "d"}
	if err := st.Write(some); err != nil {
		t.Fatalf("write file failed: %s", err)
	}
	var receive SomeType
	if err := st.Read(&receive); err != nil {
		t.Fatalf("read file failed: %s", err)
	}
	if receive != some {
		t.Fatalf("recover from file, but not equals. given: %v, received: %v", some, receive)
	}
}

func TestReadNotFound(t *testing.T) {
	dir, err := ioutil.TempDir("", "statestore-readonly")
	if err != nil {
		t.Fatalf("fail to create temporary directory: %s", err)
	}
	defer os.RemoveAll(dir)
	f, err := ioutil.TempFile(dir, "test")
	if err != nil {
		t.Fatalf("fail to create temp file: %s", err)
	}
	defer os.Remove(f.Name())
	st := NewFileStateStore(f.Name())
	type SomeType struct {
		A, B, C, D string
	}
	var receive SomeType
	if err := st.Read(&receive); err != nil {
		t.Fatalf("read file failed: %s", err)
	}
}
