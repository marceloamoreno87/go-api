package repository

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type UserValidationRepository struct {
	config.SQLCInterface
}

func NewUserValidationRepository() *UserValidationRepository {
	return &UserValidationRepository{}
}

func (repo *UserValidationRepository) GetValidationUser(id int32) (userValidation entityInterface.UserValidationInterface, err error) {
	uv, err := repo.GetDbQueries().GetValidationUser(context.Background(), id)
	if err != nil {
		return
	}
	userValidation = &entity.UserValidation{
		ID:        uv.ID,
		UserID:    uv.UserID,
		Hash:      uv.Hash,
		Used:      uv.Used,
		ExpiresIn: uv.ExpiresIn,
		CreatedAt: uv.CreatedAt,
		UpdatedAt: uv.UpdatedAt,
	}
	return
}

func (repo *UserValidationRepository) GetValidationUserByHash(hash string) (userValidation entityInterface.UserValidationInterface, err error) {
	uv, err := repo.GetDbQueries().GetValidationUserByHash(context.Background(), hash)
	if err != nil {
		return
	}
	userValidation = &entity.UserValidation{
		ID:        uv.ID,
		UserID:    uv.UserID,
		Hash:      uv.Hash,
		Used:      uv.Used,
		ExpiresIn: uv.ExpiresIn,
		CreatedAt: uv.CreatedAt,
		UpdatedAt: uv.UpdatedAt,
	}
	return
}

func (repo *UserValidationRepository) CreateValidationUser(userValidation entityInterface.UserValidationInterface) (err error) {
	err = repo.GetDbQueries().WithTx(repo.GetTx()).CreateValidationUser(context.Background(), db.CreateValidationUserParams{
		UserID:    userValidation.GetUserID(),
		Hash:      userValidation.GetHash(),
		ExpiresIn: userValidation.GetExpiresIn(),
	})
	return
}

func (repo *UserValidationRepository) SetUserValidationUsed(id int32) (err error) {
	return repo.GetDbQueries().WithTx(repo.GetTx()).SetUserValidationUsed(context.Background(), id)
}
