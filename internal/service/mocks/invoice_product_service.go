package mocks

import (
	"fmt"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"gokiosk/internal/service"
)

type InvoiceProductServiceMock struct{}

func NewInvoiceProductServiceMock() service.IInvoiceProductService {
	return InvoiceProductServiceMock{}
}

func (i InvoiceProductServiceMock) GetAllByPaginate(offset, limit int) (model.InvoiceProductSlice, error) {
	// Return nil slice and error OFFSET_AND_LIMIT_MUST_BE_POSITIVE if offset or limit is negative
	if offset < 0 || limit < 0 {
		return nil, fmt.Errorf(errors.ERR_OFFSET_AND_LIMIT_MUST_BE_POSITIVE)
	}

	// Define slice of invoices
	var slice model.InvoiceProductSlice

	// Append invoice to slice if position is in range of offset and limit
	for i, elm := range invoiceProductSlice {
		if i >= offset && i < offset+limit {
			slice = append(slice, elm)
		}
	}
	return slice, nil
}

func (i InvoiceProductServiceMock) GetByInvoiceIdAndProductID(invID, prodID string) (*model.InvoiceProduct, error) {
	// Find invoice by id
	for _, elm := range invoiceProductSlice {
		if elm.InvoiceID == invID && elm.ProductID == prodID {
			return elm, nil
		}
	}
	// Will return nil and error if invoice not found
	return nil, fmt.Errorf(errors.ERR_NOT_FOUND)
}

func (i InvoiceProductServiceMock) Create(item model.InvoiceProduct) (*model.InvoiceProduct, error) {
	// Return error if invoice id is not exists
	if findStorekeeperByID(item.InvoiceID) == nil {
		return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	}
	if findProductByID(item.ProductID) == nil {
		return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	}

	invoiceProductSlice = append(invoiceProductSlice, &item) // Add invoice to slice

	// Return invoice and nil if successful
	return &item, nil
}

func (i InvoiceProductServiceMock) Update(id string, item model.InvoiceProduct) (*model.InvoiceProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceProductServiceMock) DeleteById(id string) error {
	//TODO implement me
	panic("implement me")
}
