package main

import (
	"flag"
	"fmt"
	"github.com/just1689/just-safe/controller"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
)

var generate = flag.String("generate", "", "Generate wallet by name")
var wallet = flag.String("wallet", "wallet", "Wallet name")
var add = flag.String("add", "", "add password")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var get = flag.String("get", "", "get password")

var driver = flag.String("driver", "local", "Storage driver to use")

func main() {
	flag.Parse()

	if *driver == "local" {
		model.StorageDriver = model.LocalDriver{}
	}

	if model.StorageDriver == nil {
		panic("no storage driver, exiting")
	}

	if *generate != "" {
		generateWallet()
		return
	}

	if *add != "" {
		addSite()
		return
	}

	if *get != "" {
		getSite()
		return
	}

}

func getSite() {
	p, err := controller.GetPasswordV1(*wallet, *get, *password, *username)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%s", p))
}

func addSite() {
	err := controller.AddPasswordV1(*wallet, *add, *username, *password)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}

}

func generateWallet() {
	err := controller.CreateWalletV1(*generate, *password)
	if err != nil {
		logrus.Errorln(err)
		panic(err)
	}
}
