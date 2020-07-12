package model

var WalletFilename = "wallet.json"

type Wallet struct {
	PrivateKeyEncrypted string `json:"privateKeyEncrypted"`
	PublicKeyPlain      string `json:"publicKeyPlain"`
	Salt                string `json:"salt"`
}
