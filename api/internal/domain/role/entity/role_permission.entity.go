package entity

import (
	"errors"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
)

type RolePermission struct {
	ID            int32
	RoleID        int32
	PermissionIDs []int32
	Role          *Role
	Permissions   []*PermissionEntity.Permission
}

func NewRolePermission(roleId int32, permissionIds []int32) (rolePermission *RolePermission, err error) {
	rolePermission = &RolePermission{
		RoleID:        roleId,
		PermissionIDs: permissionIds,
	}

	err = rolePermission.Validate()
	if err != nil {
		return
	}

	return
}

func (r *RolePermission) Validate() (err error) {

	if r.RoleID == 0 {
		return errors.New("RoleId is required")
	}

	if len(r.PermissionIDs) == 0 {
		return errors.New("PermissionId is required")
	}

	return
}

func (r *RolePermission) GetID() int32 {
	return r.ID
}

func (r *RolePermission) GetRole() *Role {
	return r.Role
}

func (r *RolePermission) GetPermission() []*PermissionEntity.Permission {
	return r.Permissions
}

func (r *RolePermission) GetPermissionIDs() []int32 {
	return r.PermissionIDs
}

func (r *RolePermission) GetRoleID() int32 {
	return r.RoleID
}

func (r *RolePermission) SetID(id int32) {
	r.ID = id
}

func (r *RolePermission) SetRole(role *Role) {
	r.Role = role
}

func (r *RolePermission) SetPermission(permissions []*PermissionEntity.Permission) {
	r.Permissions = permissions
}

func (r *RolePermission) SetPermissionIDs(permissionId []int32) {
	r.PermissionIDs = permissionId
}

func (r *RolePermission) SetRoleID(roleId int32) {
	r.RoleID = roleId
}
