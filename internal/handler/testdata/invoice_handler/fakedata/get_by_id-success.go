package fakedata

import (
	"gokiosk/internal/model"
	"time"
)

var GetByIDFakeData = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER-0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}
