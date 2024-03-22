package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateAuthRevokeInputDTO struct {
	UserID int32 `json:"user_id"`
}

type UpdateAuthRevokeOutputDTO struct {
	UserID int32 `json:"user_id"`
}

type UpdateAuthRevokeUseCase struct {
	repo repositoryInterface.AuthRepositoryInterface
}

func NewUpdateAuthRevokeUseCase(db config.SQLCInterface) *UpdateAuthRevokeUseCase {
	return &UpdateAuthRevokeUseCase{
		repo: repository.NewAuthRepository(db),
	}
}

func (uc *UpdateAuthRevokeUseCase) Execute(ctx context.Context, input UpdateAuthRevokeInputDTO) (output UpdateAuthRevokeOutputDTO, err error) {
	err = uc.repo.UpdateAuthRevokeByUserID(ctx, input.UserID)
	if err != nil {
		return
	}

	output = UpdateAuthRevokeOutputDTO{
		UserID: input.UserID,
	}

	return
}
