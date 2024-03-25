package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteAvatarInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteAvatarOutputDTO struct {
	ID int32 `json:"id"`
}

type DeleteAvatarUseCase struct {
	repo repository.Avatarrepository
}

func NewDeleteAvatarUseCase(db config.SQLCInterface) *DeleteAvatarUseCase {
	return &DeleteAvatarUseCase{
		repo: repository.NewAvatarRepository(db),
	}
}

func (uc *DeleteAvatarUseCase) Execute(ctx context.Context, input DeleteAvatarInputDTO) (output DeleteAvatarOutputDTO, err error) {
	err = uc.repo.DeleteAvatar(ctx, input.ID)
	output = DeleteAvatarOutputDTO{
		ID: input.ID,
	}
	return
}
