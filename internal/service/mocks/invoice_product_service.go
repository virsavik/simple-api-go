package mocks

import (
	"gokiosk/internal/model"
	"gokiosk/internal/service"
)

type InvoiceProductServiceMock struct{}

func NewInvoiceProductServiceMock() service.IInvoiceProductService {
	return InvoiceProductServiceMock{}
}

func (i InvoiceProductServiceMock) GetAllByPaginate(offset, limit int) (model.InvoiceProductSlice, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceProductServiceMock) GetById(id string) (*model.InvoiceProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceProductServiceMock) Create(item model.InvoiceProduct) (*model.InvoiceProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceProductServiceMock) Update(id string, item model.InvoiceProduct) (*model.InvoiceProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvoiceProductServiceMock) DeleteById(id string) error {
	//TODO implement me
	panic("implement me")
}
