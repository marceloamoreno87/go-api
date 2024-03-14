package repositoryInterface

import (
	"github.com/marceloamoreno/goapi/config"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AuthRepositoryInterface interface {
	CreateToken(auth entityInterface.AuthInterface) (err error)
	GetTokenByUser() (auth entityInterface.AuthInterface, err error)
	RevokeTokenByUser(auth entityInterface.AuthInterface) error
	config.SQLCInterface
}
