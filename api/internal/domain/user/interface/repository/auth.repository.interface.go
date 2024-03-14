package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AuthRepositoryInterface interface {
	CreateToken(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error)
	GetTokenByUser() (output entityInterface.AuthInterface, err error)
	RevokeTokenByUser(auth entityInterface.AuthInterface) (output entityInterface.AuthInterface, err error)
	config.SQLCInterface
}
