package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/api"
	"github.com/GauravMakhijani/MyPassionFundSummary/internal/literals"
	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/GauravMakhijani/MyPassionFundSummary/server"

	//s"github.com/GauravMakhijani/MyPassionFundSummary/internal/5"

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

func handleDownloadPDF(w http.ResponseWriter, r *http.Request, deps *server.Dependencies, downloadRequest model.FileDownloadRequest) {
	//to be implemented by Rutuja
	//fmt.Print("hiiiiiiiii")
	response, err := deps.FileService.DownloadFile(downloadRequest)
	if err != nil {
		response := model.Response{
			Message: "fail to download file",
		}
		api.Response(w, http.StatusForbidden, response)
		return

	}
	json_response, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(json_response)
}

func handleDownloadExcel(w http.ResponseWriter, r *http.Request, deps *server.Dependencies, downloadRequest *model.FileDownloadRequest) {
	///to be implemented by Gaurav
}

func DownloadMyPassionFundSummaryHandler(deps *server.Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var downloadRequest model.FileDownloadRequest
		err := json.NewDecoder(r.Body).Decode(&downloadRequest)
		fmt.Print("hello rutuja")
		if err != nil {
			response := model.Response{
				Message: literals.ErrInvalidInput,
			}
			api.Response(w, http.StatusBadRequest, response)
		}

		fmt.Println("hiii>>>>>>>>>>>>")

		switch downloadRequest.FormatType {
		case literals.PDFFormat:
			handleDownloadPDF(w, r, deps, downloadRequest)
		case literals.ExcelFormat:
			handleDownloadExcel(w, r, deps, &downloadRequest)
		default:
			response := model.Response{
				Message: literals.ErrInvalidFormat,
			}
			api.Response(w, http.StatusBadRequest, response)
		}
	})
}
