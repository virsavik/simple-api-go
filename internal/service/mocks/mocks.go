package mocks

import (
	"gokiosk/internal/model"
	"gokiosk/internal/service/mocks/testdata"
)

// Fake data from testdata package
var invoiceSlice = testdata.Invoices
var userSlice = testdata.Users
var invoiceProductSlice = testdata.InvoiceProducts
var productSlice = testdata.Products

// Find storekeeper by id
func findStorekeeperByID(id string) *model.User {
	for _, user := range userSlice {
		if user.ID == id {
			return user
		}
	}
	return nil
}

// Check exists invoice_product has "invoice id"
func isExistsInvoiceProductWithInvoiceID(id string) bool {
	for _, invoiceProduct := range invoiceProductSlice {
		if invoiceProduct.InvoiceID == id {
			return true
		}
	}
	return false
}

func findProductByID(id string) *model.Product {
	for _, product := range productSlice {
		if product.ID == id {
			return product
		}
	}
	return nil
}
