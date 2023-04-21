package service

import (
	"testing"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GenerateFakeData(t *testing.T) {
	tests := []struct {
		name     string
		wantType model.FakeName
	}{
		{
			name: "Test generateFakeData",
			wantType: model.FakeName{
				Name: "rutuja",
				Add: model.Address{
					Line1:   "ryerruru",
					Line2:   "wgeuweueu",
					Line3:   "dbcedb",
					City:    "pune",
					State:   "Maharashtra",
					Country: "India",
					Pincode: "23743636",
				},
				Passionfund: []model.FakeData{
					{
						AccountNO:           "4273482833",
						Branch:              "pune",
						Name:                "rutuja",
						CCY:                 "INR",
						StartDate:           "08/02/1988",
						InstallmentAmount:   "798228.12",
						MaturityAmt:         "39865.10",
						DateOfMaturity:      "11/06/1979",
						Tenure:              "12.3",
						RateOfInterest:      "3.4",
						CurrentPrincipalAmt: "365266",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData := GenerateFakeData()
			assert.IsType(t, tt.wantType, gotData)
		})
	}
}

func Test_formatDate(t *testing.T) {

	type args struct {
		date string
	}
	tests := []struct {
		name     string
		args     args
		wantType interface{}
	}{
		{
			name: "test format date",
			args: args{
				date: "2021-01-01T00:00:00Z",
			},
			wantType: "01 Jan 2021",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatDate(tt.args.date)
			assert.IsType(t, tt.wantType, got)
		})
	}
}

func Test_DownloadFile(t *testing.T) {
	type args struct {
		downloadRequest model.FileDownloadRequest
	}
	tests := []struct {
		name     string
		args     args
		f        *FileServiceImpl
		wantType model.PassionFundSummaryResponse
		wantErr  bool
	}{
		{
			name: "test download file",
			args: args{
				downloadRequest: model.FileDownloadRequest{
					HashUserId: "442762277",
					FormatType: "P",
				},
			},
			f:        &FileServiceImpl{},
			wantType: model.PassionFundSummaryResponse{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileServiceImpl{}
			got, err := f.DownloadFile(tt.args.downloadRequest)
			require.NoError(t, err)

			assert.IsType(t, tt.wantType, got)

		})
	}
}

func Test_GeneratePDF(t *testing.T) {
	type args struct {
		fakename        model.FakeName
		downloadRequest model.FileDownloadRequest
	}

	tests := []struct {
		name     string
		args     args
		wantType model.PassionFundSummaryResponse
		wantErr  bool
	}{
		{
			name: "test for generating pdf",
			args: args{
				fakename: model.FakeName{
					Name: "rutuja",
					Add: model.Address{
						Line1:   "dfbhfefh",
						Line2:   "shdgshdg",
						Line3:   "dhcshgd",
						City:    "pune",
						State:   "Maharashtra",
						Country: "india",
					},
					Passionfund: []model.FakeData{
						{
							AccountNO:           "2362731727",
							Branch:              "Pune Branch",
							Name:                "rutuja",
							CCY:                 "INR",
							StartDate:           "08/02/1988",
							InstallmentAmount:   "798228.12",
							MaturityAmt:         "456622",
							DateOfMaturity:      "11/06/1979",
							Tenure:              "2",
							RateOfInterest:      "16.20",
							CurrentPrincipalAmt: "257987.20",
						},
					},
				},
				downloadRequest: model.FileDownloadRequest{
					HashUserId: "328736717",
					FormatType: "P",
				},
			},
			wantType: model.PassionFundSummaryResponse{
				Body: model.PassionFundReport{
					ReportBytes: "",
					ReportUrl:   "453626727_FDSummary_10Apr2023.pdf",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePDF(tt.args.fakename, tt.args.downloadRequest)

			require.NoError(t, err)
			assert.IsType(t, tt.wantType, got)
		})
	}
}
