package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type GetRoleUseCaseInterface interface {
	Execute(input usecase.GetRoleInputDTO) (output usecase.GetRoleOutputDTO, err error)
}

type GetRolesUseCaseInterface interface {
	Execute(input usecase.GetRolesInputDTO) (output []usecase.GetRolesOutputDTO, err error)
}

type CreateRoleUseCaseInterface interface {
	Execute(input usecase.CreateRoleInputDTO) (output usecase.CreateRoleOutputDTO, err error)
}

type UpdateRoleUseCaseInterface interface {
	Execute(input usecase.UpdateRoleInputDTO) (output usecase.UpdateRoleOutputDTO, err error)
}

type DeleteRoleUseCaseInterface interface {
	Execute(input usecase.DeleteRoleInputDTO) (output usecase.DeleteRoleOutputDTO, err error)
}

type NewGetRoleByInternalNameUseCaseInterface interface {
	Execute(input usecase.GetRoleByInternalNameInputDTO) (output usecase.GetRoleByInternalNameOutputDTO, err error)
}
