package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateAvatarInputDTO struct {
	SVG string `json:"svg"`
}

type CreateAvatarOutputDTO struct {
	SVG string `json:"svg"`
}

type CreateAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewCreateAvatarUseCase(db config.SQLCInterface) *CreateAvatarUseCase {
	return &CreateAvatarUseCase{
		repo: repository.NewAvatarRepository(db),
	}
}

func (uc *CreateAvatarUseCase) Execute(input CreateAvatarInputDTO) (output CreateAvatarOutputDTO, err error) {
	avatar, err := entity.NewAvatar(input.SVG)
	if err != nil {
		return
	}

	err = uc.repo.CreateAvatar(avatar)

	output = CreateAvatarOutputDTO{
		SVG: avatar.GetSVG(),
	}
	return
}
