package stowc

import (
	"bytes"
	"github.com/graymeta/stow"
	"io/ioutil"
	"log"
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
	err = stow.Walk(*g.client, stow.NoPrefix, 100, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		if item.Name() == path {
			r, err := item.Open()
			if err != nil {
				return err
			}
			defer r.Close()
			b, err = ioutil.ReadAll(r)
		}
		return nil
	})
	return
}

func (g GenericDriver) WriteFile(name string, data []byte) (err error) {
	r := bytes.NewReader(data)
	size := int64(len(data))
	_, err = (*g.client).Put(name, r, size, nil)
	return
}

func (g GenericDriver) ListFiles() (out chan string, err error) {
	out = make(chan string)
	err = stow.Walk(*g.client, stow.NoPrefix, 100, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		log.Println(item.Name())
		out <- item.Name()
		return nil
	})
	close(out)
	if err != nil {
		return
	}
	return

}
