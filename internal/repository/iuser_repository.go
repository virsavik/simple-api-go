package repository

import "gokiosk/internal/repository/orm"

type IUserRepository interface {
	GetByID(id string) (orm.User, error)
}
