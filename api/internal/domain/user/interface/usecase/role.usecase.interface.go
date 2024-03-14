package user_interface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type CreateRoleUseCaseInterface interface {
	Execute(input usecase.CreateRoleInputDTO) (err error)
}

type GetRoleUseCaseInterface interface {
	Execute(input usecase.GetRoleInputDTO) (output usecase.GetRoleOutputDTO, err error)
}

type GetRoleByInternalNameUseCaseInterface interface {
	Execute(input usecase.GetRoleByInternalNameInputDTO) (output usecase.GetRoleByInternalNameOutputDTO, err error)
}

type GetRolesUseCaseInterface interface {
	Execute() (output usecase.GetRolesOutputDTO, err error)
}

type UpdateRoleUseCaseInterface interface {
	Execute(input usecase.UpdateRoleInputDTO) (err error)
}

type DeleteRoleUseCaseInterface interface {
	Execute(input usecase.DeleteRoleInputDTO) (err error)
}
