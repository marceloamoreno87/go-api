package request

type CreateRolePermissionRequest struct {
	RoleID        int32   `json:"role_id" validate:"number,required"`
	PermissionIDs []int32 `json:"permission_ids" validate:"required"`
}

type GetRolePermissionRequest struct {
	RoleID int32 `json:"role_id" validate:"number,required"`
}

type DeleteRolePermissionByRoleIDRequest struct {
	RoleID int32 `json:"role_id" validate:"number,required"`
}

type UpdateRolePermissionRequest struct {
	RoleID        int32   `json:"role_id" validate:"number,required"`
	PermissionIDs []int32 `json:"permission_ids" validate:"required"`
}
