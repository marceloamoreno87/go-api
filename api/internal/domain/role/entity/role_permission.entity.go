package entity

import (
	"errors"

	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
)

type RolePermission struct {
	RoleId        int32
	Role          *Role
	PermissionIds []int32
	Permissions   []*PermissionEntity.Permission
}

func NewRolePermission(role *Role, permissions []*PermissionEntity.Permission) (rolePermission *RolePermission, err error) {

	rolePermission = &RolePermission{
		Role:          role,
		RoleId:        role.ID,
		PermissionIds: make([]int32, len(permissions)),
		Permissions:   permissions,
	}

	for i, permission := range permissions {
		rolePermission.PermissionIds[i] = permission.ID
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

	if r.Role == nil {
		return errors.New("Role is required")
	}

	if len(r.PermissionIds) == 0 {
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

func (r *RolePermission) GetPermissionId() []int32 {
	return r.PermissionIds
}

func (r *RolePermission) GetRoleId() int32 {
	return r.RoleId
}

func (r *RolePermission) SetRole(role *Role) {
	r.Role = role
}

func (r *RolePermission) SetPermission(permissions []*PermissionEntity.Permission) {
	r.Permissions = permissions
}

func (r *RolePermission) SetPermissionId(permissionId []int32) {
	r.PermissionIds = permissionId
}

func (r *RolePermission) SetRoleId(roleId int32) {
	r.RoleId = roleId
}

func (r *RolePermission) AddPermission(permission *PermissionEntity.Permission) {
	r.Permissions = append(r.Permissions, permission)
	r.PermissionIds = append(r.PermissionIds, permission.ID)
}

func (r *RolePermission) RemovePermission(permission *PermissionEntity.Permission) {
	for i, p := range r.Permissions {
		if p.ID == permission.ID {
			r.Permissions = append(r.Permissions[:i], r.Permissions[i+1:]...)
			r.PermissionIds = append(r.PermissionIds[:i], r.PermissionIds[i+1:]...)
			break
		}
	}
}

func (r *RolePermission) RemovePermissionById(permissionId int32) {
	for i, p := range r.Permissions {
		if p.ID == permissionId {
			r.Permissions = append(r.Permissions[:i], r.Permissions[i+1:]...)
			r.PermissionIds = append(r.PermissionIds[:i], r.PermissionIds[i+1:]...)
			break
		}
	}
}

func (r *RolePermission) HasPermission(permission *PermissionEntity.Permission) bool {
	for _, p := range r.Permissions {
		if p.ID == permission.ID {
			return true
		}
	}
	return false
}

func (r *RolePermission) HasPermissionById(permissionId int32) bool {
	for _, p := range r.PermissionIds {
		if p == permissionId {
			return true
		}
	}
	return false
}

func (r *RolePermission) HasRole(role *Role) bool {
	return r.Role.ID == role.ID
}

func (r *RolePermission) HasRoleId(roleId int32) bool {
	return r.RoleId == roleId
}

func (r *RolePermission) HasPermissionName(permissionName string) bool {
	for _, p := range r.Permissions {
		if p.Name == permissionName {
			return true
		}
	}
	return false
}

func (r *RolePermission) HasRoleName(roleName string) bool {
	return r.Role.Name == roleName
}
