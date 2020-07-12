package controller

import (
	"encoding/json"
	"github.com/just1689/just-safe/model"
	"github.com/sirupsen/logrus"
)

func CreateSession() (session model.Session, err error) {
	filename := model.GetSessionFilename()
	b, err := model.StorageDriver.ReadFile(filename)
	if err == nil {
		session := model.GetSessionFromBytes(b)
		if session != nil {
			return *session, nil
		}
	}
	s := model.Session{}
	b, err = json.Marshal(s)
	if err != nil {
		logrus.Errorln("could not marshal session")
		logrus.Errorln(err)
		return
	}
	err = model.StorageDriver.WriteFile(filename, b)
	if err != nil {
		logrus.Errorln("could not create session file")
		logrus.Errorln(err)
		return
	}
	return s, nil
}
