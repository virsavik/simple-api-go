package mocks

import (
	"fmt"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/service"
)

type InvoiceServiceMock struct{}

func NewInvoiceServiceMock() service.IInvoiceService {
	return InvoiceServiceMock{}
}

func (i InvoiceServiceMock) GetAllInvoices(offset, limit int) (model.InvoiceSlice, error) {
	// Return nil slice and error OFFSET_AND_LIMIT_MUST_BE_POSITIVE if offset or limit is negative
	if offset < 0 || limit < 0 {
		return nil, fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE)
	}

	// Define slice of invoices
	var slice model.InvoiceSlice

	// Append invoice to slice if position is in range of offset and limit
	for i, invoice := range invoiceSlice {
		if i >= offset && i < offset+limit {
			slice = append(slice, invoice)
		}
	}
	return slice, nil
}

func (i InvoiceServiceMock) GetInvoice(id string) (*model.Invoice, error) {
	// Find invoice by id
	for _, invoice := range invoiceSlice {
		if invoice.ID == id {
			return invoice, nil
		}
	}
	// Will return nil and error if invoice not found
	return nil, fmt.Errorf(errors.ERR_NOT_FOUND)
}

func (i InvoiceServiceMock) CreateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	// Return error if storekeeper id is not exists
	if findStorekeeperByID(invoice.StorekeeperID) == nil {
		return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	}

	invoice.ID = fmt.Sprintf("%d", len(invoiceSlice)+1) // Generate ID
	invoiceSlice = append(invoiceSlice, &invoice)       // Add invoice to slice

	// Return invoice and nil if successful
	return &invoice, nil
}

func (i InvoiceServiceMock) UpdateInvoice(id string, invoice model.Invoice) (*model.Invoice, error) {
	// Return error if storekeeper id is not exists
	if findStorekeeperByID(invoice.StorekeeperID) == nil {
		return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	}

	// Return nil and error if ID does not exist or invoice id is different ID
	if invoice.ID != id {
		return nil, fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH)
	}

	// For each invoice in slice, if id is same, update invoice
	for i, elem := range invoiceSlice {
		if elem.ID == id {
			invoiceSlice[i] = &invoice
			return &invoice, nil
		}
	}

	// Return error if invoice not found
	return nil, fmt.Errorf(errors.ERR_NOT_FOUND)
}

func (i InvoiceServiceMock) DeleteInvoice(id string) error {
	// Return "error relation exists" if at least one invoice product with invoice id
	if isExistsInvoiceProductWithInvoiceID(id) {
		return fmt.Errorf(errors.ERR_RELATION_EXISTS)
	}

	// Find invoice by id
	for i, elem := range invoiceSlice {
		if elem.ID == id {
			invoiceSlice = append(invoiceSlice[:i], invoiceSlice[i+1:]...)
			return nil
		}
	}

	// Return error if invoice not found if id is not exists
	return fmt.Errorf(errors.ERR_NOT_FOUND)
}
