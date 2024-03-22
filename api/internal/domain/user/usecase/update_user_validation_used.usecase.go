package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateUserValidationUsedInputDTO struct {
	UserID int32 `json:"user_id"`
}

type UpdateUserValidationUsedUseCase struct {
	repo repositoryInterface.UserValidationRepositoryInterface
}

func NewUpdateUserValidationUsedUseCase(db config.SQLCInterface) *UpdateUserValidationUsedUseCase {
	return &UpdateUserValidationUsedUseCase{
		repo: repository.NewUserValidationRepository(db),
	}
}

func (uc *UpdateUserValidationUsedUseCase) Execute(ctx context.Context, input UpdateUserValidationUsedInputDTO) (err error) {
	return uc.repo.UpdateUserValidationUsed(ctx, input.UserID)
}
