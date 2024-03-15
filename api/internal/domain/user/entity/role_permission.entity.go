package entity

import (
	"errors"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
)

type RolePermission struct {
	ID            int32
	RoleID        int32
	Role          entityInterface.RoleInterface
	PermissionIDs []int32
	Permissions   []entityInterface.PermissionInterface
}

func NewRolePermission(roleId int32, permissionIds []int32) (rolePermission entityInterface.RolePermissionInterface, err error) {
	rolePermission = &RolePermission{
		RoleID:        roleId,
		PermissionIDs: permissionIds,
	}

	notify := rolePermission.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	return
}

func (r *RolePermission) Validate() (notify notification.ErrorsInterface) {

	notify = notification.New()

	if r.RoleID == 0 {
		notify.AddError("RoleID is required", "role_permission.entity.role_id")
	}

	if len(r.PermissionIDs) == 0 {
		notify.AddError("PermissionIDs is required", "role_permission.entity.permission_ids")
	}

	return
}

func (r *RolePermission) GetID() int32 {
	return r.ID
}

func (r *RolePermission) GetRoleID() (id int32) {
	return r.RoleID
}

func (r *RolePermission) GetRole() entityInterface.RoleInterface {
	return r.Role
}

func (r *RolePermission) GetRolePermissionID() (id int32) {
	return r.ID
}

func (r *RolePermission) GetPermissions() []entityInterface.PermissionInterface {
	return r.Permissions
}

func (r *RolePermission) GetPermissionIDs() []int32 {
	return r.PermissionIDs
}

func (r *RolePermission) SetRolePermissionID(id int32) {
	r.ID = id
}

func (r *RolePermission) SetRoleID(roleId int32) {
	r.RoleID = roleId
}

func (r *RolePermission) SetRole(role entityInterface.RoleInterface) {
	r.Role = role
}

func (r *RolePermission) SetPermissions(permissions []entityInterface.PermissionInterface) {
	r.Permissions = permissions
}

func (r *RolePermission) SetPermissionIDs(permissionId []int32) {
	r.PermissionIDs = permissionId
}
