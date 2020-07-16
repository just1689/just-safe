package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/just1689/just-safe/client/storage"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ReadBody(writer http.ResponseWriter, request *http.Request) (stop bool, b []byte, err error) {
	defer request.Body.Close()
	b, err = ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"msg": "Bad request body"}`))
		stop = true
		return
	}
	return
}

func DecryptBody(b []byte) (payload []byte, err error) {
	e := &model.EncryptedBody{}
	err = json.Unmarshal(b, e)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	filename := model.GetSessionFilename()
	sessionBytes, err := storage.StorageDriver.ReadFile(filename)
	if err != nil {
		logrus.Errorln("could not unmarshal read session")
		logrus.Errorln(err)
		return
	}

	s := &model.Session{}
	err = json.Unmarshal(sessionBytes, s)
	if err != nil {
		logrus.Errorln("could not unmarshal bytes")
		logrus.Errorln(err)
		return
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(s.PrivateKey)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not base64 decode private key bytes")
		return
	}

	payloadBytes, err := base64.StdEncoding.DecodeString(e.Payload)
	if err != nil {
		logrus.Errorln("could not decode base64 payload")
		logrus.Errorln(err)
		return
	}

	payload, err = asymmetric.Decrypt(payloadBytes, privateKeyBytes)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not decrypt payload")
		return
	}
	return

}

func WriteJson(item interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(item)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not marshal response")
		http.Error(w, "could not marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(b)

}
