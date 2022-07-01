package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/repository/orm"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetByID(id string) (orm.User, error) {
	args := m.Called(id)

	var user orm.User
	if args.Get(0) != nil {
		user = args.Get(0).(orm.User)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return user, rErr
}
