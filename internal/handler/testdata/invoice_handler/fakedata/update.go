package fakedata

import (
	"gokiosk/internal/model"
	"time"
)

var UpdateSuccessParamFake = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateSuccessResultFake = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorNotFoundParamFake = model.Invoice{
	ID:            "X",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorNotFoundResultFake = model.Invoice{
	ID:            "X",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorStoreKeeperIDNotFoundParamFake = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}

var UpdateErrorStoreKeeperIDNotFoundResultFake = model.Invoice{
	ID:            "1",
	StorekeeperID: "KEEPER_0003",
	CreatedAt:     time.Unix(0, 0).UTC(),
	UpdatedAt:     time.Unix(0, 0).UTC(),
}
