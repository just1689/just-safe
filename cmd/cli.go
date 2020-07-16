package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/io"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
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
		item, err := model.MapToItem(input)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		url := endpoint + "api/addPassword/v1"
		b, err := json.Marshal(item)
		if err != nil {
			//handle
			logrus.Errorln(err)
			return
		}

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
		item, err := model.MapToItem(input)
		if err != nil {
			logrus.Errorln(err)
			return
		}

		url := endpoint + "api/getPassword/v1"
		b, err := json.Marshal(item)
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

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			logrus.Errorln("Failed.")
			//TODO: get error from server
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
