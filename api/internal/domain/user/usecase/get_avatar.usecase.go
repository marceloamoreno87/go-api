package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetAvatarInputDTO struct {
	ID int32 `json:"id"`
}

type GetAvatarOutputDTO struct {
	ID        int32     `json:"id"`
	SVG       string    `json:"svg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewGetAvatarUseCase() *GetAvatarUseCase {
	return &GetAvatarUseCase{
		repo: repository.NewAvatarRepository(),
	}
}

func (uc *GetAvatarUseCase) Execute(input GetAvatarInputDTO) (output GetAvatarOutputDTO, err error) {
	avatar, err := uc.repo.GetAvatar(input.ID)
	if err != nil {
		return
	}

	output = GetAvatarOutputDTO{
		ID:        avatar.GetID(),
		SVG:       avatar.GetSVG(),
		CreatedAt: avatar.GetCreatedAt(),
		UpdatedAt: avatar.GetUpdatedAt(),
	}
	return
}
