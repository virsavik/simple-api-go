package fakedata

import (
	"gokiosk/internal/repository/orm"
	"time"
)

var GetByIDFakeData = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER-0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}
