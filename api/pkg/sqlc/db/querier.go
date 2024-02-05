// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (RolePermission, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeletePermission(ctx context.Context, id int32) error
	DeleteRole(ctx context.Context, id int32) error
	DeleteRolePermission(ctx context.Context, arg DeleteRolePermissionParams) error
	DeleteUser(ctx context.Context, id int32) error
	GetPermission(ctx context.Context, id int32) (Permission, error)
	GetPermissions(ctx context.Context, arg GetPermissionsParams) ([]Permission, error)
	GetRole(ctx context.Context, id int32) (Role, error)
	GetRolePermission(ctx context.Context, arg GetRolePermissionParams) (RolePermission, error)
	GetRolePermissions(ctx context.Context, arg GetRolePermissionsParams) ([]RolePermission, error)
	GetRolePermissionsByPermissionId(ctx context.Context, arg GetRolePermissionsByPermissionIdParams) ([]RolePermission, error)
	GetRolePermissionsByRoleId(ctx context.Context, arg GetRolePermissionsByRoleIdParams) ([]RolePermission, error)
	GetRoles(ctx context.Context, arg GetRolesParams) ([]Role, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
