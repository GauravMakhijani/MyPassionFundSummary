package server

import "github.com/GauravMakhijani/MyPassionFundSummary/internal/service"

type Dependencies struct {
    FileService service.FileService
}

func InitDependencies() (*Dependencies, error) {
    return &Dependencies{
        FileService: service.NewFileService(),
    }, nil
}
