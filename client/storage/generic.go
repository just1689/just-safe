package storage

import (
	"bytes"
	"fmt"
	"github.com/graymeta/stow"
	"github.com/pkg/errors"
	"io/ioutil"
)

var DRIVER_GOOGLE = "google"
var DRIVER_LOCAL = "local"

type GenericDriver struct {
	client *stow.Container
}

func (g *GenericDriver) Init(driver string) {
	if driver == DRIVER_GOOGLE {
		g.initGoogle()
		return
	} else if driver == DRIVER_LOCAL {
		g.initLocal()
		return
	}
}

func (g GenericDriver) ReadFile(path string) (b []byte, err error) {
	found := false
	err = stow.Walk(*g.client, stow.NoPrefix, 100, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		if item.Name() == path {
			found = true
			r, err := item.Open()
			if err != nil {
				return err
			}
			defer r.Close()
			b, err = ioutil.ReadAll(r)
		}
		return nil
	})
	if !found {
		err = errors.New(fmt.Sprint("could not find file named", path))
	}
	return
}

func (g GenericDriver) WriteFile(name string, data []byte) (err error) {
	r := bytes.NewReader(data)
	size := int64(len(data))
	_, err = (*g.client).Put(name, r, size, nil)
	return
}

func (g GenericDriver) ListFiles() (out chan string, err error) {
	out = make(chan string, 256)
	defer close(out)
	err = stow.Walk(*g.client, stow.NoPrefix, 100, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		out <- item.Name()
		return nil
	})
	if err != nil {
		return
	}
	return

}

func (g GenericDriver) DeleteFile(f string) (err error) {
	return (*g.client).RemoveItem(f)
}
