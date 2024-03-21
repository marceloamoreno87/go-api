package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
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

func NewGetAvatarUseCase(db config.SQLCInterface) *GetAvatarUseCase {
	return &GetAvatarUseCase{
		repo: repository.NewAvatarRepository(db),
	}
}

func (uc *GetAvatarUseCase) Execute(ctx context.Context, input GetAvatarInputDTO) (output GetAvatarOutputDTO, err error) {
	avatar, err := uc.repo.GetAvatar(ctx, input.ID)
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
