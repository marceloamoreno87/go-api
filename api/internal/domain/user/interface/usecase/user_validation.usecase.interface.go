package usecaseInterface

import (
	"context"

	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type GetUserValidationByHashUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetUserValidationByHashInputDTO) (output usecase.GetUserValidationByHashOutputDTO, err error)
}

type GetUserValidationByUserIDUseCaseInterface interface {
	Execute(ctx context.Context, input usecase.GetUserValidationByUserIDInputDTO) (output usecase.GetUserValidationByUserIDOutputDTO, err error)
}
