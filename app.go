package main

import (
	"flag"
)

var generate = flag.String("generate", "", "Generate wallet by name")
var wallet = flag.String("wallet", "wallet", "Wallet name")
var add = flag.String("add", "", "add password")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var get = flag.String("get", "", "get password")

func main() {
	flag.Parse()

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

}

func addSite() {

}

func generateWallet() {

}
