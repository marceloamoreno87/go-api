package entityInterface

import (
	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type RolePermissionInterface interface {
	GetID() int32
	GetRoleID() (id int32)
	GetRole() RoleInterface
	GetRolePermissionID() (id int32)
	GetPermissions() []PermissionInterface
	GetPermissionIDs() []int32
	SetRolePermissionID(id int32)
	SetRoleID(roleId int32)
	SetRole(role RoleInterface)
	SetPermissions(permissions []PermissionInterface)
	SetPermissionIDs(permissionId []int32)
	Validate() (notify notification.ErrorsInterface)
}
