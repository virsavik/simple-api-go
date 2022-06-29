package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/model"
	"time"
)

type InvoiceRepositoryMock struct {
	mock.Mock
}

func (m *InvoiceRepositoryMock) GetAllByDuration(from time.Time, to time.Time) ([]model.Invoice, error) {
	args := m.Called(from, to)

	var invoices []model.Invoice
	if args.Get(0) != nil {
		invoices = args.Get(0).([]model.Invoice)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return invoices, rErr
}

func (m *InvoiceRepositoryMock) GetInvoiceProductsByInvoiceID(invoiceID string) ([]model.InvoiceProduct, error) {
	args := m.Called(invoiceID)

	var invoiceProducts []model.InvoiceProduct
	if args.Get(0) != nil {
		invoiceProducts = args.Get(0).([]model.InvoiceProduct)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return invoiceProducts, rErr
}
