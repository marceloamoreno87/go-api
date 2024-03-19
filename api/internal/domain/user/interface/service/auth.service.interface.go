package serviceInterface

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/service"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthServiceInterface interface {
	Login(input service.RequestLoginInputDTO) (output usecase.CreateAuthOutputDTO, err error)
	RefreshToken(input service.RequestRefreshTokenInputDTO) (output usecase.CreateAuthOutputDTO, err error)
}
