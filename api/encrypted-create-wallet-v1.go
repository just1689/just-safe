package api

import (
	"encoding/json"
	"github.com/just1689/just-safe/controller"
	"github.com/sirupsen/logrus"
	"net/http"
)

func encryptedCreateWalletV1(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	stop, b, err := ReadBody(writer, request)
	if stop {
		return
	}
	payload, err := DecryptBody(b)
	if err != nil {
		logrus.Errorln("could not decrypt body")
		logrus.Errorln(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	body := make(map[string]string)
	if err = json.Unmarshal(payload, &body); err != nil {
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

}
