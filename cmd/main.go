package main

import (
	"github.com/GauravMakhijani/MyPassionFundSummary/internal/controller"
	"github.com/GauravMakhijani/MyPassionFundSummary/server"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	deps, err := server.InitDependencies()
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize dependencies")
		return
	}

	router := controller.InitRouter(deps)
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":33001")

}
