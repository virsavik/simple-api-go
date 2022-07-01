package service

import (
	"github.com/stretchr/testify/require"
	"gokiosk/internal/model"
	"gokiosk/internal/repository/mocks"
	"gokiosk/internal/repository/orm"
	servmock "gokiosk/internal/service/mocks"
	"gokiosk/internal/service/testdata/invoice_service_testdata"
	"os"
	"testing"
	"time"
)

func TestInvoiceService_GetCSVReportByDuration(t *testing.T) {
	type GivenData struct {
		fakeData       []orm.Invoice
		duration       model.Duration
		adminMailAddrs []string
		mailContent    []byte
		mailErr        error
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
				fakeData:       invoice_service_testdata.FakeInvoice,
				adminMailAddrs: []string{"admin@example.com"},
				duration: model.Duration{
					From: time.Date(2006, 1, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				mailContent: []byte("Export csv successfully!"),
				mailErr:     nil,
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
			mailMock := new(servmock.MailServiceMock)
			mailMock.On("Send", tc.given.adminMailAddrs, tc.given.mailContent).Return(tc.given.mailErr)

			repoMock := new(mocks.InvoiceRepositoryMock)
			repoMock.On("GetInvoicesByDuration", tc.given.duration).Return(tc.given.fakeData, tc.err)

			service := NewInvoiceService(repoMock, mailMock)

			// When
			actualCSV, err := service.GetCSVReport(tc.given.duration)

			// Then
			require.Nil(t, err, "Should be not error")

			expectedData, err := os.ReadFile(tc.expect.csvPath)
			require.Nil(t, err, "Should be not error")

			require.Equal(t, expectedData, actualCSV, "Should be equal")
		})
	}
}
