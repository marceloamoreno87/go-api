package response

type GetRolePermissionResponse struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_id"`
}

type GetRolePermissionsResponse struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_id"`
}

type CreateRolePermissionResponse struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_id"`
}

type UpdateRolePermissionResponse struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_id"`
}

type DeleteRolePermissionByRoleIDResponse struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_id"`
}
