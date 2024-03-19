package request

type RequestCreateRoleInputDTO struct {
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestGetRoleInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}

type RequestGetRolesInputDTO struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type RequestUpdateRoleInputDTO struct {
	ID           int32  `json:"id" validate:"number,required"`
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type RequestDeleteRoleInputDTO struct {
	ID int32 `json:"id" validate:"number,required"`
}
