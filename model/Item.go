package model

type Item struct {
	Site           string `json:"site"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	WalletPassword string `json:"walletPassword"`
}
