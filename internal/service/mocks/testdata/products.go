package testdata

import (
	"github.com/volatiletech/null/v8"
	"gokiosk/internal/model"
)

var Products = model.ProductSlice{
	&model.Product{
		ID:          "1",
		Name:        "Product 1",
		Description: null.StringFrom("Description 1"),
		Amount:      10,
	},
	&model.Product{
		ID:          "2",
		Name:        "Product 2",
		Description: null.StringFrom("Description 2"),
		Amount:      10,
	},
	&model.Product{
		ID:          "3",
		Name:        "Product 3",
		Description: null.StringFrom("Description 3"),
		Amount:      10,
	},
}
