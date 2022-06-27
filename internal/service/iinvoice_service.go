package service

import "gokiosk/internal/model"

type IInvoiceService interface {
	GetAllInvoices(offset, limit int) (model.InvoiceSlice, error)

	GetInvoice(id string) (*model.Invoice, error)

	CreateInvoice(invoice model.Invoice) (*model.Invoice, error)

	UpdateInvoice(id string, invoice model.Invoice) (*model.Invoice, error)

	DeleteInvoice(id string) error
}
