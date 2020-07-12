package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/disk"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/just1689/just-safe/util/encryption/symmetric"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func CreateWalletV1(name, password string) (err error) {
	disk.CreateDir(name)

	// Create private key, public key
	private, public := asymmetric.GenerateKeys()
	publicKeyString := base64.StdEncoding.EncodeToString(public)

	//Encrypt the private key with the password
	p := []byte(password)
	p = symmetric.Pad(p)
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
	}
	b, err := json.Marshal(wallet)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not marshall wallet to json")
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/wallet.json", name), b, 0644)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could write wallet json")
		return
	}
	return
}