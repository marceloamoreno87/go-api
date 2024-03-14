package entityInterface

import "github.com/marceloamoreno/goapi/internal/shared/notification"

type RolePermissionInterface interface {
	Validate() (notify notification.ErrorsInterface)
	GetRole() RoleInterface
	GetPermissions() []PermissionInterface
	SetRole(role RoleInterface)
	SetPermissions(permissions []PermissionInterface)
	GetRoleID() (id int32)
	GetPermissionID() (id int32)
	SetRoleID(roleId int32)
	SetPermissionID(permissionId int32)
	GetPermissionIDs() []int32
	SetPermissionIDs(permissionId []int32)
	GetID() int32
	SetID(id int32)
}
