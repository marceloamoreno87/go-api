package request

type RequestGetPermissionInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetPermissionsInputDTO struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type RequestCreatePermissionInputDTO struct {
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestUpdatePermissionInputDTO struct {
	ID           int32  `json:"id" validate:"number,required"`
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestDeletePermissionInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}
