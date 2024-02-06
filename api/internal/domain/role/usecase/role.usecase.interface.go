package usecase

type CreateRoleUseCaseInterface interface {
	Execute(input CreateRoleInputDTO) (output CreateRoleOutputDTO, err error)
}

type DeleteRoleUseCaseInterface interface {
	Execute(input DeleteRoleInputDTO) (output DeleteRoleOutputDTO, err error)
}

type UpdateRoleUseCaseInterface interface {
	Execute(input UpdateRoleInputDTO) (output UpdateRoleOutputDTO, err error)
}

type GetRoleUseCaseInterface interface {
	Execute(input GetRoleInputDTO) (output GetRoleOutputDTO, err error)
}

type GetRolesUseCaseInterface interface {
	Execute() (output GetRolesOutputDTO, err error)
}
