package asymmetric

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"github.com/sirupsen/logrus"
)

var Bits = 2048

func GenerateKeys() (private, public []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, Bits)
	if err != nil {
		panic(err)
	}
	return x509.MarshalPKCS1PrivateKey(privateKey), x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

}

func Encrypt(data, publicKeyBytes []byte) (encryptedBytes []byte, err error) {
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	encryptedBytes, err = rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		data,
		nil)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return

}

func Decrypt(encryptedBytes, privateKeyBytes []byte) (decryptedBytes []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	decryptedBytes, err = privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return

}
