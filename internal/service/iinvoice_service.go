package service

import (
	"gokiosk/internal/model"
	"time"
)

type IInvoiceService interface {
	GetAllByPaginate(offset, limit int) ([]model.Invoice, error)

	GetByID(id string) (model.Invoice, error)

	GetAllByDuration(from time.Time, to time.Time) ([]model.Invoice, error)

	Create(invoice model.Invoice) (model.Invoice, error)

	Update(id string, invoice model.Invoice) (model.Invoice, error)

	DeleteByID(id string) error
}
