package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/sirupsen/logrus"
	"time"
)

func DecryptBody(b []byte) (payload []byte, err error) {
	e := &model.EncryptedBody{}
	err = json.Unmarshal(b, e)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	n := time.Now()

	filename := fmt.Sprintf("session.%v-%v-%v.json", n.Year(), n.Month(), n.Day())
	sessionBytes, err := model.StorageDriver.ReadFile(filename)
	if err != nil {
		logrus.Errorln("could not unmarshal read session")
		logrus.Errorln(err)
		return
	}

	s := &model.Session{}
	err = json.Unmarshal(sessionBytes, s)
	if err != nil {
		ok = false
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
