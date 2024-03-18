package serviceInterface

import (
	"io"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type UserServiceInterface interface {
	GetUserById(id int32) (output usecase.GetUserOutputDTO, err error)
	GetUserByEmail(email string) (output usecase.GetUserByEmailOutputDTO, err error)
	GetUsers(limit int32, offset int32) (output []usecase.GetUsersOutputDTO, err error)
	CreateUser(body io.ReadCloser) (output usecase.CreateUserOutputDTO, err error)
	UpdateUser(id int32, body io.ReadCloser) (output usecase.UpdateUserOutputDTO, err error)
	UpdateUserPassword(body io.ReadCloser) (err error)
	DeleteUser(id int32) (output usecase.DeleteUserOutputDTO, err error)
	VerifyUser(body io.ReadCloser) (err error)
	ForgotPassword(body io.ReadCloser) (err error)
}
