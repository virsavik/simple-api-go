package service

import (
	"github.com/stretchr/testify/require"
	"gokiosk/internal/model"
	"gokiosk/internal/repository/mocks"
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
				fakeData: []model.Invoice{
					{
						ID:            "1",
						StorekeeperID: "KEEPER_0001",
						CreatedAt:     time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), // 2006-01-02 15:04:05
						UpdatedAt:     time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), // 2006-01-02 15:04:05
					},
					{
						ID:            "2",
						StorekeeperID: "KEEPER_0001",
						CreatedAt:     time.Date(2006, 2, 2, 15, 4, 5, 0, time.UTC), // 2006-02-02 15:04:05
						UpdatedAt:     time.Date(2006, 2, 2, 15, 4, 5, 0, time.UTC), // 2006-02-02 15:04:05
					},
					{
						ID:            "3",
						StorekeeperID: "KEEPER_0002",
						CreatedAt:     time.Date(2010, 2, 3, 15, 4, 5, 0, time.UTC), // 2010-02-03 15:04:05
						UpdatedAt:     time.Date(2010, 2, 3, 15, 4, 5, 0, time.UTC), // 2010-02-03 15:04:05
					},
				},
			},
			expect: ExpectData{
				csvPath: "testdata/invoice_service/gen_csv-success.csv",
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
