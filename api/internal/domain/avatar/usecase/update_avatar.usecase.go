package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/avatar/entity"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
)

type UpdateAvatarInputDTO struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type UpdateAvatarUseCase struct {
	repo repository.AvatarRepositoryInterface
}

func NewUpdateAvatarUseCase(repo repository.AvatarRepositoryInterface) *UpdateAvatarUseCase {
	return &UpdateAvatarUseCase{
		repo: repo,
	}
}

func (uc *UpdateAvatarUseCase) Execute(input UpdateAvatarInputDTO) (err error) {
	avatar, err := entity.NewAvatar(input.SVG)
	if err != nil {
		return
	}

	if err = uc.repo.UpdateAvatar(avatar, input.ID); err != nil {
		return
	}

	return
}
