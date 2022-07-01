package fakedata

import (
	"gokiosk/internal/repository/orm"
	"time"
)

var CreateSuccessParamFake = orm.Invoice{
	StorekeeperID: "KEEPER_0002",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var CreateSuccessResultFake = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0002",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var CreateErrorIDMustBeEmptyFakeData = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0002",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var CreateErrorStorekeeperIDNotFoundFakeData = orm.Invoice{
	StorekeeperID: "KEEPER_XXXX",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}
