package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/just1689/just-safe/model"
	"github.com/just1689/just-safe/util/disk"
	"github.com/just1689/just-safe/util/encryption/asymmetric"
	"github.com/just1689/just-safe/util/encryption/symmetric"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var generate = flag.String("generate", "", "Generate wallet by name")
var wallet = flag.String("wallet", "wallet", "Wallet name")
var add = flag.String("add", "", "add password")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var get = flag.String("get", "", "get password")

func main() {
	flag.Parse()

	if *generate != "" {
		generateWallet()
		return
	}

	if *add != "" {
		addSite()
		return
	}

	if *get != "" {
		getSite()
		return
	}

}

func getSite() {
	logrus.Println("Site:", *get)
	logrus.Println("Username:", *username)
	logrus.Println("Wallet password:", *password)

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/wallet.json", *wallet))
	if err != nil {
		logrus.Errorln(err)
		panic("could not load wallet json")
		return
	}
	w := &model.Wallet{}
	err = json.Unmarshal(b, w)
	if err != nil {
		logrus.Errorln(err)
		panic("could not unmarshal wallet")
		return
	}

	privateKeyEncryptedBytes, err := base64.StdEncoding.DecodeString(w.PrivateKeyEncrypted)
	if err != nil {
		logrus.Errorln(err)
		panic("could not turn privateKey 64 to bytes")
		return
	}

	p := []byte(*password)
	p = symmetric.Pad(p)
	privateKeyBytes, err := symmetric.Decrypt(p, privateKeyEncryptedBytes)
	if err != nil {
		logrus.Errorln(err)
		panic("could not descrypt private key with the password")
		return
	}

	in := fmt.Sprintf("%s/%s.site.json", *wallet, *get)
	siteBytes, err := ioutil.ReadFile(in)
	if err != nil {
		logrus.Errorln(err)
		panic("could not read the site json")
		return
	}
	s := &Site{}
	err = json.Unmarshal(siteBytes, s)
	if err != nil {
		logrus.Errorln(err)
		panic("could not unmarshal site json to bytes")
		return
	}

	//TODO: find by username
	// ...................
	// ...................
	// ...................

	encryptedPasswordBytes, err := base64.StdEncoding.DecodeString(s.Entries[0].Password)
	if err != nil {
		logrus.Errorln(err)
		panic("could not decode encrypted password 64")
		return
	}

	decrypted, err := asymmetric.Decrypt(encryptedPasswordBytes, privateKeyBytes)
	if err != nil {
		logrus.Errorln(err)
		panic("could not decrypt password")
		return
	}

	logrus.Println("Decrypted password", string(decrypted))

}

func addSite() {
	logrus.Println("Site:", *add)
	logrus.Println("Username:", *username)
	logrus.Println("Password:", *password)

	//Load the wallet
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/wallet.json", *wallet))
	if err != nil {
		logrus.Errorln(err)
		panic(fmt.Sprintln("could not read wallet", *wallet))
		return
	}
	w := &model.Wallet{}
	err = json.Unmarshal(b, w)
	if err != nil {
		logrus.Errorln(err)
		panic("could not unmarshal wallet")
		return
	}
	var publicKey []byte = make([]byte, 2048)
	l, err := base64.StdEncoding.Decode(publicKey, []byte(w.PublicKeyPlain))
	publicKey = append([]byte(nil), publicKey[:l]...)
	if err != nil {
		logrus.Errorln(err)
		panic("could not get public key from base64")
		return
	}

	encryptedBytes, err := asymmetric.Encrypt([]byte(*password), publicKey)
	if err != nil {
		logrus.Errorln(err)
		panic("could not encrypt the password with wallet public key")
	}

	passwordStringEncrypted := base64.StdEncoding.EncodeToString(encryptedBytes)
	site := Site{
		Site:    *add,
		Entries: make([]Entry, 1),
	}
	site.Entries[0] = Entry{
		Username: *username,
		Password: passwordStringEncrypted,
	}
	b, err = json.Marshal(site)
	out := fmt.Sprintf("%s/%s.site.json", *wallet, *add)
	fmt.Println(out)
	ioutil.WriteFile(out, b, 0644)
}

type Site struct {
	Site    string  `json:"site"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func generateWallet() {

	disk.CreateDir(*generate)

	// Create private key, public key
	private, public := asymmetric.GenerateKeys()
	publicKeyString := base64.StdEncoding.EncodeToString(public)

	//Encrypt the private key with the password
	p := []byte(*password)
	p = symmetric.Pad(p)
	privateKeyEncrypted, err := symmetric.Encrypt(p, private)
	if err != nil {
		logrus.Errorln(err)
		panic("could not encrypt private wallet key with password")
	}

	privateKeyEncryptedString := base64.StdEncoding.EncodeToString(privateKeyEncrypted)

	wallet := model.Wallet{
		PrivateKeyEncrypted: privateKeyEncryptedString,
		PublicKeyPlain:      publicKeyString,
	}
	b, err := json.Marshal(wallet)
	if err != nil {
		logrus.Errorln(err)
		panic("could not marshall wallet to json")
		return
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/wallet.json", *generate), b, 0644)
	if err != nil {
		logrus.Errorln(err)
		panic("could write wallet json")
		return
	}
}
