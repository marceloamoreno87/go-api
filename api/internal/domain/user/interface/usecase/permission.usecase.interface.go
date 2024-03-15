package usecaseInterface

import "github.com/marceloamoreno/goapi/internal/domain/user/usecase"

type GetPermissionUseCaseInterface interface {
	Execute(input usecase.GetPermissionInputDTO) (output usecase.GetPermissionOutputDTO, err error)
}

type GetPermissionsUseCaseInterface interface {
	Execute(input usecase.GetPermissionsInputDTO) (output []usecase.GetPermissionsOutputDTO, err error)
}

type CreatePermissionUseCaseInterface interface {
	Execute(input usecase.CreatePermissionInputDTO) (output usecase.CreatePermissionOutputDTO, err error)
}

type UpdatePermissionUseCaseInterface interface {
	Execute(input usecase.UpdatePermissionInputDTO) (output usecase.UpdatePermissionOutputDTO, err error)
}

type DeletePermissionUseCaseInterface interface {
	Execute(input usecase.DeletePermissionInputDTO) (output usecase.DeletePermissionOutputDTO, err error)
}

type GetPermissionByInternalNameUseCaseInterface interface {
	Execute(input usecase.GetPermissionByInternalNameInputDTO) (output usecase.GetPermissionByInternalNameOutputDTO, err error)
}
