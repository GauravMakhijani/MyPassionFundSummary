package service

import (
	"errors"

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
	t, _ := time.Parse(time.RFC3339, date)
	return t.Format(literals.DateFormat)
}
func GenerateFakeData() (user model.FakeName, err error) {

	fake := faker.New()
	noOfData := 20

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
		user.Passionfund[i].InstallmentAmount = strconv.FormatFloat(fake.RandomFloat(2, 1000, 1000000), 'f', 2, 64)
		user.Passionfund[i].MaturityAmt = strconv.FormatFloat(fake.RandomFloat(2, 1000, 1000000), 'f', 2, 64)
		user.Passionfund[i].DateOfMaturity = fake.Time().RFC3339(currTime)
		user.Passionfund[i].Tenure = strconv.Itoa(fake.RandomNumber(2))
		user.Passionfund[i].RateOfInterest = strconv.FormatFloat(fake.RandomFloat(2, 5, 20), 'f', 2, 64)
		user.Passionfund[i].CurrentPrincipalAmt = strconv.FormatFloat(fake.RandomFloat(2, 1000, 1000000), 'f', 2, 64)
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

func GeneratePDF(fakename model.FakeName) (err error) {
	const (
		colCount     = 12
		marginH      = 5.0
		lineHt       = 4.0
		cellGap      = 2.0
		pageMaxUsage = 180.0
		newPageStart = 10.0
	)

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
	pdf.SetMargins(marginH, 10, marginH)

	pdf.SetTitle("MyPassionFundSummary", true)

	ximg, yimg, wimg, himg := 5.0, 5.0, 50.0, 30.0

	imageOptions := gofpdf.ImageOptions{
		ImageType:             "png",
		ReadDpi:               true,
		AllowNegativePosition: true,
	}
	pdf.ImageOptions(".././images.png", ximg, yimg, wimg, himg, false, imageOptions, 0, "")
	pdf.SetY(30)

	pdf.SetFont("Arial", "B", 10)
	pdf.MultiCell(100, 10, "DREAM DEPOSIT SUMMARY", "", "C", false)

	pdf.SetFont("Arial", "B", 9)
	name := "Name:"
	pdf.CellFormat(18, 5, name, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(65.3, 5, fakename.Name, "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 9)

	address := "Address:"
	pdf.CellFormat(18, 5, address, "", 0, "L", false, 0, "")

	pdf.SetFont("Arial", "", 9)

	pdf.CellFormat(65.3, 5, fakename.Add.Line1, "", 1, "L", false, 0, "")
	pdf.SetX(22.9)
	pdf.CellFormat(55, 5, fakename.Add.Line2+",", "", 1, "L", false, 0, "")
	pdf.SetX(22.9)
	pdf.CellFormat(55, 5, fakename.Add.Line3+",", "", 1, "L", false, 0, "")
	pdf.SetX(22.9)
	pdf.CellFormat(55, 5, fakename.Add.City+",", "", 1, "L", false, 0, "")
	pdf.SetX(22.9)
	pdf.CellFormat(55, 5, fakename.Add.State+",", "", 1, "L", false, 0, "")
	pdf.SetX(22.9)
	pdf.CellFormat(55, 5, fakename.Add.Country+",", "", 1, "L", false, 0, "")
	pdf.SetX(22.9)
	pdf.CellFormat(55, 5, fakename.Add.Pincode, "", 1, "L", false, 0, "")
	//to add space between address and table
	pdf.MultiCell(50, 10, "", "", "", false)

	pdf.SetFont("Arial", "B", 8)

	// Generate table header

	header := [colCount]string{"Sr.no", "Account No.", "Branch", "Name", "CCY", "Start Date", "Installment Amount", "Maturity Amount", "Date of Maturity", "Tenure (Months)", "Rate of Interest", "Current Principal Amount*"}
	colWd := []float64{10, 30, 30, 40, 13, 25, 25, 25, 25, 20, 20, 20}

	maxHt := lineHt
	y := pdf.GetY()
	for col, val := range header {
		cell.str = val
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[col]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[col] = cell
	}
	x := marginH
	for col := range header {
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
	srNo := 1
	for _, fund := range fakename.Passionfund {
		maxHt := lineHt
		// Cell height calculation loop
		count := 0
		cell.str = strconv.Itoa(srNo)
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.AccountNO
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.Branch
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.Name
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.CCY
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = formatDate(fund.StartDate)
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.InstallmentAmount
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.MaturityAmt
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = formatDate(fund.DateOfMaturity)
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.Tenure
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.RateOfInterest
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++

		cell.str = fund.CurrentPrincipalAmt
		cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		cell.ht = float64(len(cell.list)) * lineHt
		if cell.ht > maxHt {
			maxHt = cell.ht
		}
		cellList[count] = cell
		count++
		// Cell render loop
		count = 0

		x := marginH
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY := y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++

		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		pdf.Rect(x, y, colWd[count], maxHt+cellGap+cellGap, "D")
		cell = cellList[count]
		cellY = y + cellGap + (maxHt-cell.ht)/2
		for splitJ := 0; splitJ < len(cell.list); splitJ++ {
			pdf.SetXY(x+cellGap, cellY)
			pdf.CellFormat(colWd[count]-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
				"C", false, 0, "")
			cellY += lineHt
		}
		x += colWd[count]
		count++
		y += maxHt + cellGap + cellGap
		if y > pageMaxUsage {
			pdf.AddPage()
			y = newPageStart
		}
		srNo++
	}

	pdf.MultiCell(50, 10, "", "", "", false)
	pdf.SetFont("Arial", "", 9)
	currDepositBalance := "* Current Deposit Balance - Is the total installment amount paid till date towards funding of the Dream Deposit."
	line2 := "* In case of default/delay in payments of installments,the maturity value mentioned above will be different from the actual maturity value."

	pdf.MultiCell(0, 5, currDepositBalance, "", "", false)
	pdf.MultiCell(0, 5, line2, "", "", false)

	err = pdf.OutputFileAndClose("FDSummary.pdf")
	if err != nil {
		err = errors.New(literals.ErrCreatingPDF)
		return
	}

	return
}
