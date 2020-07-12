package stowc

import (
	"github.com/graymeta/stow"
	stowgs "github.com/graymeta/stow/google"
	"io/ioutil"
	"log"
	"os"
)

type GoogleDriver struct {
	client *stow.Container
}

func (g *GoogleDriver) Init() {
	stowLoc, err := stow.Dial(stowgs.Kind, stow.ConfigMap{
		stowgs.ConfigJSON:      os.Getenv("GOOGLE_JSON"),
		stowgs.ConfigProjectId: os.Getenv("GOOGLE_PROJECT_ID"),
	})
	if err != nil {
		log.Fatal(err)
	}

	stowBucket, err := stowLoc.Container(os.Getenv("BUCKET"))
	if err != nil {
		log.Fatal(err)
	}

	g.client = &stowBucket

}

func (g GoogleDriver) ReadFile(path string) (b []byte, err error) {
	err = stow.Walk(*g.client, stow.NoPrefix, 100, func(item stow.Item, err error) error {
		if err != nil {
			return err
		}
		log.Println(item.Name())

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
	if err != nil {
		return
	}
	return
}

func (g GoogleDriver) WriteFile(path string, data []byte) (err error) {
	//Implement
	return nil
}

func (g GoogleDriver) ListFiles() (out chan string, err error) {
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
