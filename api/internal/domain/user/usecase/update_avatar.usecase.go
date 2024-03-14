package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateAvatarInputDTO struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type UpdateAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewUpdateAvatarUseCase() *UpdateAvatarUseCase {
	return &UpdateAvatarUseCase{
		repo: repository.NewAvatarRepository(),
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
