package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type LocalDriver struct {
}

func (d LocalDriver) CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModeDir)
	}
}

func (d LocalDriver) ReadFile(path string) (b []byte, err error) {
	return ioutil.ReadFile(path)
}

func (d LocalDriver) WriteFile(path string, data []byte) (err error) {
	return ioutil.WriteFile(path, data, 0644)
}

func (d LocalDriver) ListFiles(p string) (out chan string, err error) {
	if _, err = os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		return
	}
	out = make(chan string)
	go func() {
		defer close(out)
		err = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			out <- path
			return nil
		})
	}()
	return out, nil

}
