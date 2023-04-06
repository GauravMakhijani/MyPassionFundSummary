package controller

import (
	"encoding/json"
	"net/http"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/api"
	"github.com/GauravMakhijani/MyPassionFundSummary/internal/literals"
	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/GauravMakhijani/MyPassionFundSummary/server"
	"github.com/sirupsen/logrus"
)

func PingHandler(rw http.ResponseWriter, rq *http.Request)  {
    response := model.Response{
        Message : "Pong",
    }
    respBytes,err := json.Marshal(response)
    if err != nil {
        logrus.WithField("err", err.Error()).Error("Error marshalling ping response")
        rw.WriteHeader(http.StatusInternalServerError)
    }
    rw.Write(respBytes)
}


func handleDownloadPDF(w http.ResponseWriter, r *http.Request , deps *server.Dependencies,downloadRequest *model.FileDownloadRequest){
    //to be implemented by Rutuja
}

func handleDownloadExcel(w http.ResponseWriter, r *http.Request , deps *server.Dependencies,downloadRequest *model.FileDownloadRequest){
    ///to be implemented by Gaurav
    passionFundSummary, err:= deps.FileService.DownloadFileAsExcel(downloadRequest)
    if err != nil {
        response := model.Response{
            Message : literals.ErrServerError,
        }
        api.Response(w,http.StatusInternalServerError,response)
        return
    }
    api.Response(w,http.StatusCreated,passionFundSummary)
}


func DownloadMyPassionFundSummaryHandler(deps *server.Dependencies) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var downloadRequest model.FileDownloadRequest
        err:= json.NewDecoder(r.Body).Decode(&downloadRequest)
        if err != nil {
            response := model.Response{
                Message : literals.ErrInvalidInput,
            }
            api.Response(w,http.StatusBadRequest,response)
            return
        }

        if(downloadRequest.HashUserId == "" || downloadRequest.FormatType == ""){
            response := model.Response{
                Message : literals.ErrInvalidInput,
            }
            api.Response(w,http.StatusBadRequest,response)
            return
        }

        switch downloadRequest.FormatType{
        case literals.PDFFormat:
            handleDownloadPDF(w,r,deps,&downloadRequest)
        case literals.ExcelFormat:
            handleDownloadExcel(w,r,deps,&downloadRequest)
        default:
            response := model.Response{
                Message : literals.ErrInvalidInput,
            }
            api.Response(w,http.StatusBadRequest,response)
        }
    })
}
