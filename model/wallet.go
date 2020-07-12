package model

type Wallet struct {
	PrivateKeyEncrypted string `json:"privateKeyEncrypted"`
	PublicKeyPlain      string `json:"publicKeyPlain"`
	Salt                string `json:"salt"`
}
