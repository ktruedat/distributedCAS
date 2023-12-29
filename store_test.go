package main

import (
	"bytes"
	"fmt"
	"io"
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
	if pathKey.FileName != expectedPathnameOriginal {
		t.Errorf("have %s want %s", pathKey.FileName, expectedPathnameOriginal)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{PathTransformFunc: CASPathTransformFunc}
	s := NewStore(opts)

	key := "specialpic"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}
	b, _ := io.ReadAll(r)
	fmt.Println(string(b))
	if string(b) != string(data) {
		t.Errorf("Want %v have %v", data, b)
	}

}
