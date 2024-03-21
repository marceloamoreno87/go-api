package usecase

import (
	"context"
	"time"

	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type GetAvatarsInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAvatarsOutputDTO struct {
	ID        int32     `json:"id"`
	SVG       string    `json:"svg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAvatarsUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewGetAvatarsUseCase(db config.SQLCInterface) *GetAvatarsUseCase {
	return &GetAvatarsUseCase{
		repo: repository.NewAvatarRepository(db),
	}
}

func (uc *GetAvatarsUseCase) Execute(ctx context.Context, input GetAvatarsInputDTO) (output []GetAvatarsOutputDTO, err error) {
	avatars, err := uc.repo.GetAvatars(ctx, input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, avatar := range avatars {
		output = append(output, GetAvatarsOutputDTO{
			ID:        avatar.GetID(),
			SVG:       avatar.GetSVG(),
			CreatedAt: avatar.GetCreatedAt(),
			UpdatedAt: avatar.GetUpdatedAt(),
		})
	}

	return
}
