package request

type RequestCreateRole struct {
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestGetRole struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetRoles struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type RequestUpdateRole struct {
	ID           int32  `json:"id" validate:"number,required"`
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestDeleteRole struct {
	ID int32 `json:"id" validate:"number,required"`
}
