package service

import (
	"testing"

	"github.com/GauravMakhijani/MyPassionFundSummary/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_generateFakeData(t *testing.T) {
	tests := []struct {
		name     string
		wantType interface{}
	}{
		{
			name:     "Test generateFakeData",
			wantType: model.UserData{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData := generateFakeData()
			assert.IsType(t, tt.wantType, gotData)
		})
	}
}

func Test_formatDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test formatDate",
			args: args{
				date: "2021-01-01T00:00:00Z",
			},
			want: "01 Jan 2021",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate := formatDate(tt.args.date)
			assert.Equal(t, tt.want, gotDate)
		})
	}
}

func TestFileServiceImpl_DownloadFileAsExcel(t *testing.T) {
	type args struct {
		downloadRequest *model.FileDownloadRequest
	}
	tests := []struct {
		name                   string
		f                      *FileServiceImpl
		args                   args
		wantPassionFundSummary model.PassionFundSummaryResponse
		wantErr                bool
	}{
		// Positive | Test DownloadFileAsExcel
        {
            name: "Positive | Test DownloadFileAsExcel",
            f: &FileServiceImpl{},
            args: args{
                downloadRequest: &model.FileDownloadRequest{
                    HashUserId: "195809309",
                    FormatType: "E",
                },
            },
            wantPassionFundSummary: model.PassionFundSummaryResponse{},
            wantErr: false,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileServiceImpl{}
			gotPassionFundSummary, err := f.DownloadFileAsExcel(tt.args.downloadRequest)
            require.NoError(t, err)
			assert.IsType(t, tt.wantPassionFundSummary, gotPassionFundSummary)
		})
	}
}
