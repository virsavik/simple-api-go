package mocks

import (
	"fmt"
	"github.com/gofrs/uuid"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/service"
	"time"
)

type InvoiceServiceMock struct{}

func NewInvoiceServiceMock() service.IInvoiceService {
	return InvoiceServiceMock{}
}

var invoiceSlice = model.InvoiceSlice{
	&model.Invoice{
		ID:            "0001",
		StorekeeperID: "KEEPER-0001",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
	&model.Invoice{
		ID:            "0002",
		StorekeeperID: "KEEPER-0002",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
	&model.Invoice{
		ID:            "0003",
		StorekeeperID: "KEEPER-0003",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
	&model.Invoice{
		ID:            "0004",
		StorekeeperID: "KEEPER-0003",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
	&model.Invoice{
		ID:            "0005",
		StorekeeperID: "KEEPER-0003",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
	&model.Invoice{
		ID:            "0006",
		StorekeeperID: "KEEPER-0004",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
}

func (i InvoiceServiceMock) GetAllInvoices(offset, limit int) (model.InvoiceSlice, error) {
	// Return nil slice and error OFFSET_AND_LIMIT_MUST_BE_POSITIVE if offset or limit is negative
	if offset < 0 || limit < 0 {
		return nil, errors.NewAppError(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE, "OFFSET_AND_LIMIT_MUST_BE_POSITIVE")
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
	return nil, fmt.Errorf("INVOICE_NOT_FOUND")
}

func (i InvoiceServiceMock) CreateInvoice(invoice model.Invoice) (*model.Invoice, error) {
	// Return nil and error if ID exists
	if invoice.ID != "" {
		return nil, fmt.Errorf("ID_MUST_BE_EMPTY")
	}
	invoice.ID = string(uuid.V4)
	invoiceSlice = append(invoiceSlice, &invoice)
	// ! Return invoice and nil if successful
	return &invoice, nil
}

func (i InvoiceServiceMock) UpdateInvoice(id string, invoice model.Invoice) (*model.Invoice, error) {
	// Return nil and error if ID does not exist or invoice id is different ID
	if invoice.ID != id {
		return nil, fmt.Errorf("ID_MUST_MATCH")
	}

	for i, elem := range invoiceSlice {
		if elem.ID == id {
			invoiceSlice[i] = &invoice
			return &invoice, nil
		}
	}
	return nil, fmt.Errorf("INVOICE_NOT_FOUND")
}

func (i InvoiceServiceMock) DeleteInvoice(id string) error {
	for i, elem := range invoiceSlice {
		if elem.ID == id {
			invoiceSlice = append(invoiceSlice[:i], invoiceSlice[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("INVOICE_NOT_FOUND")
}
