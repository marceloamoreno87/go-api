package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserActiveInputDTO struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type UpdateUserActiveOutputDTO struct {
}

type UpdateUserActiveUseCase struct {
	repo repository.Userrepository
}

func NewUpdateUserActiveUseCase(db config.SQLCInterface) *UpdateUserActiveUseCase {
	return &UpdateUserActiveUseCase{
		repo: repository.NewUserRepository(db),
	}
}

func (uc *UpdateUserActiveUseCase) Execute(ctx context.Context, input UpdateUserActiveInputDTO) (err error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return
	}
	user.SetID(input.ID)
	user.SetActive(input.Active)

	err = uc.repo.UpdateUserActive(ctx, user.GetID(), user.GetActive())

	return
}
