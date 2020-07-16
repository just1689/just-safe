package model

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type EncryptedBody struct {
	Session string `json:"session"`
	Payload string `json:"payload"`
}

func RequestToEncryptedBody(r *http.Request) (encryptedBody EncryptedBody, err error) {
	defer r.Body.Close()
	var b []byte
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not read body of http request")
		return
	}
	encryptedBody, err = BytesToEncryptedItem(b)
	return
}

func BytesToEncryptedItem(b []byte) (encryptedBody EncryptedBody, err error) {
	encryptedBody = EncryptedBody{}
	if err = json.Unmarshal(b, &encryptedBody); err != nil {
		logrus.Errorln(err)
	}
	return
}
