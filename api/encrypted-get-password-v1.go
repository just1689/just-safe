package api

import (
	"encoding/json"
	"github.com/just1689/just-safe/controller"
	"github.com/sirupsen/logrus"
	"net/http"
)

func encryptedGetPasswordV1(writer http.ResponseWriter, request *http.Request) {
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

	if !foundPassword || !foundUsername || !foundSite {
		logrus.Errorln("could not find field of [username, password, site] in body")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := controller.GetPasswordV1(site, password, username)
	if err != nil {
		logrus.Errorln(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	body = map[string]string{
		"password": result,
	}
	b, err = json.Marshal(body)
	if err != nil {
		logrus.Errorln(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(b)
}
