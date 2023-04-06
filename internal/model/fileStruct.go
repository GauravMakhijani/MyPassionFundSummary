package model

type Response struct {
    Message string `json:"message"`
}


type FileDownloadRequest struct{
    HashUserId string
    FormatType string
}


type PassionFund struct{
    AccountNo string
    BranchName string
    Name string
    CCY string
    StartDate string
    InstallmentAmount string
    MaturityAmount string
    DataOfMaturity string
    Tenure string
    RateOfInterest string
    CurrentPrincipalAmount string
}

type Address struct{
    City string
    State string
    Line1 string
    Line2 string
    Line3 string
    PostalCode string
    Country string
}

type UserData struct{
    Name string
    Address Address
    PassionFunds []PassionFund
}

type PassionFundReport struct{
    ReportBytes string `json:"reportBytes"`
    ReportUrl string    `json:"reportUrl"`
}
type PassionFundSummaryResponse struct{
    Body  PassionFundReport `json:"body"`
}
