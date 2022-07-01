package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/model"
	"gokiosk/internal/repository/orm"
)

type InvoiceRepositoryMock struct {
	mock.Mock
}

func (m *InvoiceRepositoryMock) GetInvoicesByDuration(duration model.Duration) ([]orm.Invoice, error) {
	args := m.Called(duration)

	var invoices []orm.Invoice
	if args.Get(0) != nil {
		invoices = args.Get(0).([]orm.Invoice)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return invoices, rErr
}

func (m *InvoiceRepositoryMock) GetInvoiceProductsByInvoiceID(invoiceID string) ([]orm.InvoiceProduct, error) {
	args := m.Called(invoiceID)

	var invoiceProducts []orm.InvoiceProduct
	if args.Get(0) != nil {
		invoiceProducts = args.Get(0).([]orm.InvoiceProduct)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return invoiceProducts, rErr
}
