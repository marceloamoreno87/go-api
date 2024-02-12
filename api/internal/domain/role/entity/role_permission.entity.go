package entity

import (
	"errors"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
)

type RolePermission struct {
	RoleID        int32
	Role          *Role
	PermissionIDs []int32
	Permissions   []*PermissionEntity.Permission
}

func NewRolePermission(role *Role, permissions []*PermissionEntity.Permission) (rolePermission *RolePermission, err error) {

	rolePermission = &RolePermission{
		Role:          role,
		RoleID:        role.ID,
		PermissionIDs: make([]int32, len(permissions)),
		Permissions:   permissions,
	}

	for i, permission := range permissions {
		rolePermission.PermissionIDs[i] = permission.ID
	}

	valid := rolePermission.Validate()
	if valid != nil {
		return nil, valid
	}

	return
}

func (r *RolePermission) Validate() (err error) {

	if r.RoleID == 0 {
		return errors.New("RoleId is required")
	}

	if r.Role == nil {
		return errors.New("Role is required")
	}

	if len(r.PermissionIDs) == 0 {
		return errors.New("PermissionId is required")
	}

	if len(r.Permissions) == 0 {
		return errors.New("Permissions is required")
	}

	return
}

func (r *RolePermission) GetRole() *Role {
	return r.Role
}

func (r *RolePermission) GetPermission() []*PermissionEntity.Permission {
	return r.Permissions
}

func (r *RolePermission) GetPermissionID() []int32 {
	return r.PermissionIDs
}

func (r *RolePermission) GetRoleID() int32 {
	return r.RoleID
}

func (r *RolePermission) SetRole(role *Role) {
	r.Role = role
}

func (r *RolePermission) SetPermission(permissions []*PermissionEntity.Permission) {
	r.Permissions = permissions
}

func (r *RolePermission) SetPermissionID(permissionId []int32) {
	r.PermissionIDs = permissionId
}

func (r *RolePermission) SetRoleID(roleId int32) {
	r.RoleID = roleId
}
