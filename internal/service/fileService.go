package service

import (
	"fmt"

	"log"
	"strconv"
	"time"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/literals"
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
func formatDate(date string) string {
    t,_ := time.Parse(time.RFC3339 ,date)
    return t.Format(literals.DateFormat)
}
func GenerateFakeData() (user model.FakeName, err error) {
	//fakename := []model.FakeName{}
	//generate fake name and address

	fake := faker.New()
    noOfData := 20
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
	user.Passionfund = make([]model.FakeData, noOfData)
	currTime := time.Now()
	for i := 0; i < noOfData; i++ {
		user.Passionfund[i].AccountNO = strconv.Itoa(fake.RandomNumber(10))
        user.Passionfund[i].Branch = fake.Address().City()
        user.Passionfund[i].Name = fake.Person().FirstName()
        user.Passionfund[i].CCY = "INR"
        user.Passionfund[i].StartDate = fake.Time().RFC3339(currTime)
        user.Passionfund[i].InstallmentAmount = strconv.FormatFloat(fake.RandomFloat(2,1000,1000000),'f',2,64)
        user.Passionfund[i].MaturityAmt = strconv.FormatFloat(fake.RandomFloat(2,1000,1000000),'f',2,64)
        user.Passionfund[i].DateOfMaturity = fake.Time().RFC3339(currTime)
        user.Passionfund[i].Tenure = strconv.Itoa(fake.RandomNumber(2))
        user.Passionfund[i].RateOfInterest = strconv.FormatFloat(fake.RandomFloat(2,5,20),'f',2,64)
        user.Passionfund[i].CurrentPrincipalAmt = strconv.FormatFloat(fake.RandomFloat(2,1000,1000000),'f',2,64)
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
    const (
		colCount = 12
		marginH  = 5.0
		lineHt   = 4.0
		cellGap  = 2.0
        pageMaxUsage = 180.0
        newPageStart = 10.0
	)
	// var colStrList [colCount]string
	type cellType struct {
		str  string
		list [][]byte
		ht   float64
	}
	var (
		cellList [colCount]cellType
		cell     cellType
	)
    pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	//Styling
	pdf.SetFont("Arial", "", 10)
    pdf.SetMargins(marginH, 15, marginH)

	pdf.SetTitle("MyPassionFundSummary", true)

	ximg, yimg, wimg, himg := 5.0, 5.0, 50.0, 30.0

	imageOptions := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}
	pdf.ImageOptions(".././images.png", ximg, yimg, wimg, himg, false, imageOptions, 0, "")
	pdf.SetY(35)

	pdf.SetFont("Arial", "B", 10)
	pdf.MultiCell(0, 10, "DREAM DEPOSIT SUMMARY", "", "C", false)

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

    pdf.SetFont("Arial", "B", 8)

	// Generate table header






    header := [colCount]string{"Sr.no", "Account No.", "Branch", "Name", "CCY", "Start Date", "Installment Amount", "Maturity Amount", "Date of Maturity", "Tenure (Months)","Rate of Interest", "Current Principal Amount*"}
    colWd := []float64{10,30,30,40,13,25,25,25,25,20,20,20}

    maxHt := lineHt
    y := pdf.GetY()
    for col,val:=range(header){
        cell.str =  val
        cell.list = pdf.SplitLines([]byte(cell.str), colWd[col]-cellGap-cellGap)
        cell.ht = float64(len(cell.list)) * lineHt
        if cell.ht > maxHt {
			maxHt = cell.ht
		}
        cellList[col] = cell
    }
    x:=marginH
    for col:=range(header){
        pdf.Rect(x, y, colWd[col], maxHt+cellGap+cellGap, "D")
        cell = cellList[col]
		cellY := y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[col]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[col]
    }

    pdf.Ln(6)
	pdf.SetTextColor(24, 24, 24)
	pdf.SetFillColor(255, 255, 255)
    pdf.SetFont("Arial", "", 8)

	// Rows
	y = pdf.GetY()
	srNo:= 1
    for _ ,fund := range fakename.Passionfund{
		maxHt := lineHt
		// Cell height calculation loop
        cellValues:= []string{strconv.Itoa(srNo), fund.AccountNO,fund.Branch,fund.Name,fund.CCY,
            formatDate(fund.StartDate),fund.InstallmentAmount,fund.MaturityAmt,formatDate(fund.DateOfMaturity),
            fund.Tenure,fund.RateOfInterest, fund.CurrentPrincipalAmt}

        for index,value := range(cellValues){
            cell.str = value
            cell.list = pdf.SplitLines([]byte(cell.str), colWd[index]-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
            if cell.ht > maxHt {
                maxHt = cell.ht
            }
            cellList[index] = cell
        }
		// Cell render loop
        x := marginH

        for i := range(cellValues){
            pdf.Rect(x, y, colWd[i], maxHt+cellGap+cellGap, "D")
			cell = cellList[i]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x+cellGap, cellY)
				pdf.CellFormat(colWd[i]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
					"C", false, 0, "")
				cellY += lineHt
			}
			x += colWd[i]
        }
		y += maxHt + cellGap + cellGap
        if y > pageMaxUsage {
            pdf.AddPage()
            y = newPageStart
        }
		srNo++
	}


    pdf.MultiCell(50, 10, "", "", "", false)


	currDepositBalance := "* Current Deposit Balance - Is the total installment amount paid till date towards funding of the Dream Deposit."
	line2 := "* In case of default/delay in payments of installments,the maturity value mentioned above will be different from the actual maturity value."

	pdf.MultiCell(50, 10, "", "", "", false)


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

