package fakedata

import (
	"gokiosk/internal/repository/orm"
	"time"
)

var UpdateSuccessParamFake = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateSuccessResultFake = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorNotFoundParamFake = orm.Invoice{
	ID:            "X",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorNotFoundResultFake = orm.Invoice{
	ID:            "X",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorStoreKeeperIDNotFoundParamFake = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorStoreKeeperIDNotFoundResultFake = orm.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}
