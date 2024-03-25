package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserPasswordInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordOutputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPasswordUseCase struct {
	repo repository.Userrepository
}

func NewUpdateUserPasswordUseCase(db config.SQLCInterface) *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		repo: repository.NewUserRepository(db),
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(ctx context.Context, input UpdateUserPasswordInputDTO) (output UpdateUserPasswordOutputDTO, err error) {

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return
	}
	user.SetID(input.ID)
	user.HashPassword()

	err = uc.repo.UpdateUserPassword(ctx, user.GetID(), user.GetPassword())
	output = UpdateUserPasswordOutputDTO{
		ID:       input.ID,
		Password: input.Password,
	}

	return
}
