package usecase

type CreateUserUseCaseInterface interface {
	Execute(input CreateUserInputDTO) (output CreateUserOutputDTO, err error)
}

type GetUserUseCaseInterface interface {
	Execute(input GetUserInputDTO) (output GetUserOutputDTO, err error)
}

type GetUserByEmailUseCaseInterface interface {
	Execute(input GetUserByEmailInputDTO) (output GetUserByEmailOutputDTO, err error)
}

type GetUsersUseCaseInterface interface {
	Execute() (output GetUsersOutputDTO, err error)
}

type UpdateUserUseCaseInterface interface {
	Execute(input UpdateUserInputDTO) (output UpdateUserOutputDTO, err error)
}
type DeleteUserUseCaseInterface interface {
	Execute(input DeleteUserInputDTO) (output DeleteUserOutputDTO, err error)
}
