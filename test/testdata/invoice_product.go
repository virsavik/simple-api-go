package testdata

import "gokiosk/internal/model"

var InvoiceProducts = model.InvoiceProductSlice{
	&model.InvoiceProduct{
		InvoiceID: "1",
		ProductID: "1",
	},
}
