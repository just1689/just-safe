package api

import (
	"github.com/just1689/just-safe/controller"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

func getPasswordV1(writer http.ResponseWriter, request *http.Request) {
	item, err := model.RequestToItem(request)
	if err != nil {
		msg := "could not read body"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusBadRequest)
		return
	}

	if !item.IsValidGetPassword() {
		msg := "missing field among site, username, walletPassword"
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

	result, err := controller.GetPasswordV1(item.Site, item.WalletPassword, item.Username)
	if err != nil {
		msg := "could not get password"
		logrus.Error(err)
		http.Error(writer, msg, http.StatusInternalServerError)
		return
	}
	WriteJson(model.Item{
		Password: result,
	}, writer)

}
