package service

import "gokiosk/internal/model"

type IInvoiceService interface {
	GetAllByPaginate(offset, limit int) ([]model.Invoice, error)

	GetByID(id string) (model.Invoice, error)

	Create(invoice model.Invoice) (model.Invoice, error)

	Update(id string, invoice model.Invoice) (model.Invoice, error)

	DeleteByID(id string) error
}
