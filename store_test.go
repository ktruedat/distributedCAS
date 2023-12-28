package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "sometestpic"
	pathname := CASPathTransformFunc(key)
	expectedPathname := ""
	fmt.Println(pathname)
	if pathname != expectedPathname {
		t.Errorf("have %s want %s", pathname, expectedPathname)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{PathTransformFunc: CASPathTransformFunc}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("somepic", data); err != nil {
		t.Error(err)
	}
}
