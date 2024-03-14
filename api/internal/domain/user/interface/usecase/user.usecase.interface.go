package user_interface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type CreateUserUseCaseInterface interface {
	Execute(input usecase.CreateUserInputDTO) (err error)
}
type GetUserUseCaseInterface interface {
	Execute(input usecase.GetUserInputDTO) (output usecase.GetUserOutputDTO, err error)
}
type GetUserByEmailUseCaseInterface interface {
	Execute(input usecase.GetUserByEmailInputDTO) (output usecase.GetUserByEmailOutputDTO, err error)
}
type GetUsersUseCaseInterface interface {
	Execute() (output usecase.GetUsersOutputDTO, err error)
}
type UpdateUserUseCaseInterface interface {
	Execute(input usecase.UpdateUserInputDTO) (err error)
}
type DeleteUserUseCaseInterface interface {
	Execute(input usecase.DeleteUserInputDTO) (err error)
}
