package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "sometestpic"
	pathKey := CASPathTransformFunc(key)
	expectedPathname := ""
	expectedPathnameOriginal := ""
	fmt.Println(pathKey)
	if pathKey.PathName != expectedPathname {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathname)
	}
	if pathKey.Original != expectedPathnameOriginal {
		t.Errorf("have %s want %s", pathKey.Original, expectedPathnameOriginal)
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
