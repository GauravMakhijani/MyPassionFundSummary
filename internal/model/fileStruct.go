package model

type Response struct {
    Message string `json:"message"`
}


type FileDownloadRequest struct{
    HashUserId string
    FormatType string
}

