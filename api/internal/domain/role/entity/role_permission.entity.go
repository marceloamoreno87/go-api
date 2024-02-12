package entity

import (
	"errors"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
)

type RolePermission struct {
	RoleId       int32
	PermissionId int32
	Role         *Role
	Permission   *PermissionEntity.Permission
}

func NewRolePermission(role *Role, permission *PermissionEntity.Permission) (rolePermission *RolePermission, err error) {
	rolePermission = &RolePermission{
		RoleId:       role.GetID(),
		PermissionId: permission.GetID(),
		Role:         role,
		Permission:   permission,
	}
	valid := rolePermission.Validate()
	if valid != nil {
		return nil, valid
	}

	return
}

func (r *RolePermission) Validate() (err error) {

	if r.RoleId == 0 {
		return errors.New("RoleId is required")
	}

	if r.PermissionId == 0 {
		return errors.New("PermissionId is required")
	}

	if r.Role == nil {
		return errors.New("Role is required")
	}
	if r.Permission == nil {
		return errors.New("Permission is required")
	}
	return
}

func (r *RolePermission) GetRole() *Role {
	return r.Role
}

func (r *RolePermission) GetPermission() *PermissionEntity.Permission {
	return r.Permission
}

func (r *RolePermission) SetRole(role *Role) {
	r.Role = role
}

func (r *RolePermission) SetPermission(permission *PermissionEntity.Permission) {
	r.Permission = permission
}

func (r *RolePermission) GetRoleID() int32 {
	return r.RoleId
}

func (r *RolePermission) GetPermissionID() int32 {
	return r.PermissionId
}

func (r *RolePermission) SetRoleID(roleId int32) {
	r.RoleId = roleId
}

func (r *RolePermission) SetPermissionID(permissionId int32) {
	r.PermissionId = permissionId
}
