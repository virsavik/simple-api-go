package fakedata

import (
	"gokiosk/internal/model"
	"time"
)

var CreateSuccessParamFake = model.Invoice{
	StorekeeperID: "KEEPER_0002",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var CreateSuccessResultFake = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0002",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var CreateErrorIDMustBeEmptyFakeData = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0002",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var CreateErrorStorekeeperIDNotFoundFakeData = model.Invoice{
	StorekeeperID: "KEEPER_XXXX",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}
