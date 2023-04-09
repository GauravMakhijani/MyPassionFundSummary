package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/literals"
	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/jaswdr/faker"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)


type FileService interface {
    DownloadFileAsExcel(downloadRequest *model.FileDownloadRequest) (passionFundSummary model.PassionFundSummaryResponse,err error)
    DownloadFileAsPDF() error

}

type FileServiceImpl struct {
}


func NewFileService() FileService {
    return &FileServiceImpl{}
}

func generateFakeData() (user model.UserData){

    fake := faker.New()
    user.Name = fake.Person().FirstName()
    user.Address =  model.Address{
        City: fake.Address().City(),
        State: fake.Address().State(),
        Line1: fake.Address().StreetAddress(),
        Line2: fake.Address().SecondaryAddress(),
        Line3: fake.Address().StreetName(),
        PostalCode: fake.Address().PostCode(),
        Country: fake.Address().Country(),
    }
    user.PassionFunds = make([]model.PassionFund, 5)
    currTime := time.Now()
    for i:=0;i<5;i++{
        user.PassionFunds[i].AccountNo = strconv.Itoa(fake.RandomNumber(10))
        user.PassionFunds[i].BranchName = fake.Address().City()
        user.PassionFunds[i].Name = fake.Person().FirstName()
        user.PassionFunds[i].CCY = "INR"
        user.PassionFunds[i].StartDate = fake.Time().RFC3339(currTime)
        user.PassionFunds[i].InstallmentAmount = strconv.FormatFloat(fake.RandomFloat(2,1000,1000000),'f',2,64)
        user.PassionFunds[i].MaturityAmount = strconv.FormatFloat(fake.RandomFloat(2,1000,1000000),'f',2,64)
        user.PassionFunds[i].DataOfMaturity = fake.Time().RFC3339(currTime)
        user.PassionFunds[i].Tenure = strconv.Itoa(fake.RandomNumber(2))
        user.PassionFunds[i].RateOfInterest = strconv.FormatFloat(fake.RandomFloat(2,5,20),'f',2,64)
        user.PassionFunds[i].CurrentPrincipalAmount = strconv.FormatFloat(fake.RandomFloat(2,1000,1000000),'f',2,64)
    }
    return
}
func formatDate(date string) string {
    t,_ := time.Parse(time.RFC3339 ,date)
    return t.Format(literals.DateFormat)
}

func (f *FileServiceImpl) DownloadFileAsExcel(downloadRequest *model.FileDownloadRequest) (passionFundSummary model.PassionFundSummaryResponse,err error)  {
    file := excelize.NewFile()

    file.SetSheetName("Sheet1",literals.SheetName)

    userData:= generateFakeData()

    file.SetCellValue(literals.SheetName,"A1",literals.BankName)
    file.SetCellValue(literals.SheetName,"A3","Name:")
    file.SetCellValue(literals.SheetName,"A4","Address:")
    file.SetCellValue(literals.SheetName,"B3",userData.Name)
    file.SetCellValue(literals.SheetName,"B4",userData.Address.Line1)
    file.SetCellValue(literals.SheetName,"B5",userData.Address.Line2)
    file.SetCellValue(literals.SheetName,"B6",userData.Address.Line3)
    file.SetCellValue(literals.SheetName,"B7",userData.Address.City)
    file.SetCellValue(literals.SheetName,"B8",userData.Address.State)
    file.SetCellValue(literals.SheetName,"B9",userData.Address.Country)
    file.SetCellValue(literals.SheetName,"B10",userData.Address.PostalCode)
    file.MergeCell(literals.SheetName,"C6","L6")
    file.SetColWidth(literals.SheetName,"A","M",30)
    file.SetCellValue(literals.SheetName,"C6",literals.SummaryHeading)

    style, _ := file.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"}})

    file.SetCellStyle(literals.SheetName, "C6", "L6", style)

    tableCol := []string{"Sr.no", "Account No", "Branch Name", "Name", "CCY", "Start Date", "Installment Amount", "Maturity Amount", "Date of Maturity", "Tenure (Months)","Rate of Interest", "Current Principal Amount*"}

    ch := 'A'
    for _,col :=range tableCol {
        cell:= fmt.Sprintf("%c11",ch)
        file.SetCellValue(literals.SheetName,cell,col)
        ch++
    }

    i:= 12
    start:=i
    for sr,fund := range userData.PassionFunds{
        maturityAmount,_ := strconv.ParseFloat(fund.MaturityAmount,64)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("A%d",i),sr+1)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("B%d",i),fund.AccountNo)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("C%d",i),fund.BranchName)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("D%d",i),fund.Name)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("E%d",i),fund.CCY)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("F%d",i),formatDate(fund.StartDate))
        file.SetCellValue(literals.SheetName,fmt.Sprintf("G%d",i),fund.InstallmentAmount)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("H%d",i), maturityAmount)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("I%d",i),formatDate(fund.DataOfMaturity))
        file.SetCellValue(literals.SheetName,fmt.Sprintf("J%d",i),fund.Tenure)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("K%d",i),fund.RateOfInterest)
        file.SetCellValue(literals.SheetName,fmt.Sprintf("L%d",i),fund.CurrentPrincipalAmount)
        i++
    }

    file.SetCellValue(literals.SheetName,fmt.Sprintf("G%d",i),"TotalINR")
    formula := fmt.Sprintf("=SUM(H%d:H%d)",start,i-1)
    file.SetCellFormula(literals.SheetName,fmt.Sprintf("H%d",i),formula)

    i++
    file.SetCellValue(literals.SheetName,fmt.Sprintf("B%d",i),"Disclaimer-")
    i++
    file.MergeCell(literals.SheetName,fmt.Sprintf("B%d",i),fmt.Sprintf("B%d",i+1))
    file.MergeCell(literals.SheetName,fmt.Sprintf("B%d",i),fmt.Sprintf("E%d",i))
    i++
    file.SetCellValue(literals.SheetName,fmt.Sprintf("B%d",i),literals.Disclaimer)

    d := time.Now()
    dateOfCreation := d.Format("02Jan2006")
    fileName := downloadRequest.HashUserId + "_MyPassionFundSummary_"+dateOfCreation+ ".xlsx"

    if err= file.SaveAs(fileName);err!=nil{
        logrus.WithField("err", err.Error()).Error(err)
        err = errors.New(literals.ErrCreatingExcelFile)
        return
    }

    bytes,err:= ioutil.ReadFile("./"+fileName)
    if err != nil {
        logrus.WithField("err", err.Error()).Error(err)
        return
    }

    base64Encoding := base64.StdEncoding.EncodeToString(bytes)

    passionFundSummary.Body.ReportBytes = base64Encoding
    passionFundSummary.Body.ReportUrl = fileName


    return
}


func (f *FileServiceImpl) DownloadFileAsPDF() error {
    return nil
}
