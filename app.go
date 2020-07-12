package main

import (
	"flag"
	"github.com/just1689/just-safe/api"
	"github.com/just1689/just-safe/client/storage"
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
		d := storage.GenericDriver{}
		d.Init(driver)
		storage.StorageDriver = d
	}
	if storage.StorageDriver == nil {
		panic("no storage driver, exiting")
	}

	if os.Getenv("PORT") != "" {
		logrus.Infoln("> started")
		api.Listen()
	}

}
