package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserValidationRepositoryInterface interface {
	CreateValidationUser(userValidation entityInterface.UserValidationInterface) (err error)
	GetValidationUser(id int32) (userValidation entityInterface.UserValidationInterface, err error)
	GetValidationUserByHash(hash string) (userValidation entityInterface.UserValidationInterface, err error)
	UpdateUserValidationUsed(id int32) (err error)
	config.SQLCInterface
}
