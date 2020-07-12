package model

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Session struct {
	PrivateKey string `json:"privateKey,omitempty"`
	PublicKey  string `json:"publicKey,omitempty"`
}

func GetSessionFilename() string {
	d, m, y := time.Now().Date()
	return fmt.Sprintf("session.%v-%v-%v.json", y, int(m), d)
}

func IsSessionFilename(f string) bool {
	return strings.HasPrefix(f, "session.")
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
