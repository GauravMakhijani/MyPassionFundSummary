package service

import (
	"fmt"
	"log"
	"time"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/jaswdr/faker"
	"github.com/jung-kurt/gofpdf"
)

type FileService interface {
	DownloadFile() (model.FileDownloadResponse, error)
}

type FileServiceImpl struct {
}

func NewFileService() FileService {
	return &FileServiceImpl{}
}
func GenerateFakeData() (user model.FakeName, err error) {
	//fakename := []model.FakeName{}
	//generate fake name and address

	fake := faker.New()
	// fakenames := model.FakeName{
	user.Name = fake.Person().Name()
	user.Add = model.Address{
		Line1:   fake.Address().BuildingNumber(),
		Line2:   fake.Address().SecondaryAddress(),
		Line3:   fake.Address().StreetName(),
		City:    fake.Address().City(),
		State:   fake.Address().State(),
		Country: fake.Address().Country(),
		Pincode: fake.Address().PostCode(),
	}
	user.Passionfund = make([]model.FakeData, 5)
	for i := 0; i < 5; i++ {
		user.Passionfund[i].AccountNO = fake.RandomNumber(14)
		user.Passionfund[i].Branch = fake.App().Name()
		user.Passionfund[i].Name = fake.Person().FirstName()
		user.Passionfund[i].CCY = fake.Currency().Country()
		user.Passionfund[i].StartDate = fake.Time().RFC1123(time.Time{})
		user.Passionfund[i].InstallmentAmount = fake.Currency().Number()
		user.Passionfund[i].MaturityAmt = fake.Currency().Number()
		user.Passionfund[i].DateOfMaturity = fake.Time().RFC3339(time.Time{})
		user.Passionfund[i].Tenure = fake.RandomNumber(5)
		user.Passionfund[i].RateOfInterest = fake.RandomNumber(4)
		user.Passionfund[i].CurrentPrincipalAmt = fake.Currency().Number()

	}

	return

}

func (f *FileServiceImpl) DownloadFile() (response model.FileDownloadResponse, err error) {
	fakeData, err := GenerateFakeData()
	if err != nil {
		return
	}

	err = GeneratePDF(fakeData)
	if err != nil {
		log.Fatal(err)
	}
	return model.FileDownloadResponse{}, nil
}

func GeneratePDF(fakename model.FakeName) error {
	pdf := gofpdf.New("P", "mm", "A3", "")
	pdf.AddPage()

	//Styling
	pdf.SetFont("Arial", "", 10)
	pdf.SetMargins(1, 0, 1)
	pdf.SetTitle("MyPassionFundSummary", true)

	x, y, w, h := 5.0, 5.0, 50.0, 30.0

	imageOptions := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}

	pdf.ImageOptions(".././images.png", x, y, w, h, false, imageOptions, 0, "")
	pdf.SetY(35)
	//pdf.CellFormat(0, 10, "DREAM DEPOSIT SUMMARY")
	pdf.SetFont("Arial", "B", 10)
	pdf.MultiCell(0, 10, "DREAM DEPOSIT SUMMARY", "", "C", false)
	//pdf.SetY(y + h + 5)
	pdf.SetFont("Arial", "B", 10)
	name := "Name"
	pdf.MultiCell(70, 5, name+fakename.Name, "", "L", false)

	pdf.MultiCell(50, 5, fmt.Sprintf("Address:%s\t", fakename.Add.Line1+","), "", "L", false)

	pdf.MultiCell(50, 4, fakename.Add.Line2+",", "", "C", false)
	pdf.MultiCell(50, 4, fakename.Add.Line3+",", "", "C", false)
	pdf.MultiCell(50, 4, fakename.Add.City+",", "", "C", false)
	pdf.MultiCell(50, 4, fakename.Add.State+",", "", "C", false)
	pdf.MultiCell(50, 4, fakename.Add.Country+",", "", "C", false)
	pdf.MultiCell(50, 4, fakename.Add.Pincode+",", "", "C", false)
	//to add space between address and table
	pdf.MultiCell(50, 10, "", "", "", false)
	//fmt.Sprintf("Name:%s\t", fakename.Name)
	//pdf.SetY(y + h + 10)
	pdf.SetFont("Arial", "", 7)

	// Generate table header
	pdf.CellFormat(25, 6, "Account No", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 6, "Branch", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 6, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 6, "CCY", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 6, "Start date", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 6, "Installment Amount", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 6, "Maturity Amount", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 6, "Date of Maturity ", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 6, "Tenure", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 6, "Rate of Interest", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 6, "Current Principal amt", "1", 1, "C", false, 0, "")

	//for _, fakeName := range fakename {

	for _, data := range fakename.Passionfund {
		fmt.Print("is data gets print")
		pdf.CellFormat(25, 6, fmt.Sprintf("%d", (data.AccountNO)), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 6, data.Branch, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, data.Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, data.CCY, "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 6, data.StartDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 6, fmt.Sprintf("%d", (data.InstallmentAmount)), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 6, fmt.Sprintf("%d", (data.MaturityAmt)), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 6, data.DateOfMaturity, "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 6, fmt.Sprintf("%d", (data.Tenure)), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, fmt.Sprintf("%d", (data.RateOfInterest)), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 6, fmt.Sprintf("%d", (data.CurrentPrincipalAmt)), "1", 1, "C", false, 0, "")

	}
	//}
	currDepositBalance := "1)Current Deposit Balance-Is the total installment amount paid till date towards funding ofthe Dream Deposit."
	line2 := "2)In case of default/delay in payments of installments,the maturity value mentioned above will be different from the actual maturity value."
	//pdf.SetY(-1)
	pdf.MultiCell(50, 10, "", "", "", false)

	//pdf.SetY(y + h + 10)

	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 5, currDepositBalance, "", "", false)
	pdf.MultiCell(0, 5, line2, "", "", false)

	err := pdf.OutputFileAndClose("FDSummary.pdf")
	if err != nil {
		fmt.Println("ok here 4")

		return err
	}

	return nil

}
