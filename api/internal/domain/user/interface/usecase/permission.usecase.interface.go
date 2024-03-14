package user_interface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type CreatePermissionUseCaseInterface interface {
	Execute(input usecase.CreatePermissionInputDTO) (err error)
}

type GetPermissionUseCaseInterface interface {
	Execute(id int32) (output usecase.GetPermissionOutputDTO, err error)
}

type GetPermissionByInternalNameUseCaseInterface interface {
	Execute(input usecase.GetPermissionByInternalNameInputDTO) (output usecase.GetPermissionByInternalNameOutputDTO, err error)
}

type GetPermissionsUseCaseInterface interface {
	Execute(input usecase.GetPermissionsInputDTO) (output []usecase.GetPermissionsOutputDTO, err error)
}

type UpdatePermissionUseCaseInterface interface {
	Execute(id int32, input usecase.UpdatePermissionInputDTO) (err error)
}

type DeletePermissionUseCaseInterface interface {
	Execute(id int32) (err error)
}
