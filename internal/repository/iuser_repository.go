package repository

import "gokiosk/internal/model"

type IUserRepository interface {
	GetByID(id string) (model.User, error)
}
