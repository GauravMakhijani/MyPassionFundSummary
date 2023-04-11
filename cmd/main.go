package main

import (
	"os"
	"path/filepath"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/controller"
	"github.com/GauravMakhijani/MyPassionFundSummary/server"
	"github.com/robfig/cron/v3"
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
	c := cron.New()
	c.AddFunc("@every 0h20m", RemoveExpiredFiles)
	c.Start()
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":8080")

}

func RemoveExpiredFiles() {

	dir := "../internal/storage"
	logrus.Info("files deletion initiated")
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return
		}
	}
	logrus.Info("files deleted successfully")
}
