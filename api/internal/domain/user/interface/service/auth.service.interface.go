package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthServiceInterface interface {
	Login(input request.RequestLoginInputDTO) (output usecase.CreateAuthOutputDTO, err error)
	RefreshToken(input request.RequestRefreshTokenInputDTO) (output usecase.CreateAuthOutputDTO, err error)
}
