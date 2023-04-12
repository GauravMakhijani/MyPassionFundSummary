package server

import service "github.com/GauravMakhijani/MyPassionFundSummary/internal/Service"

type Dependencies struct {
    FileService service.FileService
}

func InitDependencies() (*Dependencies, error) {
    return &Dependencies{
        FileService: service.NewFileService(),
    }, nil
}
