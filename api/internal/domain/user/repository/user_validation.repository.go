package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserValidationRepository struct {
	DB config.SQLCInterface
}

func NewUserValidationRepository() *UserValidationRepository {
	return &UserValidationRepository{
		DB: config.NewSqlc(config.DB),
	}
}

func (repo *UserValidationRepository) GetValidationUser(id int32) (output entityInterface.UserValidationInterface, err error) {

	uv, err := repo.DB.GetDbQueries().GetValidationUser(context.Background(), id)
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

func (repo *UserValidationRepository) GetValidationUserByHash(hash string) (output entityInterface.UserValidationInterface, err error) {

	uv, err := repo.DB.GetDbQueries().GetValidationUserByHash(context.Background(), hash)
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

func (repo *UserValidationRepository) CreateValidationUser(userValidation entityInterface.UserValidationInterface) (err error) {

	return repo.DB.GetDbQueries().CreateValidationUser(context.Background(), db.CreateValidationUserParams{
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
	})
}

func (repo *UserValidationRepository) UpdateUserValidationUsed(id int32) (err error) {

	return repo.DB.GetDbQueries().UpdateUserValidationUsed(context.Background(), id)
}
