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

-- name: CreateUser :exec
INSERT INTO users (
  name,
  email,
  password,
  role_id
) VALUES (
  $1, $2, $3, $4
);

-- name: UpdateUser :exec
UPDATE users SET
  name = $1,
  email = $2,
  password = $3,
  role_id = $4
WHERE id = $5;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

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

-- name: CreateRole :exec
INSERT INTO roles (
  name,
  internal_name,
  description
) VALUES (
  $1, $2, $3
);

-- name: UpdateRole :exec
UPDATE roles SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;

-- name: GetPermission :one
SELECT * FROM permissions
WHERE id = $1 LIMIT 1;

-- name: GetPermissionByInternalName :one
SELECT * FROM permissions
WHERE internal_name = $1 LIMIT 1;

-- name: GetPermissions :many
SELECT * FROM permissions
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreatePermission :exec
INSERT INTO permissions (
  name,
  internal_name,
  description
) VALUES (
  $1, $2, $3
);

-- name: UpdatePermission :exec
UPDATE permissions SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1;

-- name: CreateRolePermission :exec
INSERT INTO role_permissions (
  role_id,
  permission_id
) VALUES (
  $1, $2
);

-- name: GetRolePermissions :many
SELECT * FROM role_permissions
INNER JOIN permissions ON role_permissions.permission_id = permissions.id
INNER JOIN roles ON role_permissions.role_id = roles.id
WHERE role_id = $1
ORDER BY permission_id ASC;

-- name: DeleteRolePermission :exec
DELETE FROM role_permissions
WHERE role_id = $1;