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

func (repo *UserValidationRepository) GetUserValidationByUserID(id int32) (output entityInterface.UserValidationInterface, err error) {

	uv, err := repo.DB.GetDbQueries().GetUserValidationByUserID(context.Background(), id)
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

func (repo *UserValidationRepository) GetUserValidationByHash(hash string) (output entityInterface.UserValidationInterface, err error) {

	uv, err := repo.DB.GetDbQueries().GetUserValidationByHash(context.Background(), hash)
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

func (repo *UserValidationRepository) CreateUserValidation(userValidation entityInterface.UserValidationInterface) (output entityInterface.UserValidationInterface, err error) {
	newUserValidation, err := repo.DB.GetDbQueries().CreateUserValidation(context.Background(), db.CreateValidationUserParams{
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

func (repo *UserValidationRepository) UpdateUserValidationUsed(id int32) (err error) {

	return repo.DB.GetDbQueries().UpdateUserValidationUsed(context.Background(), id)
}
