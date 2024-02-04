package repository

import "github.com/marceloamoreno/goapi/internal/domain/user/entity"

type UserRepositoryInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsers(limit int32, offset int32) ([]*entity.User, error)
	UpdateUser(user *entity.User, id int64) (*entity.User, error)
	DeleteUser(id int64) error
}
