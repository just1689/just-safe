package main

import (
	"flag"
	"github.com/just1689/just-safe/api"
	"github.com/just1689/just-safe/client/stowc"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	flag.Parse()
	logrus.Infoln("Starting up...")
	logrus.Infoln("")
	logrus.Infoln("")
	logrus.Infoln("")
	driver := os.Getenv("driver")
	if driver == "" {
		driver = "google" //TODO: replace with local
	}

	if driver != "" {
		d := stowc.GenericDriver{}
		d.Init(driver)
		stowc.StorageDriver = d
	}
	if stowc.StorageDriver == nil {
		panic("no storage driver, exiting")
	}

	if os.Getenv("PORT") != "" {
		logrus.Infoln("> started")
		api.Listen()
	}

}
