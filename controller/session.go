package controller

import (
	"encoding/base64"
	"encoding/json"
	"github.com/just1689/just-safe/client/storage"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/sirupsen/logrus"
)

func CreateSession() (session model.Session, err error) {
	filename := model.GetSessionFilename()
	b, err := storage.StorageDriver.ReadFile(filename)
	if err == nil {
		session := model.GetSessionFromBytes(b)
		if session != nil {
			return *session, nil
		}
	}

	go GCSessions()

	private, public := generatePrivatePublicKeyPair64()
	s := model.Session{
		PrivateKey: private,
		PublicKey:  public,
	}
	b, err = json.Marshal(s)
	if err != nil {
		logrus.Errorln("could not marshal session")
		logrus.Errorln(err)
		return
	}
	err = storage.StorageDriver.WriteFile(filename, b)
	if err != nil {
		logrus.Errorln("could not create session file")
		logrus.Errorln(err)
		return
	}
	return s, nil
}

func generatePrivatePublicKeyPair64() (private, public string) {
	priv, publ := asymmetric.GenerateKeys()
	private = base64.StdEncoding.EncodeToString(priv)
	public = base64.StdEncoding.EncodeToString(publ)
	return

}

func GCSessions() {
	logrus.Infoln("> Starting GC")
	out, err := storage.StorageDriver.ListFiles()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	for filename := range out {
		if model.IsSessionFilename(filename) {
			if model.GetSessionFilename() != filename {
				logrus.Infoln(">> Deleting old session", filename)
				err := storage.StorageDriver.DeleteFile(filename)
				if err != nil {
					logrus.Errorln("   FAIL")
				} else {
					logrus.Infoln("   OK")
				}
			}
		}
	}
}
