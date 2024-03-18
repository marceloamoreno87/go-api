package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type CreateAuthUseCaseInterface interface {
	Execute(input usecase.CreateAuthInputDTO) (output usecase.CreateAuthOutputDTO, err error)
}
type LoginUserUseCaseInterface interface {
	Execute(input usecase.LoginUserInputDTO) (output usecase.LoginUserOutputDTO, err error)
}

type GetAuthByUserIDUseCaseInterface interface {
	Execute(input usecase.GetAuthByUserIDInputDTO) (output usecase.GetAuthByUserIDOutputDTO, err error)
}

type GetAuthByTokenUseCaseInterface interface {
	Execute(input usecase.GetAuthByTokenInputDTO) (output usecase.GetAuthByTokenOutputDTO, err error)
}

type UpdateAuthRevokeUseCaseInterface interface {
	Execute(input usecase.UpdateAuthRevokeInputDTO) (output usecase.UpdateAuthRevokeOutputDTO, err error)
}

type GetAuthByRefreshTokenUseCase interface {
	Execute(input usecase.GetAuthByRefreshTokenInputDTO) (output usecase.GetAuthByRefreshTokenOutputDTO, err error)
}
