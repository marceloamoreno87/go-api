package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(input request.RequestCreateUserInputDTO) (output usecase.CreateUserOutputDTO, err error)
	GetUser(input request.RequestGetUserInputDTO) (output usecase.GetUserOutputDTO, err error)
	GetUsers(input request.RequestGetUsersInputDTO) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(input request.RequestUpdateUserInputDTO) (output usecase.UpdateUserOutputDTO, err error)
	DeleteUser(input request.RequestDeleteUserInputDTO) (output usecase.DeleteUserOutputDTO, err error)
	UpdateUserPassword(input request.RequestUpdateUserPasswordInputDTO) (err error)
	ForgotPassword(input request.RequestForgotPasswordInputDTO) (err error)
	VerifyUser(input request.RequestVerifyUserInputDTO) (err error)
}
