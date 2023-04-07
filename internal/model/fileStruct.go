package model

type Response struct {
	Message string `json:"message"`
}

type FileDownloadRequest struct {
	HashUserId string `json:"hashUserId"`
	FormatType string `json:"formatType"`
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
	AccountNO           string
	Branch              string
	Name                string
	CCY                 string
	StartDate           string
	InstallmentAmount   string
	MaturityAmt         string
	DateOfMaturity      string
	Tenure              string
	RateOfInterest      string
	CurrentPrincipalAmt string
}

type cellType struct {
	str  string
	list [][]byte
	ht   float64
}
