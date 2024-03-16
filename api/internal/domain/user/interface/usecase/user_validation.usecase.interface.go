package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type GetUserValidationByHashUseCaseInterface interface {
	Execute(input usecase.GetUserValidationByHashInputDTO) (output usecase.GetUserValidationByHashOutputDTO, err error)
}
