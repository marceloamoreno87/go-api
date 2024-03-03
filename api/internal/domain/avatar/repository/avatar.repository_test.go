package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/entity"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateAvatar(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewAvatarRepository(db)

	avatar := &entity.Avatar{
		ID:  1,
		SVG: "Test",
	}

	createAvatarSQL := `-- name: CreateAvatar :exec INSERT INTO avatars \( svg \) VALUES \( \$1 \)`
	mock.ExpectBegin()
	mock.ExpectExec(createAvatarSQL).
		WithArgs(avatar.SVG).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = rr.Begin()
	assert.NoError(t, err)
	newAvatar := rr.CreateAvatar(avatar)
	assert.NoError(t, newAvatar)
	err = rr.Commit()
	assert.NoError(t, err)

}

func TestGetAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewAvatarRepository(db)

	avatar := &entity.Avatar{
		ID:  1,
		SVG: "Test",
	}

	rows := sqlmock.NewRows([]string{"id", "svg", "created_at", "updated_at"}).
		AddRow(avatar.ID, avatar.SVG, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetAvatar :one SELECT id, svg, created_at, updated_at FROM avatars WHERE id = \$1 LIMIT 1`).
		WithArgs(avatar.ID).
		WillReturnRows(rows)

	r, err := rr.GetAvatar(avatar.ID)
	assert.NoError(t, err)
	assert.Equal(t, avatar.SVG, r.SVG)
}

func TestDeleteAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewAvatarRepository(db)

	avatar := &entity.Avatar{
		ID:  1,
		SVG: "Test",
	}

	deleteAvatarSQL := `-- name: DeleteAvatar :exec DELETE FROM avatars WHERE id = \$1`
	mock.ExpectBegin()
	mock.ExpectExec(deleteAvatarSQL).
		WithArgs(avatar.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = rr.Begin()
	assert.NoError(t, err)
	deletedAvatar := rr.DeleteAvatar(avatar.ID)
	assert.NoError(t, deletedAvatar)
	err = rr.Commit()
	assert.NoError(t, err)

}

func TestUpdateAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewAvatarRepository(db)

	avatar := &entity.Avatar{
		ID:  1,
		SVG: "Test",
	}

	updateAvatarSQL := `-- name: UpdateAvatar :exec UPDATE avatars SET svg = \$1 WHERE id = \$2`
	mock.ExpectBegin()
	mock.ExpectExec(updateAvatarSQL).
		WithArgs(avatar.SVG, avatar.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = rr.Begin()
	assert.NoError(t, err)
	updatedAvatar := rr.UpdateAvatar(avatar, avatar.ID)
	assert.NoError(t, updatedAvatar)
	err = rr.Commit()
	assert.NoError(t, err)

}

func TestGetAvatars(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := repository.NewAvatarRepository(db)

	avatar := &entity.Avatar{
		ID:  1,
		SVG: "Test",
	}

	rows := sqlmock.NewRows([]string{"id", "svg", "created_at", "updated_at"}).
		AddRow(avatar.ID, avatar.SVG, time.Now(), time.Now())

	mock.ExpectQuery(`-- name: GetAvatars :many SELECT id, svg, created_at, updated_at FROM avatars ORDER BY id ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(int32(10), int32(0)).
		WillReturnRows(rows)

	avatars, err := rr.GetAvatars(10, 0)

	assert.NoError(t, err)

	assert.Equal(t, 1, len(avatars))
	assert.Equal(t, avatar.SVG, avatars[0].SVG)
}
