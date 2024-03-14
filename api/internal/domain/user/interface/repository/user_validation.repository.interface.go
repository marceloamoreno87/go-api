package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserValidationRepositoryInterface interface {
	CreateValidationUser(userValidation entityInterface.UserValidationInterface) (output entityInterface.UserValidationInterface, err error)
	GetValidationUser(id int32) (output entityInterface.UserValidationInterface, err error)
	GetValidationUserByHash(hash string) (output entityInterface.UserValidationInterface, err error)
	UpdateUserValidationUsed(id int32) (output entityInterface.UserValidationInterface, err error)
	config.SQLCInterface
}
