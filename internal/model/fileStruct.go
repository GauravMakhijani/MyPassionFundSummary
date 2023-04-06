package model

type Response struct {
	Message string `json:"message"`
}

type FileDownloadRequest struct {
	HashUserId string `json:"hashUserId"`
	FormatType string `json:"format_type"`
}

type FileDownloadResponse struct {
	ReportBytes     string `json:"reportBytes"`
	ReportGenerated string `json:"reportGenerated"`
}

type Address struct {
	Line1   string
	Line2   string
	Line3   string
	City    string
	State   string
	Country string
	Pincode string
}

type FakeName struct {
	Name        string
	Add         Address
	Passionfund []FakeData
}

type FakeData struct {
	AccountNO           int
	Branch              string
	Name                string
	CCY                 string
	StartDate           string
	InstallmentAmount   int
	MaturityAmt         int
	DateOfMaturity      string
	Tenure              int
	RateOfInterest      int
	CurrentPrincipalAmt int
}
