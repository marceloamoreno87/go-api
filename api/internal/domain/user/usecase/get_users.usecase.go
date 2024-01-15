package usecase

import (
	"context"

	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

type GetUsersUseCase struct {
	UserRepository *db.Queries
}

func NewGetUsersUseCase(userRepository *db.Queries) *GetUsersUseCase {
	return &GetUsersUseCase{
		UserRepository: userRepository,
	}
}

func (uc *GetUsersUseCase) Execute(GetUsersParams db.GetUsersParams) ([]db.User, error) {
	repo, err := uc.UserRepository.GetUsers(context.Background(), GetUsersParams)
	if err != nil {
		return []db.User{}, err
	}
	return repo, nil
}
