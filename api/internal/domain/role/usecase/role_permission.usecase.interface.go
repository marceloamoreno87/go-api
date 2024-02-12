package usecase

type CreateRolePermissionsUseCaseInterface interface {
	Execute(input CreateRolePermissionInputDTO) (output CreateRolePermissionOutputDTO, err error)
}

type GetRolePermissionUseCaseInterface interface {
	Execute(input GetRolePermissionsInputDTO) (output GetRolePermissionsOutputDTO, err error)
}

type UpdateRolePermissionUseCaseInterface interface {
	Execute(input UpdateRolePermissionInputDTO) (output UpdateRolePermissionOutputDTO, err error)
}
