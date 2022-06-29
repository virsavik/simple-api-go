package repository

import (
	"gokiosk/internal/model"
	"time"
)

type IInvoiceRepository interface {
	GetAllByDuration(from time.Time, to time.Time) ([]model.Invoice, error)

	GetInvoiceProductsByInvoiceID(invoiceID string) ([]model.InvoiceProduct, error)
}
