package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/model"
)

type InvoiceServiceMock struct {
	mock.Mock
}

func (i InvoiceServiceMock) GetAllByPaginate(offset, limit int) (model.InvoiceSlice, error) {
	// this records that the method was called and passes in the value
	// it was called with
	args := i.Called(offset, limit)

	var slice model.InvoiceSlice
	// Get the first argument (the slice)
	if args.Get(0) != nil {
		slice = args.Get(0).(model.InvoiceSlice)
	}

	var rErr error
	// Get the second argument (the error)
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return slice, rErr
}

func (i InvoiceServiceMock) GetById(id string) (*model.Invoice, error) {
	//// Find invoice by id
	//for _, invoice := range invoiceSlice {
	//	if invoice.ID == id {
	//		return invoice, nil
	//	}
	//}
	//// Will return nil and error if invoice not found
	//return nil, fmt.Errorf(errors.ERR_NOT_FOUND)
	return nil, nil
}

func (i InvoiceServiceMock) Create(invoice model.Invoice) (*model.Invoice, error) {
	//// Return error if storekeeper id is not exists
	//if findStorekeeperByID(invoice.StorekeeperID) == nil {
	//	return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	//}
	//
	//invoice.ID = fmt.Sprintf("%d", len(invoiceSlice)+1) // Generate ID
	//invoiceSlice = append(invoiceSlice, &invoice)       // Add invoice to slice
	//
	//// Return invoice and nil if successful
	//return &invoice, nil
	return nil, nil
}

func (i InvoiceServiceMock) Update(id string, invoice model.Invoice) (*model.Invoice, error) {
	//// Return error if storekeeper id is not exists
	//if findStorekeeperByID(invoice.StorekeeperID) == nil {
	//	return nil, fmt.Errorf(errors.ERR_RELATION_DOES_NOT_EXIST)
	//}
	//
	//// Return nil and error if ID does not exist or invoice id is different ID
	//if invoice.ID != id {
	//	return nil, fmt.Errorf(errors.ERR_ID_MUST_BE_MATCH)
	//}
	//
	//// For each invoice in slice, if id is same, update invoice
	//for i, elem := range invoiceSlice {
	//	if elem.ID == id {
	//		invoiceSlice[i] = &invoice
	//		return &invoice, nil
	//	}
	//}
	//
	//// Return error if invoice not found
	//return nil, fmt.Errorf(errors.ERR_NOT_FOUND)
	return nil, nil
}

func (i InvoiceServiceMock) DeleteById(id string) error {
	//// Return "error relation exists" if at least one invoice product with invoice id
	//if isExistsInvoiceProductWithInvoiceID(id) {
	//	return fmt.Errorf(errors.ERR_RELATION_EXISTS)
	//}
	//
	//// Find invoice by id
	//for i, elem := range invoiceSlice {
	//	if elem.ID == id {
	//		invoiceSlice = append(invoiceSlice[:i], invoiceSlice[i+1:]...)
	//		return nil
	//	}
	//}
	//
	//// Return error if invoice not found if id is not exists
	//return fmt.Errorf(errors.ERR_NOT_FOUND)
	return nil
}

func findStorekeeperByID(id string) *model.User {
	//for _, user := range userSlice {
	//	if user.ID == id {
	//		return user
	//	}
	//}
	//return nil
	return nil
}

func isExistsInvoiceProductWithInvoiceID(id string) bool {
	//for _, invoiceProduct := range invoiceProductSlice {
	//	if invoiceProduct.InvoiceID == id {
	//		return true
	//	}
	//}
	//return false
	return false
}
