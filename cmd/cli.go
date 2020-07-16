package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/io"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {
	endpoint := "http://localhost:8080/"

	if os.Args[1] == "add" {
		input, err := io.PromptForInput([]io.NextInput{
			io.PromptSite,
			io.PromptUsername,
			io.PromptPassword,
			io.PromptWalletPassword,
		})
		if err != nil {
			logrus.Errorln(err)
			return
		}
		site := input["site"]
		username := input["username"]
		password := input["password"]
		walletPassword := input["walletPassword"]

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

		io.LoadingBar(2000)

		http.Post(url, "application/json", bytes.NewReader(b))

	} else if os.Args[1] == "get" {
		input, err := io.PromptForInput([]io.NextInput{
			io.PromptSite,
			io.PromptUsername,
			io.PromptWalletPassword,
		})
		if err != nil {
			logrus.Errorln(err)
			return
		}
		site := input["site"]
		username := input["username"]
		walletPassword := input["walletPassword"]

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
		go io.LoadingBar(2000)

		resp, err := http.Post(url, "application/json", bytes.NewReader(b))
		if err != nil {
			logrus.Error(err)
			return
		}
		//fmt.Println(resp.StatusCode)
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

		//fmt.Println("")
		//fmt.Println("")

		pass := r["password"]
		if os.Args[2] == "clipboard" || os.Args[2] == "c" {
			io.WriteToClipboard(pass)
		}

		if os.Args[2] == "keyboard" || os.Args[2] == "k" {
			if err != nil {
				panic(err)
			}
			keyboard, err := io.GenerateKeyboardScript(pass)
			if err != nil {
				logrus.Errorln(err)
				return
			}

			for i := 1; i <= 3; i++ {
				fmt.Println("Writing in", i)
				time.Sleep(500 * time.Millisecond)
			}
			keyboard()
		}

	}

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
