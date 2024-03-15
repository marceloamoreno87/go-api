package repositoryInterface

import (
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AuthRepositoryInterface interface {
	CreateAuth(auth entityInterface.AuthInterface) (err error)
	GetAuthByUserID(userId int32) (output entityInterface.AuthInterface, err error)
	UpdateAuthRevokeByUserID(id int32) (err error)
	GetAuthByToken(userId int32, token string) (output entityInterface.AuthInterface, err error)
	GetAuthByRefreshToken(userId int32, refreshToken string) (output entityInterface.AuthInterface, err error)
}
