package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user entityInterface.UserInterface) (output entityInterface.UserInterface, err error)
	GetUser(ctx context.Context, id int32) (output entityInterface.UserInterface, err error)
	GetUserByEmail(ctx context.Context, email string) (output entityInterface.UserInterface, err error)
	GetUsers(ctx context.Context, limit int32, offset int32) (output []entityInterface.UserInterface, err error)
	UpdateUser(ctx context.Context, user entityInterface.UserInterface, id int32) (ouerr error)
	DeleteUser(ctx context.Context, id int32) (err error)
	UpdateUserPassword(ctx context.Context, id int32, password string) (err error)
	UpdateUserActive(ctx context.Context, id int32, active bool) (err error)
}
