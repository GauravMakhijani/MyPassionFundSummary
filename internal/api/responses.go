package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Response(w http.ResponseWriter,status int, response interface{}){
    respBytes, err := json.Marshal(response)
	if err != nil {
		logrus.WithField("err", err.Error()).Error()
		status = http.StatusInternalServerError
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)
}
