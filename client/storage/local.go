package storage

import (
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/local"
	"github.com/sirupsen/logrus"
	"os"
)

func (g *GenericDriver) initLocal() {

	stowLoc, err := stow.Dial(local.Kind, stow.ConfigMap{
		local.ConfigKeyPath: os.Getenv("WALLET_PATH"),
	})
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not connect to local storage")
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
