package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type NewCreateAuthUseCaseInterface interface {
	Execute(input usecase.CreateAuthInputDTO) (output usecase.CreateAuthOutputDTO, err error)
}
type NewLoginUserUseCaseInterface interface {
	Execute(input usecase.LoginUserInputDTO) (output usecase.LoginUserOutputDTO, err error)
}

type NewGetAuthByUserIDUseCaseInterface interface {
	Execute(input usecase.GetAuthByUserIDInputDTO) (output usecase.GetAuthByUserIDOutputDTO, err error)
}

type NewGetAuthByTokenUseCaseInterface interface {
	Execute(input usecase.GetAuthByTokenInputDTO) (output usecase.GetAuthByTokenOutputDTO, err error)
}

type NewUpdateAuthRevokeUseCaseInterface interface {
	Execute(input usecase.UpdateAuthRevokeInputDTO) (output usecase.UpdateAuthRevokeOutputDTO, err error)
}

type NewGetAuthByRefreshTokenUseCase interface {
	Execute(input usecase.GetAuthByRefreshTokenInputDTO) (output usecase.GetAuthByRefreshTokenOutputDTO, err error)
}

type NewCheckTokenUseCaseInterface interface {
	Execute(input usecase.CheckTokenInputDTO) (output usecase.CheckTokenOutputDTO, err error)
}

type NewCheckRefreshTokenUseCaseInterface interface {
	Execute(input usecase.CheckRefreshTokenInputDTO) (output usecase.CheckRefreshTokenOutputDTO, err error)
}

type NewCreateUserUseCaseInterface interface {
	Execute(input usecase.CreateUserInputDTO) (output usecase.CreateUserOutputDTO, err error)
}

type NewUpdateUserPasswordUseCaseInterface interface {
	Execute(input usecase.UpdateUserPasswordInputDTO) (output usecase.UpdateUserPasswordOutputDTO, err error)
}
