package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/model"
)

type InvoiceServiceMock struct {
	mock.Mock
}

func (i *InvoiceServiceMock) GetAllByPaginate(offset, limit int) ([]model.Invoice, error) {
	// this records that the method was called and passes in the value
	// it was called with
	args := i.Called(offset, limit)

	var slice []model.Invoice
	// Get the first argument (the slice)
	if args.Get(0) != nil {
		slice = args.Get(0).([]model.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return slice, rErr
}

func (i *InvoiceServiceMock) GetByID(id string) (model.Invoice, error) {
	// Pass id as parameter and get arguments to return
	args := i.Called(id)

	var invoice model.Invoice
	// Get the first argument (the invoice)
	if args.Get(0) != nil {
		invoice = args.Get(0).(model.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return invoice, rErr
}

func (i *InvoiceServiceMock) Create(invoice model.Invoice) (model.Invoice, error) {
	// Pass id as parameter and get arguments to return
	args := i.Called(invoice)

	var inv model.Invoice
	// Get the first argument (the invoice)
	if args.Get(0) != nil {
		inv = args.Get(0).(model.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return inv, rErr
}

func (i *InvoiceServiceMock) Update(id string, invoice model.Invoice) (model.Invoice, error) {
	// Pass id as parameter and get arguments to return
	args := i.Called(id, invoice)

	var inv model.Invoice
	// Get the first argument (the invoice)
	if args.Get(0) != nil {
		inv = args.Get(0).(model.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return inv, rErr
}

func (i *InvoiceServiceMock) DeleteByID(id string) error {
	// Pass id as parameter and get arguments to return
	args := i.Called(id)

	var rErr error
	// Get the first argument (the error)
	if args.Get(0) != nil {
		rErr = args.Error(0)
	}

	return rErr
}
