package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/GauravMakhijani/MyPassionFundSummary/server"
	"github.com/sirupsen/logrus"
)

func PingHandler(rw http.ResponseWriter, rq *http.Request) {
	response := model.Response{
		Message: "Pong",
	}
	respBytes, err := json.Marshal(response)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("Error marshalling ping response")
		rw.WriteHeader(http.StatusInternalServerError)
	}
	rw.Write(respBytes)
}

func handleDownloadPDF(w http.ResponseWriter, r *http.Request, deps *server.Dependencies, downloadRequest *model.FileDownloadRequest) {
	//to be implemented by Rutuja

	err := deps.service.DownloadFile(&downloadRequest)
    if err!=nil{
        
    }
}

func handleDownloadExcel(w http.ResponseWriter, r *http.Request, deps *server.Dependencies, downloadRequest *model.FileDownloadRequest) {
	///to be implemented by Gaurav
}

func DownloadMyPassionFundSummaryHandler(deps *server.Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var downloadRequest model.FileDownloadRequest
		err := json.NewDecoder(r.Body).Decode(&downloadRequest)
		if err != nil {
			response := model.Response{
				Message: "err- invalid input",
			}
			respBytes, err := json.Marshal(response)
			if err != nil {
				logrus.WithField("err", err.Error()).Error("Error marshalling ping response")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(respBytes)
		}

		switch downloadRequest.FormatType {
		case "P":
			handleDownloadPDF(w, r, deps, &downloadRequest)
			// err:=deps.service.DownloadFile(r.Context(),&downloadRequest)
			// if err
		case "E":
			handleDownloadExcel(w, r, deps, &downloadRequest)
		default:
			response := model.Response{
				Message: "err- invalid format",
			}
			respBytes, err := json.Marshal(response)
			if err != nil {
				logrus.WithField("err", err.Error()).Error("Error marshalling ping response")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(respBytes)
		}
	})
}
