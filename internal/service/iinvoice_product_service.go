package service

import "gokiosk/internal/model"

type IInvoiceProductService interface {
	GetAllByPaginate(offset, limit int) (model.InvoiceProductSlice, error)

	GetByInvoiceIdAndProductID(invoiceID, productID string) (*model.InvoiceProduct, error)

	Create(item model.InvoiceProduct) (*model.InvoiceProduct, error)

	Update(id string, item model.InvoiceProduct) (*model.InvoiceProduct, error)

	DeleteById(id string) error
}
