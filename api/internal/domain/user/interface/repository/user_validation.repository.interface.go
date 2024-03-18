package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type UserValidationRepositoryInterface interface {
	CreateUserValidation(userValidation entityInterface.UserValidationInterface) (output entityInterface.UserValidationInterface, err error)
	GetUserValidationByUserID(id int32) (output entityInterface.UserValidationInterface, err error)
	GetUserValidationByHash(hash string) (output entityInterface.UserValidationInterface, err error)
	UpdateUserValidationUsed(id int32) (err error)
}
