package service

import (
	"fmt"
	"log"
	"strconv"
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
	user.Passionfund = make([]model.FakeData, 3)
	for i := 0; i < 3; i++ {
		user.Passionfund[i].AccountNO = strconv.Itoa(fake.RandomNumber(10))
		user.Passionfund[i].Branch = fake.Address().City()
		user.Passionfund[i].Name = fake.Person().FirstName()
		user.Passionfund[i].CCY = "INR"
		user.Passionfund[i].StartDate = fake.Time().RFC1123(time.Time{})
		user.Passionfund[i].InstallmentAmount = strconv.FormatFloat(fake.RandomFloat(2, 1000, 1000000), 'f', 2, 64)
		user.Passionfund[i].MaturityAmt = strconv.FormatFloat(fake.RandomFloat(2, 1000, 1000000), 'f', 2, 64)
		user.Passionfund[i].DateOfMaturity = fake.Time().RFC3339(time.Time{})
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

func GeneratePDF(fakename model.FakeName) error {

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	//////////////////////////////////////////////////////////////////
	const (
		colCount = 12
		marginH  = 1.0
		lineHt   = 5.5
		cellGap  = 2.0
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
	/////////////////////////////////////////////////////////////////////////

	//Styling
	pdf.SetFont("Arial", "", 10)
	pdf.SetMargins(2, 0, 1)
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
	pdf.SetFont("Arial", "B", 13)
	pdf.MultiCell(100, 10, "DREAM DEPOSIT SUMMARY", "", "R", false)
	// pdf.SetY(42)
	pdf.SetFont("Arial", "B", 10)
	name := "Name:"
	pdf.CellFormat(18, 5, name, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(65.3, 5, fakename.Name, "", 1, "L", false, 0, "")
	//pdf.SetY(42)
	pdf.SetFont("Arial", "B", 10)

	address := "Address:"

	pdf.CellFormat(18, 5, address, "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)

	pdf.CellFormat(65.3, 5, fakename.Add.Line1, "", 1, "L", false, 0, "")
	pdf.SetX(20.1)
	pdf.CellFormat(55, 5, fakename.Add.Line2+",", "", 1, "L", false, 0, "")
	pdf.SetX(20.1)
	pdf.CellFormat(55, 5, fakename.Add.Line3+",", "", 1, "L", false, 0, "")
	pdf.SetX(20.1)
	pdf.CellFormat(55, 5, fakename.Add.City+",", "", 1, "L", false, 0, "")
	pdf.SetX(20.1)
	pdf.CellFormat(55, 5, fakename.Add.State+",", "", 1, "L", false, 0, "")
	pdf.SetX(20.1)
	pdf.CellFormat(55, 5, fakename.Add.Country+",", "", 1, "L", false, 0, "")
	pdf.SetX(20.1)
	pdf.CellFormat(55, 5, fakename.Add.Pincode, "", 1, "L", false, 0, "")

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	header := [colCount]string{"Sr.no", "Account No.", "Branch Name", "Name", "CCY", "Start Date", "Installment Amount", "Maturity Amount", "Date of Maturity", "Tenure (Months)", "Rate of Interest", "Current Principal Amount*"}

	// strList := loremList()
	pdf.SetMargins(marginH, 15, marginH)
	pdf.SetFont("Arial", "", 14)
	//pdf.AddPage()

	colWd := []float64{15, 25, 20, 25, 10, 40, 30, 25, 25, 25, 25, 20}

	// Headers
	pdf.SetTextColor(224, 224, 224)
	pdf.SetFillColor(64, 64, 64)
	for colJ := 0; colJ < colCount; colJ++ {
		// cell.str = string(header[colJ])
		// 	cell.list = pdf.SplitLines([]byte(cell.str), colWd[count]-cellGap-cellGap)
		// 	cell.ht = float64(len(cell.list)) * lineHt
		// 	if cell.ht > maxHt {
		// 		maxHt = cell.ht
		// 	}
		// 	cellList[count] = cell
		pdf.CellFormat(colWd[colJ], 10, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetTextColor(24, 24, 24)
	pdf.SetFillColor(255, 255, 255)

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

		cell.str = fund.StartDate
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

		cell.str = fund.DateOfMaturity
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
		srNo++
	}

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	err := pdf.OutputFileAndClose("FDSummary.pdf")
	if err != nil {
		fmt.Println("ok here 4")

		return err
	}

	return nil

}

// func setText(w float64, h float64, pdf *gofpdf.Fpdf, text string) {
// 	x := pdf.GetX()
// 	y := pdf.GetY()
// 	width, _ := pdf.GetPageSize()
// 	_, _, right, _ := pdf.GetMargins()
// 	if x+w > width-right {
// 		x = pdf.GetX() - w + 3
// 		// move to the next line in the same cell
// 		y = y + h

// 		// 	// move to the next line in the same row
// 		//	pdf.SetXY(x, y+h)
// 		pdf.SetXY(x, y)

// 		// print the remaining text
// 		pdf.CellFormat(w, h, text, "1", 0, "C", false, 0, "")
// 	} else {
// 		// print the text in the current cell
// 		pdf.CellFormat(w, h, text, "1", 0, "C", false, 0, "")
// 	}
// }
