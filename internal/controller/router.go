package controller

import (
	"github.com/GauravMakhijani/MyPassionFundSummary/server"
	"github.com/gorilla/mux"
)

func InitRouter (deps *server.Dependencies) (router *mux.Router) {
    router = mux.NewRouter()
    router.HandleFunc("/ping", PingHandler).Methods("GET")
    router.HandleFunc("/download-mypassionfund-summary", DownloadMyPassionFundSummaryHandler(deps)).Methods("GET")
    return
}
