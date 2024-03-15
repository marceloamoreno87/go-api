package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AuthRepositoryInterface interface {
	CreateAuth(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error)
	GetAuthByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error)
	RevokeAuthByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error)
}
