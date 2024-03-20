package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthServiceInterface interface {
	Login(input request.RequestLogin) (output usecase.CreateAuthOutputDTO, err error)
	RefreshToken(input request.RequestRefreshToken) (output usecase.CreateAuthOutputDTO, err error)
}
