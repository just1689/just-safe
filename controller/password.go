package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/just1689/just-safe/util/encryption/symmetric"
	"github.com/sirupsen/logrus"
)

func AddPasswordV1(site, username, password string) (err error) {

	//Load the wallet
	b, err := model.StorageDriver.ReadFile(fmt.Sprintf("wallet.json"))
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not read wallet")
		return
	}
	w := &model.Wallet{}
	err = json.Unmarshal(b, w)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not unmarshal wallet")
		return
	}
	var publicKey = make([]byte, 2048)
	l, err := base64.StdEncoding.Decode(publicKey, []byte(w.PublicKeyPlain))
	publicKey = append([]byte(nil), publicKey[:l]...)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not get public key from base64")
		return
	}

	encryptedBytes, err := asymmetric.Encrypt([]byte(password), publicKey)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not encrypt the password with wallet public key")
	}

	passwordStringEncrypted := base64.StdEncoding.EncodeToString(encryptedBytes)

	s, _ := readSite(site)
	if s == nil {
		s = &model.Site{
			Site:    site,
			Entries: make([]model.Entry, 0),
		}
	}
	s.AddItem(username, passwordStringEncrypted)
	b, err = json.Marshal(s)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could marshal site to json")
		return

	}
	filename := model.GetSiteFilename(site)
	model.StorageDriver.WriteFile(filename, b)
	return
}

func GetPasswordV1(site, walletPassword, username string) (sitePassword string, err error) {
	b, err := model.StorageDriver.ReadFile(fmt.Sprintf("wallet.json"))
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not load wallet json")
		return
	}
	w := &model.Wallet{}
	err = json.Unmarshal(b, w)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not unmarshal wallet")
		return
	}

	privateKeyEncryptedBytes, err := base64.StdEncoding.DecodeString(w.PrivateKeyEncrypted)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not turn privateKey 64 to bytes")
		return
	}

	p := []byte(walletPassword)
	p = symmetric.Pad(p, w.Salt)
	privateKeyBytes, err := symmetric.Decrypt(p, privateKeyEncryptedBytes)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not decrypt private key with the password")
		return
	}

	s, err := readSite(site)
	if err != nil {
		logrus.Errorln("could not read site from storage")
		logrus.Errorln(err)
		return
	}

	var correctEntry model.Entry
	found := false

	for _, entry := range s.Entries {
		if entry.Username == username {
			correctEntry = entry
			found = true
			break
		}
	}

	if !found {
		err = errors.New("could not find username in entries")
		logrus.Errorln(err)
		logrus.Errorln(err.Error())
		return
	}

	encryptedPasswordBytes, err := base64.StdEncoding.DecodeString(correctEntry.Password)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not decode encrypted password 64")
		return
	}

	decrypted, err := asymmetric.Decrypt(encryptedPasswordBytes, privateKeyBytes)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not decrypt password")
		return
	}

	sitePassword = string(decrypted)
	return

}

func CheckPassword(walletPassword string) bool {
	w, err := GetWalletV1()
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not wallet")
		return false
	}
	privateKeyEncryptedBytes, err := base64.StdEncoding.DecodeString(w.PrivateKeyEncrypted)
	if err != nil {
		logrus.Errorln("could not get password")
		//TODO: handle errors in a more correct way
		return false
	}

	p := []byte(walletPassword)
	p = symmetric.Pad(p, w.Salt)
	privateKeyBytes, err := symmetric.Decrypt(p, privateKeyEncryptedBytes)
	if err != nil {
		return false
	}
	return len(privateKeyBytes) > 0
}

func readSite(site string) (s *model.Site, err error) {
	filename := model.GetSiteFilename(site)
	siteBytes, err := model.StorageDriver.ReadFile(filename)
	if err != nil {
		return
	}
	s = &model.Site{}
	err = json.Unmarshal(siteBytes, s)
	if err != nil {
		return
	}
	return
}
