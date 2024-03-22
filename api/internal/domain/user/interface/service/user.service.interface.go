package serviceInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, input request.RequestCreateUser) (output usecase.CreateUserOutputDTO, err error)
	GetUser(ctx context.Context, input request.RequestGetUser) (output usecase.GetUserOutputDTO, err error)
	GetUsers(ctx context.Context, input request.RequestGetUsers) (output []usecase.GetUsersOutputDTO, err error)
	UpdateUser(ctx context.Context, input request.RequestUpdateUser) (output usecase.UpdateUserOutputDTO, err error)
	DeleteUser(ctx context.Context, input request.RequestDeleteUser) (output usecase.DeleteUserOutputDTO, err error)
	UpdateUserPassword(ctx context.Context, input request.RequestUpdateUserPassword) (err error)
	ForgotPassword(ctx context.Context, input request.RequestForgotPassword) (err error)
	VerifyUser(ctx context.Context, input request.RequestVerifyUser) (err error)
}
