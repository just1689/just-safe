package api

import (
	"encoding/json"
	"github.com/just1689/just-safe/controller"
	"github.com/sirupsen/logrus"
	"net/http"
)

func createSessionV1(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	session, err := controller.CreateSession()
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not create session")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	session.PrivateKey = ""
	b, err := json.Marshal(session)
	if err != nil {
		logrus.Errorln("could not marshal session")
		logrus.Errorln(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(b)
}
