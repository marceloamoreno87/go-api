package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserValidationRepositoryInterface interface {
	CreateUserValidation(ctx context.Context, userValidation entityInterface.UserValidationInterface) (output entityInterface.UserValidationInterface, err error)
	GetUserValidationByUserID(ctx context.Context, id int32) (output entityInterface.UserValidationInterface, err error)
	GetUserValidationByHash(ctx context.Context, hash string) (output entityInterface.UserValidationInterface, err error)
	UpdateUserValidationUsed(ctx context.Context, id int32) (err error)
}
