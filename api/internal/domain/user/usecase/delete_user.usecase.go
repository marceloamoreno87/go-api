package usecase

import (
	"context"

	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

type DeleteUserUseCase struct {
	UserRepository *db.Queries
}

func NewDeleteUserUseCase(userRepository *db.Queries) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(id int64) (err error) {
	_, err = uc.UserRepository.GetUser(context.Background(), id)
	if err != nil {
		return err
	}

	err = uc.UserRepository.DeleteUser(context.Background(), id)
	if err != nil {
		return
	}
	return
}
