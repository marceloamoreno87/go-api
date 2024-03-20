package request

type RequestGetPermission struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetPermissions struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type RequestCreatePermission struct {
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestUpdatePermission struct {
	ID           int32  `json:"id" validate:"number,required"`
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestDeletePermission struct {
	ID int32 `json:"id" validate:"number,required"`
}
