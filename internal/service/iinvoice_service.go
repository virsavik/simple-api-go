package service

import (
	"gokiosk/internal/model"
	"gokiosk/internal/repository/orm"
	"time"
)

type IInvoiceService interface {
	GetAllByPaginate(page model.Paginate) ([]orm.Invoice, error)

	GetByID(id string) (orm.Invoice, error)

	GetAllByDuration(from time.Time, to time.Time) ([]orm.Invoice, error)

	Create(invoice orm.Invoice) (orm.Invoice, error)

	Update(id string, invoice orm.Invoice) (orm.Invoice, error)

	DeleteByID(id string) error
}
