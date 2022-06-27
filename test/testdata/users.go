package testdata

import "gokiosk/internal/model"

var Users = model.UserSlice{
	&model.User{
		ID:   "KEEPER-0001",
		Name: "Wilson",
	},
	&model.User{
		ID:   "KEEPER-0002",
		Name: "Knight",
	},
	&model.User{
		ID:   "KEEPER-0003",
		Name: "Hollow",
	},
	&model.User{
		ID:   "KEEPER-0004",
		Name: "Lion",
	},
}
