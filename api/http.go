package api

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/controller"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

func Listen() {
	http.HandleFunc("/api/createWallet/v1", func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			logrus.Errorln("could not read body")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		body := make(map[string]string)
		if err = json.Unmarshal(b, &body); err != nil {
			logrus.Errorln("could not unmarshal body")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		password, found := body["password"]
		if !found {
			logrus.Errorln("could not find password in body")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		controller.CreateWalletV1("wallet", password)

	})

	http.HandleFunc("/api/getPassword/v1", func(writer http.ResponseWriter, request *http.Request) {

	})

	http.HandleFunc("/api/addPassword/v1", func(writer http.ResponseWriter, request *http.Request) {

	})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

}
