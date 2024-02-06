-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password,
  role_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
  name = $1,
  email = $2,
  password = $3,
  role_id = $4
WHERE id = $5
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: GetRoleByInternalName :one
SELECT * FROM roles
WHERE internal_name = $1 LIMIT 1;

-- name: GetRoles :many
SELECT * FROM roles
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreateRole :one
INSERT INTO roles (
  name,
  internal_name,
  description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateRole :one
UPDATE roles SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4
RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles
WHERE id = $1
RETURNING *;

-- name: GetPermission :one
SELECT * FROM permissions
WHERE id = $1 LIMIT 1;

-- name: GetPermissions :many
SELECT * FROM permissions
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreatePermission :one
INSERT INTO permissions (
  name,
  internal_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdatePermission :one
UPDATE permissions SET
  name = $1,
  internal_name = $2
WHERE id = $3
RETURNING *;

-- name: DeletePermission :one
DELETE FROM permissions
WHERE id = $1
RETURNING *;

-- name: GetRolePermission :one
SELECT * FROM role_permissions
WHERE role_id = $1 AND permission_id = $2 LIMIT 1;

-- name: GetRolePermissions :many
SELECT * FROM role_permissions
ORDER BY role_id ASC, permission_id ASC
LIMIT $1 OFFSET $2;

-- name: CreateRolePermission :one
INSERT INTO role_permissions (
  role_id,
  permission_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteRolePermission :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2
RETURNING *;

-- name: GetRolePermissionsByRoleId :many
SELECT * FROM role_permissions
WHERE role_id = $1
ORDER BY permission_id ASC
LIMIT $2 OFFSET $3;

-- name: GetRolePermissionsByPermissionId :many
SELECT * FROM role_permissions
WHERE permission_id = $1
ORDER BY role_id ASC
LIMIT $2 OFFSET $3;