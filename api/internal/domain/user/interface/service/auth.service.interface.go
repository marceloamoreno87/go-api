package serviceInterface

import (
	"io"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type AuthServiceInterface interface {
	Login(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error)
	RefreshToken(body io.ReadCloser) (output usecase.CreateAuthOutputDTO, err error)
}
