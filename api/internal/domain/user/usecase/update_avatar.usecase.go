package usecase

import (
	"context"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateAvatarInputDTO struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type UpdateAvatarOutputDTO struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type UpdateAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewUpdateAvatarUseCase(db config.SQLCInterface) *UpdateAvatarUseCase {
	return &UpdateAvatarUseCase{
		repo: repository.NewAvatarRepository(db),
	}
}

func (uc *UpdateAvatarUseCase) Execute(ctx context.Context, input UpdateAvatarInputDTO) (output UpdateAvatarOutputDTO, err error) {
	avatar, err := entity.NewAvatar(input.SVG)
	if err != nil {
		return
	}

	err = uc.repo.UpdateAvatar(ctx, avatar, input.ID)
	output = UpdateAvatarOutputDTO{
		ID:  avatar.GetID(),
		SVG: avatar.GetSVG(),
	}
	return
}
