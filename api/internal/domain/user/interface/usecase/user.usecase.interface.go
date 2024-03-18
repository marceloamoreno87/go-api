package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type GetUserByEmailUseCaseInterface interface {
	Execute(input usecase.GetUserByEmailInputDTO) (output usecase.GetUserByEmailOutputDTO, err error)
}

type CreateUserUseCaseInterface interface {
	Execute(input usecase.CreateUserInputDTO) (output usecase.CreateUserOutputDTO, err error)
}

type GetUserUseCaseInterface interface {
	Execute(input usecase.GetUserInputDTO) (output usecase.GetUserOutputDTO, err error)
}

type GetUsersUseCaseInterface interface {
	Execute(input usecase.GetUsersInputDTO) (output []usecase.GetUsersOutputDTO, err error)
}

type UpdateUserUseCaseInterface interface {
	Execute(input usecase.UpdateUserInputDTO) (output usecase.UpdateUserOutputDTO, err error)
}

type DeleteUserUseCaseInterface interface {
	Execute(input usecase.DeleteUserInputDTO) (output usecase.DeleteUserOutputDTO, err error)
}

type UpdateUserPasswordUseCaseInterface interface {
	Execute(input usecase.UpdateUserPasswordInputDTO) (output usecase.UpdateUserPasswordOutputDTO, err error)
}

type CreateUserValidationUseCaseInterface interface {
	Execute(input usecase.CreateUserValidationInputDTO) (output usecase.CreateUserValidationOutputDTO, err error)
}


