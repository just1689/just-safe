package api

import (
	"github.com/just1689/just-safe/client/storage"
	"github.com/just1689/just-safe/controller"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

func createWalletV1(writer http.ResponseWriter, request *http.Request) {
	item, err := model.RequestToItem(request)
	if err != nil {
		msg := "could not read body"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	if item.IsAddWallet() {
		msg := "no walletPassword field provided"
		logrus.Errorln(msg)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	_, err = storage.StorageDriver.ReadFile(model.WalletFilename)
	if err == nil {
		msg := "wallet already exists"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	controller.CreateWalletV1(item.WalletPassword)

}
