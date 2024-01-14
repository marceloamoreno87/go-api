// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, bio
) VALUES (
  $1, $2
)
RETURNING id, name, bio
`

type CreateUserParams struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Bio)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, bio FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, bio FROM users
ORDER BY name
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
  set name = $2,
  bio = $3
WHERE id = $1
`

type UpdateUserParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.ID, arg.Name, arg.Bio)
	return err
}