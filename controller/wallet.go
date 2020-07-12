package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/client/stowc"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/just1689/just-safe/util/encryption/symmetric"
	"github.com/sirupsen/logrus"
)

func GetWalletV1() (wallet model.Wallet, err error) {
	b, err := stowc.StorageDriver.ReadFile(model.WalletFilename)
	if err != nil {
		logrus.Errorln("could not open wallet json file")
		logrus.Errorln(err)
		return
	}
	w := &model.Wallet{}
	err = json.Unmarshal(b, w)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not unmarshal wallet")
		return
	}
	wallet = *w
	return
}

func CreateWalletV1(password string) (err error) {

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
	err = stowc.StorageDriver.WriteFile(fmt.Sprintf(model.WalletFilename), b)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could write wallet json")
		return
	}
	return
}
