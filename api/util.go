package api

import (
	"io/ioutil"
	"net/http"
)

func ReadBody(writer http.ResponseWriter, request *http.Request) (stop bool, b []byte, err error) {
	defer request.Body.Close()
	b, err = ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		stop = true
		return
	}
	return
}
