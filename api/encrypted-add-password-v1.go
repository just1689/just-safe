package api

import (
	"encoding/json"
	"github.com/just1689/just-safe/controller"
	"github.com/sirupsen/logrus"
	"net/http"
)

func encryptedAddPasswordV1(writer http.ResponseWriter, request *http.Request) {
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

	site, foundSite := body["site"]
	username, foundUsername := body["username"]
	password, foundPassword := body["password"]
	walletPassword, foundWalletPassword := body["walletPassword"]
	if !foundPassword || !foundUsername || !foundSite || !foundWalletPassword {
		logrus.Errorln("could not find field of [username, password, site, walletPassword] in body")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	walletPasswordOK := controller.CheckPassword(walletPassword)
	if !walletPasswordOK {
		logrus.Errorln("bad password")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = controller.AddPasswordV1(site, username, password)
	if err != nil {
		logrus.Errorln(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
