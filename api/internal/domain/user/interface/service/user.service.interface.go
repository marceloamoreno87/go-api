package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(input request.RequestCreateUser) (output usecase.CreateUserOutputDTO, err error)
	GetUser(input request.RequestGetUser) (output usecase.GetUserOutputDTO, err error)
	GetUsers(input request.RequestGetUsers) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(input request.RequestUpdateUser) (output usecase.UpdateUserOutputDTO, err error)
	DeleteUser(input request.RequestDeleteUser) (output usecase.DeleteUserOutputDTO, err error)
	UpdateUserPassword(input request.RequestUpdateUserPassword) (err error)
	ForgotPassword(input request.RequestForgotPassword) (err error)
	VerifyUser(input request.RequestVerifyUser) (err error)
}
