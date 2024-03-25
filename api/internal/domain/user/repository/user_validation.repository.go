package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserValidationrepository interface {
	CreateUserValidation(ctx context.Context, userValidation *entity.UserValidation) (output *entity.UserValidation, err error)
	GetUserValidationByUserID(ctx context.Context, id int32) (output *entity.UserValidation, err error)
	GetUserValidationByHash(ctx context.Context, hash string) (output *entity.UserValidation, err error)
	UpdateUserValidationUsed(ctx context.Context, id int32) (err error)
}

type UserValidationRepository struct {
	db config.SQLCInterface
}

func NewUserValidationRepository(db config.SQLCInterface) *UserValidationRepository {
	return &UserValidationRepository{
		db: db,
	}
}

func (repo *UserValidationRepository) GetUserValidationByUserID(ctx context.Context, id int32) (output *entity.UserValidation, err error) {
	uv, err := repo.db.GetDbQueries().GetUserValidationByUserID(ctx, id)
	if err != nil {
		return
	}
	output = &entity.UserValidation{
		ID:        uv.ID,
		UserID:    uv.UserID,
		Hash:      uv.Hash,
		ExpiresIn: uv.ExpiresIn,
		Used:      uv.Used,
		CreatedAt: uv.CreatedAt,
		UpdatedAt: uv.UpdatedAt,
	}

	return
}

func (repo *UserValidationRepository) GetUserValidationByHash(ctx context.Context, hash string) (output *entity.UserValidation, err error) {
	uv, err := repo.db.GetDbQueries().GetUserValidationByHash(ctx, hash)
	if err != nil {
		return
	}
	output = &entity.UserValidation{
		ID:        uv.ID,
		UserID:    uv.UserID,
		Hash:      uv.Hash,
		ExpiresIn: uv.ExpiresIn,
		Used:      uv.Used,
		CreatedAt: uv.CreatedAt,
		UpdatedAt: uv.UpdatedAt,
	}
	return
}

func (repo *UserValidationRepository) CreateUserValidation(ctx context.Context, userValidation *entity.UserValidation) (output *entity.UserValidation, err error) {
	newUserValidation, err := repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreateUserValidation(ctx, db.CreateUserValidationParams{
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
	})
	if err != nil {
		return
	}
	output = &entity.UserValidation{
		ID:        newUserValidation.ID,
		UserID:    newUserValidation.UserID,
		Hash:      newUserValidation.Hash,
		ExpiresIn: newUserValidation.ExpiresIn,
		Used:      newUserValidation.Used,
		CreatedAt: newUserValidation.CreatedAt,
		UpdatedAt: newUserValidation.UpdatedAt,
	}
	return

}

func (repo *UserValidationRepository) UpdateUserValidationUsed(ctx context.Context, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).UpdateUserValidationUsed(ctx, id)
}
