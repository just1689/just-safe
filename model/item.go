package model

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Site           string `json:"site,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	WalletPassword string `json:"walletPassword,omitempty"`
}

func (i Item) IsAddWallet() bool {
	if i.WalletPassword == "" {
		return false
	}
	return true
}

func (i Item) IsValidAddPassword() bool {
	if i.Site == "" {
		return false
	}
	if i.Username == "" {
		return false
	}
	if i.Password == "" {
		return false
	}
	if i.WalletPassword == "" {
		return false
	}
	return true
}
func (i Item) IsValidGetPassword() bool {
	if i.Site == "" {
		return false
	}
	if i.Username == "" {
		return false
	}
	if i.WalletPassword == "" {
		return false
	}
	return true
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

func RequestToItem(r *http.Request) (item Item, err error) {
	defer r.Body.Close()
	var b []byte
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		logrus.Errorln("could not read body of http request")
		return
	}
	item = Item{}
	if err = json.Unmarshal(b, &item); err != nil {
		logrus.Errorln(err)
	}
	return

}
