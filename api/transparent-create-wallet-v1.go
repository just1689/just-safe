package api

import (
	"encoding/json"
	"github.com/just1689/just-safe/client/storage"
	"github.com/just1689/just-safe/controller"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

func createWalletV1(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	stop, b, err := ReadBody(writer, request)
	if stop {
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

	_, err = storage.StorageDriver.ReadFile(model.WalletFilename)
	if err == nil {
		logrus.Errorln("wallet already exists")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	controller.CreateWalletV1(password)

}
