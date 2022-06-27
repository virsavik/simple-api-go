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

var invoiceSlice = model.InvoiceSlice{
	&model.Invoice{
		ID:            "1",
		StorekeeperID: "KEEPER-0001",
	},
	&model.Invoice{
		ID:            "2",
		StorekeeperID: "KEEPER-0002",
	},
	&model.Invoice{
		ID:            "3",
		StorekeeperID: "KEEPER-0003",
	},
	&model.Invoice{
		ID:            "4",
		StorekeeperID: "KEEPER-0003",
	},
	&model.Invoice{
		ID:            "5",
		StorekeeperID: "KEEPER-0003",
	},
	&model.Invoice{
		ID:            "6",
		StorekeeperID: "KEEPER-0004",
	},
}

var invoiceProductSlice = model.InvoiceProductSlice{
	&model.InvoiceProduct{
		InvoiceID: "1",
		ProductID: "1",
	},
}

var userSlice = model.UserSlice{
	&model.User{
		ID:   "KEEPER-0001",
		Name: "Wilson",
	},
	&model.User{
		ID:   "KEEPER-0002",
		Name: "Knight",
	},
	&model.User{
		ID:   "KEEPER-0003",
		Name: "Hollow",
	},
	&model.User{
		ID:   "KEEPER-0004",
		Name: "Lion",
	},
}

func (i InvoiceServiceMock) GetAllInvoices(offset, limit int) (model.InvoiceSlice, error) {
	// Return nil slice and error OFFSET_AND_LIMIT_MUST_BE_POSITIVE if offset or limit is negative
	if offset < 0 || limit < 0 {
		return nil, fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE)
	}
	var slice model.InvoiceSlice
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
	if findStorekeeperByID(invoice.StorekeeperID) == nil {
		return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	}
	invoice.ID = fmt.Sprintf("%d", len(invoiceSlice)+1) // Generate ID
	invoiceSlice = append(invoiceSlice, &invoice)       // Add invoice to slice
	// ! Return invoice and nil if successful
	return &invoice, nil
}

func (i InvoiceServiceMock) UpdateInvoice(id string, invoice model.Invoice) (*model.Invoice, error) {
	if findStorekeeperByID(invoice.StorekeeperID) == nil {
		return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	}

	// Return nil and error if ID does not exist or invoice id is different ID
	if invoice.ID != id {
		return nil, fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH)
	}

	for i, elem := range invoiceSlice {
		if elem.ID == id {
			invoiceSlice[i] = &invoice
			return &invoice, nil
		}
	}
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

func findStorekeeperByID(id string) *model.User {
	for _, user := range userSlice {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func isExistsInvoiceProductWithInvoiceID(id string) bool {
	for _, invoiceProduct := range invoiceProductSlice {
		if invoiceProduct.InvoiceID == id {
			return true
		}
	}
	return false
}
