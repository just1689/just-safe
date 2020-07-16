package model

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Item struct {
	Site           string `json:"site,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	WalletPassword string `json:"walletPassword,omitempty"`
}

func MapToItem(m map[string]string) (item Item, err error) {
	b, err := json.Marshal(m)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	item = Item{}
	err = json.Unmarshal(b, &item)
	return
}
