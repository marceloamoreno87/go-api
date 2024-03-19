package request

type RequestCreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id" validate:"number,required"`
	PermissionIDs []int32 `json:"permission_ids" validate:"required"`
}

type RequestGetRolePermissionInputDTO struct {
	RoleID int32 `json:"role_id" validate:"number,required"`
}

type RequestDeleteRolePermissionByRoleIDInputDTO struct {
	RoleID int32 `json:"role_id" validate:"number,required"`
}

type RequestUpdateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id" validate:"number,required"`
	PermissionIDs []int32 `json:"permission_ids" validate:"required"`
}
