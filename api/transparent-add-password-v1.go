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
		msg := "could not read body"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	if !item.IsValidAddPassword() {
		msg := "missing field among site, username, password, walletPassword"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	walletPasswordOK := controller.CheckPassword(item.WalletPassword)
	if !walletPasswordOK {
		msg := "password was incorrect"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	err = controller.AddPasswordV1(item.Site, item.Username, item.Password)
	if err != nil {
		msg := "could not add password"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
