package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type GetRolePermissionsUseCaseInterface interface {
	Execute(input usecase.GetRolePermissionsInputDTO) (output []usecase.GetRolePermissionsOutputDTO, err error)
}

type CreateRolePermissionUseCaseInterface interface {
	Execute(input usecase.CreateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error)
}

type DeleteRolePermissionByRoleIDUseCaseInterface interface {
	Execute(input usecase.DeleteRolePermissionByRoleIDInputDTO) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error)
}
