package request

type CreateRoleRequest struct {
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type GetRoleRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}

type GetRolesRequest struct {
	Limit  int32 `json:"limit" validate:"number,required"`
	Offset int32 `json:"offset" validate:"number,required"`
}

type UpdateRoleRequest struct {
	ID           int32  `json:"id" validate:"number,required"`
	Name         string `json:"name" validate:"required"`
	InternalName string `json:"internal_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type DeleteRoleRequest struct {
	ID int32 `json:"id" validate:"number,required"`
}
