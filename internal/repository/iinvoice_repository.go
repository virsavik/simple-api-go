package repository

import (
	"gokiosk/internal/model"
	"gokiosk/internal/repository/orm"
)

type IInvoiceRepository interface {
	GetInvoicesByDuration(duration model.Duration) ([]orm.Invoice, error)

	GetInvoiceProductsByInvoiceID(invoiceID string) ([]orm.InvoiceProduct, error)
}
