package service

type FileService interface {
    DownloadFile(hashUserId string,formatType string) error
}

type FileServiceImpl struct {
}


func NewFileService() FileService {
    return &FileServiceImpl{}
}

func (f *FileServiceImpl) DownloadFile(hashUserId string,formatType string) error {
    return nil
}
