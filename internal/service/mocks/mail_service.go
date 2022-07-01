package mocks

import "github.com/stretchr/testify/mock"

type MailServiceMock struct {
	mock.Mock
}

func (m *MailServiceMock) Send(to []string, msg []byte) error {
	args := m.Called(to, msg)

	var err error
	if args.Get(0) != nil {
		err = args.Error(0)
	}

	return err
}
