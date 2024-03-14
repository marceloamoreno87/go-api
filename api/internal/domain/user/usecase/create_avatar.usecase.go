package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateAvatarInputDTO struct {
	SVG string `json:"svg"`
}

type CreateAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewCreateAvatarUseCase() *CreateAvatarUseCase {
	return &CreateAvatarUseCase{
		repo: repository.NewAvatarRepository(),
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
