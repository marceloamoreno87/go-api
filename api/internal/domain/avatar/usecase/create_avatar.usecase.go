package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/avatar/entity"
	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
)

type CreateAvatarInputDTO struct {
	SVG string `json:"svg"`
}

type CreateAvatarUseCase struct {
	repo repository.AvatarRepositoryInterface
}

func NewCreateAvatarUseCase(repo repository.AvatarRepositoryInterface) *CreateAvatarUseCase {
	return &CreateAvatarUseCase{
		repo: repo,
	}
}

func (uc *CreateAvatarUseCase) Execute(input CreateAvatarInputDTO) (err error) {
	avatar, err := entity.NewAvatar(input.SVG)
	if err != nil {
		return
	}

	if err = uc.repo.CreateAvatar(avatar); err != nil {
		return
	}

	return
}
