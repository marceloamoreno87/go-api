package usecase

type CreatePermissionUseCaseInterface interface {
	Execute(input *CreatePermissionInputDTO) (err error)
}

type GetPermissionUseCaseInterface interface {
	Execute(id int32) (output *GetPermissionOutputDTO, err error)
}

type GetPermissionByInternalNameUseCaseInterface interface {
	Execute(input GetPermissionByInternalNameInputDTO) (output GetPermissionByInternalNameOutputDTO, err error)
}

type GetPermissionsUseCaseInterface interface {
	Execute(input GetPermissionsInputDTO) (output []*GetPermissionsOutputDTO, err error)
}

type UpdatePermissionUseCaseInterface interface {
	Execute(id int32, input *UpdatePermissionInputDTO) (err error)
}

type DeletePermissionUseCaseInterface interface {
	Execute(id int32) (err error)
}
