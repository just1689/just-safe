package api

import (
	"fmt"
	"net/http"
	"os"
)

func Listen() {

	http.HandleFunc("/api/createWallet/v1", createWalletV1)
	http.HandleFunc("/api/getPassword/v1", getPasswordV1)
	http.HandleFunc("/api/addPassword/v1", addPasswordV1)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

}
