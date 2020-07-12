package stowc

import (
	"github.com/graymeta/stow"
	stowgs "github.com/graymeta/stow/google"
	"github.com/sirupsen/logrus"
	"os"
)

func (g *GenericDriver) initGoogle() {
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
