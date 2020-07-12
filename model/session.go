package model

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Session struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func GetSessionFilename() string {
	n := time.Now()
	return fmt.Sprintf("session.%v-%v-%v.json", n.Year(), n.Month(), n.Day())
}

func GetSessionFromBytes(b []byte) *Session {
	s := &Session{}
	err := json.Unmarshal(b, s)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not unmarshal session from bytes")
		logrus.Errorln(string(b))
	}
	return s
}
