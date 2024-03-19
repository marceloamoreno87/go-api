package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(input service.RequestCreateUserInputDTO) (output usecase.CreateUserOutputDTO, err error)
	GetUser(input service.RequestGetUserInputDTO) (output usecase.GetUserOutputDTO, err error)
	GetUsers(input service.RequestGetUsersInputDTO) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(input service.RequestUpdateUserInputDTO) (output usecase.UpdateUserOutputDTO, err error)
	DeleteUser(input service.RequestDeleteUserInputDTO) (output usecase.DeleteUserOutputDTO, err error)
	UpdateUserPassword(input service.RequestUpdateUserPasswordInputDTO) (err error)
	ForgotPassword(input service.RequestForgotPasswordInputDTO) (err error)
	VerifyUser(input service.RequestVerifyUserInputDTO) (err error)
}
