package usecase

type CreateRoleUseCaseInterface interface {
	Execute(input CreateRoleInputDTO) (err error)
}

type GetRoleUseCaseInterface interface {
	Execute(input GetRoleInputDTO) (output GetRoleOutputDTO, err error)
}

type GetRoleByInternalNameUseCaseInterface interface {
	Execute(input GetRoleByInternalNameInputDTO) (output GetRoleByInternalNameOutputDTO, err error)
}

type GetRolesUseCaseInterface interface {
	Execute() (output GetRolesOutputDTO, err error)
}

type UpdateRoleUseCaseInterface interface {
	Execute(input UpdateRoleInputDTO) (err error)
}

type DeleteRoleUseCaseInterface interface {
	Execute(input DeleteRoleInputDTO) (err error)
}
