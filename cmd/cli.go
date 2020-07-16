package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"syscall"
)

func main() {
	endpoint := "http://localhost:8080/"

	if os.Args[1] == "add" {
		site, username, password, walletPassword := credentials()
		url := endpoint + "api/addPassword/v1"
		body := map[string]string{
			"site":           site,
			"username":       username,
			"walletPassword": walletPassword,
			"password":       password,
		}
		b, err := json.Marshal(body)
		if err != nil {
			//handle
			logrus.Errorln(err)
			return
		}
		//TODO: loading bar
		http.Post(url, "application/json", bytes.NewReader(b))

	} else if os.Args[1] == "get" {
		site, username, walletPassword := getter()
		url := endpoint + "api/getPassword/v1"
		body := map[string]string{
			"site":           site,
			"username":       username,
			"walletPassword": walletPassword,
		}
		b, err := json.Marshal(body)
		if err != nil {
			//handle
			logrus.Errorln(err)
			return
		}
		//TODO: loading bar
		resp, err := http.Post(url, "application/json", bytes.NewReader(b))
		if err != nil {
			logrus.Error(err)
			return
		}
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}

		r := make(map[string]string)
		err = json.Unmarshal(b, &r)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println("")
		fmt.Println("")

		if err := clipboard.WriteAll(r["password"]); err != nil {
			panic(err)
		}
		fmt.Println("COPIED TO CLIPBOARD!")

	}

}

func credentials() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	var bytePassword []byte
	var err error

	fmt.Print("Enter site: ")
	site, _ := reader.ReadString('\n')

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	if runtime.GOOS == "windows" {
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on unix
	} else {
		bytePassword, err = terminal.ReadPassword(0)
	}
	if err != nil {
		//Handle error
		return "", "", "", ""
	}
	password := string(bytePassword)
	fmt.Println("")

	fmt.Print("Enter Wallet password: ")
	if runtime.GOOS == "windows" {
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on unix
	} else {
		bytePassword, err = terminal.ReadPassword(0)
	}
	if err != nil {
		//Handle error
		return "", "", "", ""
	}
	walletPassword := string(bytePassword)

	return strings.TrimSpace(site), strings.TrimSpace(username), strings.TrimSpace(password), strings.TrimSpace(walletPassword)
}

func getter() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	var bytePassword []byte
	var err error

	fmt.Print("Enter site: ")
	site, _ := reader.ReadString('\n')

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Wallet password: ")
	if runtime.GOOS == "windows" {
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on unix
	} else {
		bytePassword, err = terminal.ReadPassword(0)
	}
	if err != nil {
		//Handle error
		return "", "", ""
	}
	walletPassword := string(bytePassword)

	return strings.TrimSpace(site), strings.TrimSpace(username), strings.TrimSpace(walletPassword)
}
