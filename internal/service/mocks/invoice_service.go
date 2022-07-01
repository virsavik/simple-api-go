package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/model"
	"gokiosk/internal/repository/orm"
	"time"
)

type InvoiceServiceMock struct {
	mock.Mock
}

func (i *InvoiceServiceMock) GetAllByPaginate(page model.Paginate) ([]orm.Invoice, error) {
	// this records that the method was called and passes in the value
	// it was called with
	args := i.Called(page)

	var slice []orm.Invoice
	// Get the first argument (the slice)
	if args.Get(0) != nil {
		slice = args.Get(0).([]orm.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return slice, rErr
}

func (i *InvoiceServiceMock) GetByID(id string) (orm.Invoice, error) {
	// Pass id as parameter and get arguments to return
	args := i.Called(id)

	var invoice orm.Invoice
	// Get the first argument (the invoice)
	if args.Get(0) != nil {
		invoice = args.Get(0).(orm.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return invoice, rErr
}

func (i *InvoiceServiceMock) GetAllByDuration(from time.Time, to time.Time) ([]orm.Invoice, error) {
	args := i.Called(from, to)

	var data []orm.Invoice
	if args.Get(0) != nil {
		data = args.Get(0).([]orm.Invoice)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}
	return data, rErr
}

func (i *InvoiceServiceMock) Create(invoice orm.Invoice) (orm.Invoice, error) {
	// Pass id as parameter and get arguments to return
	args := i.Called(invoice)

	var inv orm.Invoice
	// Get the first argument (the invoice)
	if args.Get(0) != nil {
		inv = args.Get(0).(orm.Invoice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return inv, rErr
}

func (i *InvoiceServiceMock) Update(id string, invoice orm.Invoice) (orm.Invoice, error) {
	// Pass id as parameter and get arguments to return
	args := i.Called(id, invoice)

	var inv orm.Invoice
	// Get the first argument (the invoice)
	if args.Get(0) != nil {
		inv = args.Get(0).(orm.Invoice)
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
