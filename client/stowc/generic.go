package stowc

import (
	"github.com/graymeta/stow"
	stowgs "github.com/graymeta/stow/google"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
)

type GenericDriver struct {
	client *stow.Container
}

func (g *GenericDriver) InitGoogle() {
	stowLoc, err := stow.Dial(stowgs.Kind, stow.ConfigMap{
		stowgs.ConfigJSON:      os.Getenv("GOOGLE_JSON"),
		stowgs.ConfigProjectId: os.Getenv("GOOGLE_PROJECT_ID"),
	})
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not connect to cloud storage")
		panic(err)
	}
	stowBucket, err := stowLoc.Container(os.Getenv("BUCKET"))
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not connect to cloud storage")
		panic(err)
	}
	g.client = &stowBucket

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

func (g GenericDriver) WriteFile(path string, data []byte) (err error) {
	//Implement
	return nil
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
