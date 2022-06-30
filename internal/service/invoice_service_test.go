package service

import (
	"github.com/stretchr/testify/require"
	"gokiosk/internal/model"
	"gokiosk/internal/repository/mocks"
	"gokiosk/internal/service/testdata/invoice_service_testdata"
	"os"
	"testing"
	"time"
)

func TestInvoiceService_GenCSVReportByDuration(t *testing.T) {
	type GivenData struct {
		fakeData []model.Invoice
	}

	type ExpectData struct {
		csvPath string
	}

	tsc := map[string]struct {
		given  GivenData
		expect ExpectData
		err    error
	}{
		"success": {
			given: GivenData{
				fakeData: invoice_service_testdata.FakeInvoice,
			},
			expect: ExpectData{
				csvPath: "testdata/invoice_service_testdata/gen_csv-success.csv",
			},
			err: nil,
		},
	}

	for desc, tc := range tsc {
		t.Run(desc, func(t *testing.T) {
			// Given
			repoMock := new(mocks.InvoiceRepositoryMock)
			fromTime := time.Date(2006, 1, 1, 0, 0, 0, 0, time.UTC)
			toTime := time.Date(2010, 2, 4, 0, 0, 0, 0, time.UTC)
			repoMock.On("GetAllByDuration", fromTime, toTime).Return(tc.given.fakeData, tc.err)
			service := NewInvoiceService(repoMock)

			// When
			actualCSV, err := service.GenCSVReportByDuration(fromTime, toTime)

			// Then
			require.Nil(t, err, "Should be not error")

			expectedData, err := os.ReadFile(tc.expect.csvPath)
			require.Nil(t, err, "Should be not error")

			require.Equal(t, expectedData, actualCSV, "Should be equal")
		})
	}
}
