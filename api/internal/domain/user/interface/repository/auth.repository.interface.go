package repositoryInterface

import (
	"context"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
)

type AuthRepositoryInterface interface {
	CreateAuth(ctx context.Context, input entityInterface.AuthInterface) (err error)
	GetAuthByUserID(ctx context.Context, userId int32) (output entityInterface.AuthInterface, err error)
	UpdateAuthRevokeByUserID(ctx context.Context, id int32) (err error)
	GetAuthByToken(ctx context.Context, userId int32, token string) (output entityInterface.AuthInterface, err error)
	GetAuthByRefreshToken(ctx context.Context, userId int32, refreshToken string) (output entityInterface.AuthInterface, err error)
}
