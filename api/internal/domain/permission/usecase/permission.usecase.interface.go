package usecase

type CreatePermissionUseCaseInterface interface {
	Execute(input *CreatePermissionInputDTO) (*CreatePermissionOutputDTO, error)
}

type GetPermissionUseCaseInterface interface {
	Execute(id int32) (*GetPermissionOutputDTO, error)
}

type GetPermissionsUseCaseInterface interface {
	Execute(limit int32, offset int32) ([]*GetPermissionsOutputDTO, error)
}

type UpdatePermissionUseCaseInterface interface {
	Execute(id int32, input *UpdatePermissionInputDTO) (*UpdatePermissionOutputDTO, error)
}

type DeletePermissionUseCaseInterface interface {
	Execute(id int32) error
}
