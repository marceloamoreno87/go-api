// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateAvatar(ctx context.Context, svg string) error
	CreatePermission(ctx context.Context, arg CreatePermissionParams) error
	CreateRole(ctx context.Context, arg CreateRoleParams) error
	CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) error
	CreateToken(ctx context.Context, arg CreateTokenParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	CreateValidationUser(ctx context.Context, arg CreateValidationUserParams) error
	DeleteAvatar(ctx context.Context, id int32) error
	DeletePermission(ctx context.Context, id int32) error
	DeleteRole(ctx context.Context, id int32) error
	DeleteRolePermission(ctx context.Context, roleID int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetAvatar(ctx context.Context, id int32) (Avatar, error)
	GetAvatars(ctx context.Context, arg GetAvatarsParams) ([]Avatar, error)
	GetPermission(ctx context.Context, id int32) (Permission, error)
	GetPermissionByInternalName(ctx context.Context, internalName string) (Permission, error)
	GetPermissions(ctx context.Context, arg GetPermissionsParams) ([]Permission, error)
	GetRole(ctx context.Context, id int32) (Role, error)
	GetRoleByInternalName(ctx context.Context, internalName string) (Role, error)
	GetRolePermission(ctx context.Context, id int32) ([]RolePermission, error)
	GetRolePermissionsByRole(ctx context.Context, roleID int32) ([]GetRolePermissionsByRoleRow, error)
	GetRoles(ctx context.Context, arg GetRolesParams) ([]Role, error)
	GetTokenByUser(ctx context.Context, userID int32) (Auth, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserWithAvatar(ctx context.Context, id int32) (GetUserWithAvatarRow, error)
	GetUserWithRole(ctx context.Context, id int32) (GetUserWithRoleRow, error)
	GetUserWithRoleAndAvatar(ctx context.Context, id int32) (GetUserWithRoleAndAvatarRow, error)
	GetUserWithValidationUser(ctx context.Context, id int32) (GetUserWithValidationUserRow, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	GetUsersWithAvatar(ctx context.Context, arg GetUsersWithAvatarParams) ([]GetUsersWithAvatarRow, error)
	GetUsersWithRole(ctx context.Context, arg GetUsersWithRoleParams) ([]GetUsersWithRoleRow, error)
	GetUsersWithRoleAndAvatar(ctx context.Context, arg GetUsersWithRoleAndAvatarParams) ([]GetUsersWithRoleAndAvatarRow, error)
	GetValidationUser(ctx context.Context, userID int32) (UsersValidation, error)
	GetValidationUserByHash(ctx context.Context, hash string) (UsersValidation, error)
	RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error)
	RevokeTokenByUser(ctx context.Context, userID int32) error
	UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) error
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserActive(ctx context.Context, arg UpdateUserActiveParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserValidationUsed(ctx context.Context, id int32) error
}

var _ Querier = (*Queries)(nil)
