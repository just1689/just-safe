package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/just1689/just-safe/util/encryption/symmetric"
	"github.com/sirupsen/logrus"
)

func CreateWalletV1(name, password string) (err error) {

	// Create private key, public key
	private, public := asymmetric.GenerateKeys()
	publicKeyString := base64.StdEncoding.EncodeToString(public)

	//Encrypt the private key with the password
	p := []byte(password)
	salt := util.RandStringRunes(32)
	p = symmetric.Pad(p, salt)
	privateKeyEncrypted, err := symmetric.Encrypt(p, private)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not encrypt private wallet key with password")
		return
	}

	privateKeyEncryptedString := base64.StdEncoding.EncodeToString(privateKeyEncrypted)

	wallet := model.Wallet{
		PrivateKeyEncrypted: privateKeyEncryptedString,
		PublicKeyPlain:      publicKeyString,
		Salt:                salt,
	}
	b, err := json.Marshal(wallet)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not marshall wallet to json")
		return
	}
	err = model.StorageDriver.WriteFile(fmt.Sprintf("wallet.json"), b)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could write wallet json")
		return
	}
	return
}
