package user_interface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type LoginUseCaseInterface interface {
	Execute(input usecase.LoginInputDTO) (err error)
}
