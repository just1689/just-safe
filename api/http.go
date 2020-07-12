package api

import (
	"fmt"
	"net/http"
	"os"
)

func Listen() {

	//Session
	http.HandleFunc("/api/createSession/v1", createSessionV1)

	//API
	http.HandleFunc("/api/createWallet/v1", createWalletV1)
	http.HandleFunc("/api/getPassword/v1", getPasswordV1)
	http.HandleFunc("/api/addPassword/v1", addPasswordV1)

	//Second-layer encrypted
	http.HandleFunc("/api/encrypted/createWallet/v1", encryptedCreateWalletV1)
	http.HandleFunc("/api/encrypted/getPassword/v1", encryptedGetPasswordV1)
	http.HandleFunc("/api/encrypted/addPassword/v1", encryptedAddPasswordV1)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

}
