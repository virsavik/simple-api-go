package fakedata

import (
	"gokiosk/internal/model"
	"time"
)

var GetAllFakeData = []model.Invoice{
	{
		ID:            "1",
		StorekeeperID: "KEEPER-0001",
		CreatedAt:     time.Unix(0, 0).UTC(),
		UpdatedAt:     time.Unix(0, 0).UTC(),
	},
	{
		ID:            "2",
		StorekeeperID: "KEEPER-0002",
		CreatedAt:     time.Unix(0, 0).UTC(),
		UpdatedAt:     time.Unix(0, 0).UTC(),
	},
}
