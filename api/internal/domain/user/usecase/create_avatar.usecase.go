package usecase

import (
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type CreateAvatarInputDTO struct {
	SVG string `json:"svg"`
}

type CreateAvatarOutputDTO struct {
	ID        int32  `json:"id"`
	SVG       string `json:"svg"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewCreateAvatarUseCase() *CreateAvatarUseCase {
	return &CreateAvatarUseCase{
		repo: repository.NewAvatarRepository(),
	}
}

func (uc *CreateAvatarUseCase) Execute(input CreateAvatarInputDTO) (output CreateAvatarOutputDTO, err error) {
	avatar, err := entity.NewAvatar(input.SVG)
	if err != nil {
		return
	}

	a, err := uc.repo.CreateAvatar(avatar)
	if err != nil {
		return
	}

	output = CreateAvatarOutputDTO{
		ID:        a.GetID(),
		SVG:       a.GetSVG(),
		CreatedAt: a.GetCreatedAt().String(),
		UpdatedAt: a.GetUpdatedAt().String(),
	}
	return
}
