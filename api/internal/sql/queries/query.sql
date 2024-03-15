-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserWithRole :one
SELECT * FROM users
INNER JOIN roles ON users.role_id = roles.id
WHERE users.id = $1 LIMIT 1;

-- name: GetUserWithAvatar :one
SELECT * FROM users
INNER JOIN avatars ON users.id = avatars.user_id
WHERE users.id = $1 LIMIT 1;

-- name: GetUserWithRoleAndAvatar :one
SELECT * FROM users
INNER JOIN roles ON users.role_id = roles.id
INNER JOIN avatars ON users.id = avatars.user_id
WHERE users.id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: GetUsersWithRole :many
SELECT * FROM users
INNER JOIN roles ON users.role_id = roles.id
ORDER BY users.id ASC
LIMIT $1 OFFSET $2;

-- name: GetUserWithValidationUser :one
SELECT * FROM users
INNER JOIN users_validation ON users.id = users_validation.user_id
WHERE users.id = $1 ORDER BY users_validation.id DESC LIMIT 1;

-- name: GetUsersWithAvatar :many
SELECT * FROM users
INNER JOIN avatars ON users.id = avatars.user_id
ORDER BY users.id ASC
LIMIT $1 OFFSET $2;

-- name: GetUsersWithRoleAndAvatar :many
SELECT * FROM users
INNER JOIN roles ON users.role_id = roles.id
INNER JOIN avatars ON users.id = avatars.user_id
ORDER BY users.id ASC
LIMIT $1 OFFSET $2;

-- name: CreateUser :exec
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
;

-- name: UpdateUser :exec
UPDATE users SET
  name = $1,
  email = $2,
  password = $3,
  active = $4,
  role_id = $5,
  avatar_id = $6
WHERE id = $7
;

-- name: UpdateUserPassword :exec
UPDATE users SET
  password = $1
WHERE id = $2
;

-- name: UpdateUserActive :exec
UPDATE users SET
  active = $1
WHERE id = $2
;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
;

-- name: GetAuthByUserID :one
SELECT * FROM auth
WHERE user_id = $1 and active is true ORDER BY id DESC LIMIT 1;

-- name: GetAuthByToken :one
SELECT * FROM auth
WHERE token = $1 and user_id = $2 and active is true ORDER BY id DESC LIMIT 1;

-- name: GetAuthByRefreshToken :one
SELECT * FROM auth
WHERE refresh_token = $1 and user_id = $2 and active is true ORDER BY id DESC LIMIT 1;

-- name: CreateAuth :exec
INSERT INTO auth (
  user_id,
  token,
  refresh_token,
  token_expires_in,
  refresh_token_expires_in
) VALUES (
  $1, $2, $3, $4, $5
)
;

-- name: UpdateAuthRevokeByUserID :exec
UPDATE auth SET
  active = false
WHERE user_id = $1
;

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
)
;

-- name: UpdateRole :exec
UPDATE roles SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4
;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1
;

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
)
;

-- name: UpdatePermission :exec
UPDATE permissions SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4
;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1
;

-- name: CreateRolePermission :exec
INSERT INTO role_permissions (
  role_id,
  permission_id
) VALUES (
  $1, $2
)
;

-- name: GetRolePermission :many
SELECT * FROM role_permissions
WHERE id = $1 LIMIT 1;

-- name: GetRolePermissionsByRole :many
SELECT * FROM role_permissions
INNER JOIN permissions ON role_permissions.permission_id = permissions.id
INNER JOIN roles ON role_permissions.role_id = roles.id
WHERE role_id = $1
ORDER BY permission_id ASC;

-- name: DeleteRolePermission :exec
DELETE FROM role_permissions
WHERE role_id = $1
;

-- name: GetAvatar :one
SELECT * FROM avatars
WHERE id = $1 LIMIT 1;

-- name: GetAvatars :many
SELECT * FROM avatars
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreateAvatar :exec
INSERT INTO avatars (
  svg
) VALUES (
  $1
)
;

-- name: UpdateAvatar :exec
UPDATE avatars SET
  svg = $1
WHERE id = $2
;

-- name: DeleteAvatar :exec
DELETE FROM avatars
WHERE id = $1
;

-- name: GetValidationUser :one
SELECT * FROM users_validation
WHERE user_id = $1 ORDER BY id DESC LIMIT 1;

-- name: GetValidationUserByHash :one
SELECT * FROM users_validation
WHERE hash = $1 and used is false LIMIT 1;

-- name: CreateValidationUser :exec
INSERT INTO users_validation (
  user_id,
  hash,
  expires_in
) VALUES (
  $1, $2, $3
)
;

-- name: UpdateUserValidationUsed :exec
UPDATE users_validation SET
  used = true
WHERE id = $1
;
