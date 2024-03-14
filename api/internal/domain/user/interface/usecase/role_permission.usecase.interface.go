package user_interface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type CreateRolePermissionsUseCaseInterface interface {
	Execute(input usecase.CreateRolePermissionInputDTO) (err error)
}

type GetRolePermissionUseCaseInterface interface {
	Execute(input usecase.GetRolePermissionsInputDTO) (output usecase.GetRolePermissionsOutputDTO, err error)
}

type UpdateRolePermissionUseCaseInterface interface {
	Execute(input usecase.UpdateRolePermissionInputDTO) (err error)
}
