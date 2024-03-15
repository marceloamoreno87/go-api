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

-- name: CreateUser :one
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
RETURNING *;

-- name: RegisterUser :one
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
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
  name = $1,
  email = $2,
  password = $3,
  active = $4,
  role_id = $5,
  avatar_id = $6
WHERE id = $7
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users SET
  password = $1
WHERE id = $2
RETURNING *;

-- name: UpdateUserActive :one
UPDATE users SET
  active = $1
WHERE id = $2
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- name: GetAuthByUser :one
SELECT * FROM auth
WHERE user_id = $1 and active is true LIMIT 1;

-- name: CreateAuth :one
INSERT INTO auth (
  user_id,
  token,
  refresh_token,
  expires_in
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: RevokeAuthByUser :one
UPDATE auth SET
  active = false
WHERE user_id = $1
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

-- name: GetPermissionByInternalName :one
SELECT * FROM permissions
WHERE internal_name = $1 LIMIT 1;

-- name: GetPermissions :many
SELECT * FROM permissions
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreatePermission :one
INSERT INTO permissions (
  name,
  internal_name,
  description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdatePermission :one
UPDATE permissions SET
  name = $1,
  internal_name = $2,
  description = $3
WHERE id = $4
RETURNING *;

-- name: DeletePermission :one
DELETE FROM permissions
WHERE id = $1
RETURNING *;

-- name: CreateRolePermission :one
INSERT INTO role_permissions (
  role_id,
  permission_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRolePermission :many
SELECT * FROM role_permissions
WHERE id = $1 LIMIT 1;

-- name: GetRolePermissionsByRole :many
SELECT * FROM role_permissions
INNER JOIN permissions ON role_permissions.permission_id = permissions.id
INNER JOIN roles ON role_permissions.role_id = roles.id
WHERE role_id = $1
ORDER BY permission_id ASC;

-- name: DeleteRolePermission :one
DELETE FROM role_permissions
WHERE role_id = $1
RETURNING *;

-- name: GetAvatar :one
SELECT * FROM avatars
WHERE id = $1 LIMIT 1;

-- name: GetAvatars :many
SELECT * FROM avatars
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CreateAvatar :one
INSERT INTO avatars (
  svg
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateAvatar :one
UPDATE avatars SET
  svg = $1
WHERE id = $2
RETURNING *;

-- name: DeleteAvatar :one
DELETE FROM avatars
WHERE id = $1
RETURNING *;

-- name: GetValidationUser :one
SELECT * FROM users_validation
WHERE user_id = $1 ORDER BY id DESC LIMIT 1;

-- name: GetValidationUserByHash :one
SELECT * FROM users_validation
WHERE hash = $1 and used is false LIMIT 1;

-- name: CreateValidationUser :one
INSERT INTO users_validation (
  user_id,
  hash,
  expires_in
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUserValidationUsed :one
UPDATE users_validation SET
  used = true
WHERE id = $1
RETURNING *;
