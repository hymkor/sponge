package main

import (
	"os"
)

type LazyFile struct {
	Name string
	Perm os.FileMode
	Flag int

	fd *os.File
}

func (L *LazyFile) Write(buf []byte) (int, error) {
	if L.fd == nil {
		var err error
		L.fd, err = os.OpenFile(L.Name, L.Flag, L.Perm)
		if err != nil {
			return 0, err
		}
	}
	return L.fd.Write(buf)
}

func (f *LazyFile) Close() error {
	if f.fd == nil {
		return nil
	}
	err := f.fd.Close()
	f.fd = nil
	return err
}
