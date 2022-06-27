package testdata

import "gokiosk/internal/model"

var Invoices = model.InvoiceSlice{
	&model.Invoice{
		ID:            "1",
		StorekeeperID: "KEEPER-0001",
	},
	&model.Invoice{
		ID:            "2",
		StorekeeperID: "KEEPER-0002",
	},
	&model.Invoice{
		ID:            "3",
		StorekeeperID: "KEEPER-0003",
	},
	&model.Invoice{
		ID:            "4",
		StorekeeperID: "KEEPER-0003",
	},
	&model.Invoice{
		ID:            "5",
		StorekeeperID: "KEEPER-0003",
	},
	&model.Invoice{
		ID:            "6",
		StorekeeperID: "KEEPER-0004",
	},
}
