package api

import (
	"encoding/json"
	"github.com/just1689/just-safe/controller"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func addPasswordV1(writer http.ResponseWriter, request *http.Request) {
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

	site, foundSite := body["site"]
	username, foundUsername := body["username"]
	password, foundPassword := body["password"]
	if !foundPassword || !foundUsername || !foundSite {
		logrus.Errorln("could not find field of [username, password, site] in body")
		writer.WriteHeader(http.StatusBadRequest)
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
