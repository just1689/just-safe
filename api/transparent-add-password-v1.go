package api

import (
	"github.com/just1689/just-safe/controller"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

func addPasswordV1(writer http.ResponseWriter, request *http.Request) {
	item, err := model.RequestToItem(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "could not read body to item", http.StatusBadRequest)
		return
	}

	if !item.IsValidAddPassword() {
		logrus.Errorln("could not find field of [username, password, site, walletPassword] in body")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	walletPasswordOK := controller.CheckPassword(item.WalletPassword)
	if !walletPasswordOK {
		logrus.Errorln("bad password")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = controller.AddPasswordV1(item.Site, item.Username, item.Password)
	if err != nil {
		logrus.Errorln(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
