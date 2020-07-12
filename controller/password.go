package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/just1689/just-safe/util/encryption/symmetric"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func AddPasswordV1(wallet, site, username, password string) (err error) {
	logrus.Println("Site:", site)
	logrus.Println("Username:", username)
	logrus.Println("Password:", password)

	//Load the wallet
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/wallet.json", wallet))
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln(fmt.Sprintln("could not read wallet", wallet))
		return
	}
	w := &model.Wallet{}
	err = json.Unmarshal(b, w)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not unmarshal wallet")
		return
	}
	var publicKey []byte = make([]byte, 2048)
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
	s := model.Site{
		Site:    site,
		Entries: make([]model.Entry, 1),
	}
	s.Entries[0] = model.Entry{
		Username: username,
		Password: passwordStringEncrypted,
	}
	b, err = json.Marshal(s)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could marshal site to json")
		return

	}
	out := fmt.Sprintf("%s/%s.site.json", wallet, site)
	fmt.Println(out)
	ioutil.WriteFile(out, b, 0644)
	return
}

func GetPasswordV1(wallet, site, walletPassword, username string) (sitePassword string, err error) {
	logrus.Println("Site:", site)
	logrus.Println("Username:", username)
	logrus.Println("Wallet password:", walletPassword)

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/wallet.json", wallet))
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
	p = symmetric.Pad(p)
	privateKeyBytes, err := symmetric.Decrypt(p, privateKeyEncryptedBytes)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not descrypt private key with the password")
		return
	}

	in := fmt.Sprintf("%s/%s.site.json", wallet, site)
	siteBytes, err := ioutil.ReadFile(in)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not read the site json")
		return
	}
	s := &model.Site{}
	err = json.Unmarshal(siteBytes, s)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not unmarshal site json to bytes")
		return
	}

	//TODO: find by username
	// ...................
	// ...................
	// ...................

	encryptedPasswordBytes, err := base64.StdEncoding.DecodeString(s.Entries[0].Password)
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

	logrus.Println("Decrypted password", string(decrypted))
	sitePassword = string(decrypted)
	return

}
