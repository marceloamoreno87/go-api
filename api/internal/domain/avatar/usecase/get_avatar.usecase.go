package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
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
	repo repository.AvatarRepositoryInterface
}

func NewGetAvatarUseCase(repo repository.AvatarRepositoryInterface) *GetAvatarUseCase {
	return &GetAvatarUseCase{
		repo: repo,
	}
}

func (uc *GetAvatarUseCase) Execute(input GetAvatarInputDTO) (output GetAvatarOutputDTO, err error) {
	avatar, err := uc.repo.GetAvatar(input.ID)
	if err != nil {
		return
	}

	output = GetAvatarOutputDTO{
		ID:        avatar.ID,
		SVG:       avatar.SVG,
		CreatedAt: avatar.CreatedAt,
		UpdatedAt: avatar.UpdatedAt,
	}
	return
}
