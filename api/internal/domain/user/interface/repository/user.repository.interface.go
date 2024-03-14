package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserRepositoryInterface interface {
	CreateUser(user entityInterface.UserInterface) (output entityInterface.UserInterface, err error)
	GetUser(id int32) (output entityInterface.UserInterface, err error)
	GetUserByEmail(email string) (output entityInterface.UserInterface, err error)
	GetUsers(limit int32, offset int32) (output []entityInterface.UserInterface, err error)
	UpdateUser(user entityInterface.UserInterface, id int32) (output entityInterface.UserInterface, err error)
	DeleteUser(id int32) (output entityInterface.UserInterface, err error)
	RegisterUser(user entityInterface.UserInterface) (output entityInterface.UserInterface, err error)
	UpdateUserPassword(user entityInterface.UserInterface, id int32) (output entityInterface.UserInterface, err error)
}
