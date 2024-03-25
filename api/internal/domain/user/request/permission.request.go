package request

type GetPermissionRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}

type GetPermissionsRequest struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type CreatePermissionRequest struct {
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type UpdatePermissionRequest struct {
	ID           int32  `json:"id" validate:"number,required"`
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type DeletePermissionRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}
