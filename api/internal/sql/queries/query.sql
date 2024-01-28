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
  password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
  name = $1,
  email = $2,
  password = $3
WHERE id = $4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
