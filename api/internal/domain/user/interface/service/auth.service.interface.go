package serviceInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, input request.RequestLogin) (output usecase.CreateAuthOutputDTO, err error)
	RefreshToken(ctx context.Context, input request.RequestRefreshToken) (output usecase.CreateAuthOutputDTO, err error)
}
