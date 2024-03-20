package request

type RequestCreateRolePermission struct {
	RoleID        int32   `json:"role_id" validate:"number,required"`
	PermissionIDs []int32 `json:"permission_ids" validate:"required"`
}

type RequestGetRolePermission struct {
	RoleID int32 `json:"role_id" validate:"number,required"`
}

type RequestDeleteRolePermissionByRoleID struct {
	RoleID int32 `json:"role_id" validate:"number,required"`
}

type RequestUpdateRolePermission struct {
	RoleID        int32   `json:"role_id" validate:"number,required"`
	PermissionIDs []int32 `json:"permission_ids" validate:"required"`
}
