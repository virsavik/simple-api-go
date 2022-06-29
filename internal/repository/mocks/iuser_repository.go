package mocks

import (
	"github.com/stretchr/testify/mock"
	"gokiosk/internal/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetByID(id string) (model.User, error) {
	args := m.Called(id)

	var user model.User
	if args.Get(0) != nil {
		user = args.Get(0).(model.User)
	}

	var rErr error
	if args.Get(1) != nil {
		rErr = args.Error(1)
	}

	return user, rErr
}
