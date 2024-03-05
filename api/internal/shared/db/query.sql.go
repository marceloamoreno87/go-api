// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const createAvatar = `-- name: CreateAvatar :exec
INSERT INTO avatars (
  svg
) VALUES (
  $1
)
`

func (q *Queries) CreateAvatar(ctx context.Context, svg string) error {
	_, err := q.exec(ctx, q.createAvatarStmt, createAvatar, svg)
	return err
}

const createPermission = `-- name: CreatePermission :exec
INSERT INTO permissions (
  name,
  internal_name,
  description
) VALUES (
  $1, $2, $3
)
`

type CreatePermissionParams struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) error {
	_, err := q.exec(ctx, q.createPermissionStmt, createPermission, arg.Name, arg.InternalName, arg.Description)
	return err
}

const createRole = `-- name: CreateRole :exec
INSERT INTO roles (
  name,
  internal_name,
  description
) VALUES (
  $1, $2, $3
)
`

type CreateRoleParams struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) error {
	_, err := q.exec(ctx, q.createRoleStmt, createRole, arg.Name, arg.InternalName, arg.Description)
	return err
}

const createRolePermission = `-- name: CreateRolePermission :exec
INSERT INTO role_permissions (
  role_id,
  permission_id
) VALUES (
  $1, $2
)
`

type CreateRolePermissionParams struct {
	RoleID       int32 `json:"role_id"`
	PermissionID int32 `json:"permission_id"`
}

func (q *Queries) CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) error {
	_, err := q.exec(ctx, q.createRolePermissionStmt, createRolePermission, arg.RoleID, arg.PermissionID)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
  name,
  email,
  password,
  active,
  role_id,
  avatar_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Active,
		arg.RoleID,
		arg.AvatarID,
	)
	return err
}

const createValidationUser = `-- name: CreateValidationUser :exec
INSERT INTO users_validation (
  user_id,
  hash,
  expires_in
) VALUES (
  $1, $2, $3
)
`

type CreateValidationUserParams struct {
	UserID    int32  `json:"user_id"`
	Hash      string `json:"hash"`
	ExpiresIn int32  `json:"expires_in"`
}

func (q *Queries) CreateValidationUser(ctx context.Context, arg CreateValidationUserParams) error {
	_, err := q.exec(ctx, q.createValidationUserStmt, createValidationUser, arg.UserID, arg.Hash, arg.ExpiresIn)
	return err
}

const deleteAvatar = `-- name: DeleteAvatar :exec
DELETE FROM avatars
WHERE id = $1
`

func (q *Queries) DeleteAvatar(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteAvatarStmt, deleteAvatar, id)
	return err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deletePermissionStmt, deletePermission, id)
	return err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteRoleStmt, deleteRole, id)
	return err
}

const deleteRolePermission = `-- name: DeleteRolePermission :exec
DELETE FROM role_permissions
WHERE role_id = $1
`

func (q *Queries) DeleteRolePermission(ctx context.Context, roleID int32) error {
	_, err := q.exec(ctx, q.deleteRolePermissionStmt, deleteRolePermission, roleID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getAvatar = `-- name: GetAvatar :one
SELECT id, svg, created_at, updated_at FROM avatars
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAvatar(ctx context.Context, id int32) (Avatar, error) {
	row := q.queryRow(ctx, q.getAvatarStmt, getAvatar, id)
	var i Avatar
	err := row.Scan(
		&i.ID,
		&i.Svg,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAvatars = `-- name: GetAvatars :many
SELECT id, svg, created_at, updated_at FROM avatars
ORDER BY id ASC
LIMIT $1 OFFSET $2
`

type GetAvatarsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAvatars(ctx context.Context, arg GetAvatarsParams) ([]Avatar, error) {
	rows, err := q.query(ctx, q.getAvatarsStmt, getAvatars, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Avatar
	for rows.Next() {
		var i Avatar
		if err := rows.Scan(
			&i.ID,
			&i.Svg,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPermission = `-- name: GetPermission :one
SELECT id, name, internal_name, description, created_at, updated_at FROM permissions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPermission(ctx context.Context, id int32) (Permission, error) {
	row := q.queryRow(ctx, q.getPermissionStmt, getPermission, id)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.InternalName,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPermissionByInternalName = `-- name: GetPermissionByInternalName :one
SELECT id, name, internal_name, description, created_at, updated_at FROM permissions
WHERE internal_name = $1 LIMIT 1
`

func (q *Queries) GetPermissionByInternalName(ctx context.Context, internalName string) (Permission, error) {
	row := q.queryRow(ctx, q.getPermissionByInternalNameStmt, getPermissionByInternalName, internalName)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.InternalName,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPermissions = `-- name: GetPermissions :many
SELECT id, name, internal_name, description, created_at, updated_at FROM permissions
ORDER BY id ASC
LIMIT $1 OFFSET $2
`

type GetPermissionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPermissions(ctx context.Context, arg GetPermissionsParams) ([]Permission, error) {
	rows, err := q.query(ctx, q.getPermissionsStmt, getPermissions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.InternalName,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRole = `-- name: GetRole :one
SELECT id, name, internal_name, description, created_at, updated_at FROM roles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, id int32) (Role, error) {
	row := q.queryRow(ctx, q.getRoleStmt, getRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.InternalName,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoleByInternalName = `-- name: GetRoleByInternalName :one
SELECT id, name, internal_name, description, created_at, updated_at FROM roles
WHERE internal_name = $1 LIMIT 1
`

func (q *Queries) GetRoleByInternalName(ctx context.Context, internalName string) (Role, error) {
	row := q.queryRow(ctx, q.getRoleByInternalNameStmt, getRoleByInternalName, internalName)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.InternalName,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRolePermission = `-- name: GetRolePermission :many
SELECT id, role_id, permission_id FROM role_permissions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRolePermission(ctx context.Context, id int32) ([]RolePermission, error) {
	rows, err := q.query(ctx, q.getRolePermissionStmt, getRolePermission, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RolePermission
	for rows.Next() {
		var i RolePermission
		if err := rows.Scan(&i.ID, &i.RoleID, &i.PermissionID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRolePermissionsByRole = `-- name: GetRolePermissionsByRole :many
SELECT role_permissions.id, role_id, permission_id, permissions.id, permissions.name, permissions.internal_name, permissions.description, permissions.created_at, permissions.updated_at, roles.id, roles.name, roles.internal_name, roles.description, roles.created_at, roles.updated_at FROM role_permissions
INNER JOIN permissions ON role_permissions.permission_id = permissions.id
INNER JOIN roles ON role_permissions.role_id = roles.id
WHERE role_id = $1
ORDER BY permission_id ASC
`

type GetRolePermissionsByRoleRow struct {
	ID             int32     `json:"id"`
	RoleID         int32     `json:"role_id"`
	PermissionID   int32     `json:"permission_id"`
	ID_2           int32     `json:"id_2"`
	Name           string    `json:"name"`
	InternalName   string    `json:"internal_name"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ID_3           int32     `json:"id_3"`
	Name_2         string    `json:"name_2"`
	InternalName_2 string    `json:"internal_name_2"`
	Description_2  string    `json:"description_2"`
	CreatedAt_2    time.Time `json:"created_at_2"`
	UpdatedAt_2    time.Time `json:"updated_at_2"`
}

func (q *Queries) GetRolePermissionsByRole(ctx context.Context, roleID int32) ([]GetRolePermissionsByRoleRow, error) {
	rows, err := q.query(ctx, q.getRolePermissionsByRoleStmt, getRolePermissionsByRole, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRolePermissionsByRoleRow
	for rows.Next() {
		var i GetRolePermissionsByRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.RoleID,
			&i.PermissionID,
			&i.ID_2,
			&i.Name,
			&i.InternalName,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_3,
			&i.Name_2,
			&i.InternalName_2,
			&i.Description_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoles = `-- name: GetRoles :many
SELECT id, name, internal_name, description, created_at, updated_at FROM roles
ORDER BY id ASC
LIMIT $1 OFFSET $2
`

type GetRolesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRoles(ctx context.Context, arg GetRolesParams) ([]Role, error) {
	rows, err := q.query(ctx, q.getRolesStmt, getRoles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.InternalName,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password, active, role_id, avatar_id, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password, active, role_id, avatar_id, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserWithAvatar = `-- name: GetUserWithAvatar :one
SELECT users.id, name, email, password, active, role_id, avatar_id, users.created_at, users.updated_at, avatars.id, svg, avatars.created_at, avatars.updated_at FROM users
INNER JOIN avatars ON users.id = avatars.user_id
WHERE users.id = $1 LIMIT 1
`

type GetUserWithAvatarRow struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Active      bool      `json:"active"`
	RoleID      int32     `json:"role_id"`
	AvatarID    int32     `json:"avatar_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ID_2        int32     `json:"id_2"`
	Svg         string    `json:"svg"`
	CreatedAt_2 time.Time `json:"created_at_2"`
	UpdatedAt_2 time.Time `json:"updated_at_2"`
}

func (q *Queries) GetUserWithAvatar(ctx context.Context, id int32) (GetUserWithAvatarRow, error) {
	row := q.queryRow(ctx, q.getUserWithAvatarStmt, getUserWithAvatar, id)
	var i GetUserWithAvatarRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.Svg,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
	)
	return i, err
}

const getUserWithRole = `-- name: GetUserWithRole :one
SELECT users.id, users.name, email, password, active, role_id, avatar_id, users.created_at, users.updated_at, roles.id, roles.name, internal_name, description, roles.created_at, roles.updated_at FROM users
INNER JOIN roles ON users.role_id = roles.id
WHERE users.id = $1 LIMIT 1
`

type GetUserWithRoleRow struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Active       bool      `json:"active"`
	RoleID       int32     `json:"role_id"`
	AvatarID     int32     `json:"avatar_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ID_2         int32     `json:"id_2"`
	Name_2       string    `json:"name_2"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt_2  time.Time `json:"created_at_2"`
	UpdatedAt_2  time.Time `json:"updated_at_2"`
}

func (q *Queries) GetUserWithRole(ctx context.Context, id int32) (GetUserWithRoleRow, error) {
	row := q.queryRow(ctx, q.getUserWithRoleStmt, getUserWithRole, id)
	var i GetUserWithRoleRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.Name_2,
		&i.InternalName,
		&i.Description,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
	)
	return i, err
}

const getUserWithRoleAndAvatar = `-- name: GetUserWithRoleAndAvatar :one
SELECT users.id, users.name, email, password, active, role_id, avatar_id, users.created_at, users.updated_at, roles.id, roles.name, internal_name, description, roles.created_at, roles.updated_at, avatars.id, svg, avatars.created_at, avatars.updated_at FROM users
INNER JOIN roles ON users.role_id = roles.id
INNER JOIN avatars ON users.id = avatars.user_id
WHERE users.id = $1 LIMIT 1
`

type GetUserWithRoleAndAvatarRow struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Active       bool      `json:"active"`
	RoleID       int32     `json:"role_id"`
	AvatarID     int32     `json:"avatar_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ID_2         int32     `json:"id_2"`
	Name_2       string    `json:"name_2"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt_2  time.Time `json:"created_at_2"`
	UpdatedAt_2  time.Time `json:"updated_at_2"`
	ID_3         int32     `json:"id_3"`
	Svg          string    `json:"svg"`
	CreatedAt_3  time.Time `json:"created_at_3"`
	UpdatedAt_3  time.Time `json:"updated_at_3"`
}

func (q *Queries) GetUserWithRoleAndAvatar(ctx context.Context, id int32) (GetUserWithRoleAndAvatarRow, error) {
	row := q.queryRow(ctx, q.getUserWithRoleAndAvatarStmt, getUserWithRoleAndAvatar, id)
	var i GetUserWithRoleAndAvatarRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.Name_2,
		&i.InternalName,
		&i.Description,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.ID_3,
		&i.Svg,
		&i.CreatedAt_3,
		&i.UpdatedAt_3,
	)
	return i, err
}

const getUserWithValidationUser = `-- name: GetUserWithValidationUser :one
SELECT users.id, name, email, password, active, role_id, avatar_id, users.created_at, updated_at, users_validation.id, user_id, hash, expires_in, users_validation.created_at FROM users
INNER JOIN users_validation ON users.id = users_validation.user_id
WHERE users.id = $1 ORDER BY users_validation.id DESC LIMIT 1
`

type GetUserWithValidationUserRow struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Active      bool      `json:"active"`
	RoleID      int32     `json:"role_id"`
	AvatarID    int32     `json:"avatar_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ID_2        int32     `json:"id_2"`
	UserID      int32     `json:"user_id"`
	Hash        string    `json:"hash"`
	ExpiresIn   int32     `json:"expires_in"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

func (q *Queries) GetUserWithValidationUser(ctx context.Context, id int32) (GetUserWithValidationUserRow, error) {
	row := q.queryRow(ctx, q.getUserWithValidationUserStmt, getUserWithValidationUser, id)
	var i GetUserWithValidationUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.UserID,
		&i.Hash,
		&i.ExpiresIn,
		&i.CreatedAt_2,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, email, password, active, role_id, avatar_id, created_at, updated_at FROM users
ORDER BY id ASC
LIMIT $1 OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.getUsersStmt, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Active,
			&i.RoleID,
			&i.AvatarID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersWithAvatar = `-- name: GetUsersWithAvatar :many
SELECT users.id, name, email, password, active, role_id, avatar_id, users.created_at, users.updated_at, avatars.id, svg, avatars.created_at, avatars.updated_at FROM users
INNER JOIN avatars ON users.id = avatars.user_id
ORDER BY users.id ASC
LIMIT $1 OFFSET $2
`

type GetUsersWithAvatarParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersWithAvatarRow struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Active      bool      `json:"active"`
	RoleID      int32     `json:"role_id"`
	AvatarID    int32     `json:"avatar_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ID_2        int32     `json:"id_2"`
	Svg         string    `json:"svg"`
	CreatedAt_2 time.Time `json:"created_at_2"`
	UpdatedAt_2 time.Time `json:"updated_at_2"`
}

func (q *Queries) GetUsersWithAvatar(ctx context.Context, arg GetUsersWithAvatarParams) ([]GetUsersWithAvatarRow, error) {
	rows, err := q.query(ctx, q.getUsersWithAvatarStmt, getUsersWithAvatar, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersWithAvatarRow
	for rows.Next() {
		var i GetUsersWithAvatarRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Active,
			&i.RoleID,
			&i.AvatarID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Svg,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersWithRole = `-- name: GetUsersWithRole :many
SELECT users.id, users.name, email, password, active, role_id, avatar_id, users.created_at, users.updated_at, roles.id, roles.name, internal_name, description, roles.created_at, roles.updated_at FROM users
INNER JOIN roles ON users.role_id = roles.id
ORDER BY users.id ASC
LIMIT $1 OFFSET $2
`

type GetUsersWithRoleParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersWithRoleRow struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Active       bool      `json:"active"`
	RoleID       int32     `json:"role_id"`
	AvatarID     int32     `json:"avatar_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ID_2         int32     `json:"id_2"`
	Name_2       string    `json:"name_2"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt_2  time.Time `json:"created_at_2"`
	UpdatedAt_2  time.Time `json:"updated_at_2"`
}

func (q *Queries) GetUsersWithRole(ctx context.Context, arg GetUsersWithRoleParams) ([]GetUsersWithRoleRow, error) {
	rows, err := q.query(ctx, q.getUsersWithRoleStmt, getUsersWithRole, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersWithRoleRow
	for rows.Next() {
		var i GetUsersWithRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Active,
			&i.RoleID,
			&i.AvatarID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Name_2,
			&i.InternalName,
			&i.Description,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersWithRoleAndAvatar = `-- name: GetUsersWithRoleAndAvatar :many
SELECT users.id, users.name, email, password, active, role_id, avatar_id, users.created_at, users.updated_at, roles.id, roles.name, internal_name, description, roles.created_at, roles.updated_at, avatars.id, svg, avatars.created_at, avatars.updated_at FROM users
INNER JOIN roles ON users.role_id = roles.id
INNER JOIN avatars ON users.id = avatars.user_id
ORDER BY users.id ASC
LIMIT $1 OFFSET $2
`

type GetUsersWithRoleAndAvatarParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersWithRoleAndAvatarRow struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Active       bool      `json:"active"`
	RoleID       int32     `json:"role_id"`
	AvatarID     int32     `json:"avatar_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ID_2         int32     `json:"id_2"`
	Name_2       string    `json:"name_2"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt_2  time.Time `json:"created_at_2"`
	UpdatedAt_2  time.Time `json:"updated_at_2"`
	ID_3         int32     `json:"id_3"`
	Svg          string    `json:"svg"`
	CreatedAt_3  time.Time `json:"created_at_3"`
	UpdatedAt_3  time.Time `json:"updated_at_3"`
}

func (q *Queries) GetUsersWithRoleAndAvatar(ctx context.Context, arg GetUsersWithRoleAndAvatarParams) ([]GetUsersWithRoleAndAvatarRow, error) {
	rows, err := q.query(ctx, q.getUsersWithRoleAndAvatarStmt, getUsersWithRoleAndAvatar, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersWithRoleAndAvatarRow
	for rows.Next() {
		var i GetUsersWithRoleAndAvatarRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Active,
			&i.RoleID,
			&i.AvatarID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Name_2,
			&i.InternalName,
			&i.Description,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.ID_3,
			&i.Svg,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getValidationUser = `-- name: GetValidationUser :one
SELECT id, user_id, hash, expires_in, created_at FROM users_validation
WHERE user_id = $1 ORDER BY id DESC LIMIT 1
`

func (q *Queries) GetValidationUser(ctx context.Context, userID int32) (UsersValidation, error) {
	row := q.queryRow(ctx, q.getValidationUserStmt, getValidationUser, userID)
	var i UsersValidation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Hash,
		&i.ExpiresIn,
		&i.CreatedAt,
	)
	return i, err
}

const getValidationUserByHash = `-- name: GetValidationUserByHash :one
SELECT id, user_id, hash, expires_in, created_at FROM users_validation
WHERE hash = $1 LIMIT 1
`

func (q *Queries) GetValidationUserByHash(ctx context.Context, hash string) (UsersValidation, error) {
	row := q.queryRow(ctx, q.getValidationUserByHashStmt, getValidationUserByHash, hash)
	var i UsersValidation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Hash,
		&i.ExpiresIn,
		&i.CreatedAt,
	)
	return i, err
}

const registerUser = `-- name: RegisterUser :one
INSERT INTO users (
  name,
  email,
  password,
  active,
  role_id,
  avatar_id
) VALUES (
  $1, $2, $3, $4, (select id from roles where internal_name = 'user'), (select id from avatars where id = 1)
)
RETURNING id, name, email, password, active, role_id, avatar_id, created_at, updated_at
`

type RegisterUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error) {
	row := q.queryRow(ctx, q.registerUserStmt, registerUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Active,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Active,
		&i.RoleID,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAvatar = `-- name: UpdateAvatar :exec
UPDATE avatars SET
  svg = $1
WHERE id = $2
`

type UpdateAvatarParams struct {
	Svg string `json:"svg"`
	ID  int32  `json:"id"`
}

func (q *Queries) UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) error {
	_, err := q.exec(ctx, q.updateAvatarStmt, updateAvatar, arg.Svg, arg.ID)
	return err
}

const updatePermission = `-- name: UpdatePermission :exec
UPDATE permissions SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4
`

type UpdatePermissionParams struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
	ID           int32  `json:"id"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error {
	_, err := q.exec(ctx, q.updatePermissionStmt, updatePermission,
		arg.Name,
		arg.InternalName,
		arg.Description,
		arg.ID,
	)
	return err
}

const updateRole = `-- name: UpdateRole :exec
UPDATE roles SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4
`

type UpdateRoleParams struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
	ID           int32  `json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.exec(ctx, q.updateRoleStmt, updateRole,
		arg.Name,
		arg.InternalName,
		arg.Description,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET
  name = $1,
  email = $2,
  password = $3,
  active = $4,
  role_id = $5,
  avatar_id = $6
WHERE id = $7
`

type UpdateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
	RoleID   int32  `json:"role_id"`
	AvatarID int32  `json:"avatar_id"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Active,
		arg.RoleID,
		arg.AvatarID,
		arg.ID,
	)
	return err
}
