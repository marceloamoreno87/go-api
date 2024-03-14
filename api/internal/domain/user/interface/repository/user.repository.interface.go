package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserRepositoryInterface interface {
	CreateUser(user entityInterface.UserInterface) (err error)
	GetUser(id int32) (entityInterface.UserInterface, error)
	GetUserByEmail(email string) (entityInterface.UserInterface, error)
	GetUsers(limit int32, offset int32) ([]entityInterface.UserInterface, error)
	UpdateUser(user entityInterface.UserInterface, id int32) (err error)
	DeleteUser(id int32) (err error)
	RegisterUser(user entityInterface.UserInterface) (userOutput entityInterface.UserInterface, err error)
	UpdatePasswordUser(user entityInterface.UserInterface, id int32) (err error)
}
