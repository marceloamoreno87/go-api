package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/avatar/repository"
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
	repo repository.AvatarRepositoryInterface
}

func NewGetAvatarsUseCase(repo repository.AvatarRepositoryInterface) *GetAvatarsUseCase {
	return &GetAvatarsUseCase{
		repo: repo,
	}
}

func (uc *GetAvatarsUseCase) Execute(input GetAvatarsInputDTO) (output []GetAvatarsOutputDTO, err error) {
	avatars, err := uc.repo.GetAvatars(input.Limit, input.Offset)
	if err != nil {
		return
	}

	for _, avatar := range avatars {
		output = append(output, GetAvatarsOutputDTO{
			ID:        avatar.ID,
			SVG:       avatar.SVG,
			CreatedAt: avatar.CreatedAt,
			UpdatedAt: avatar.UpdatedAt,
		})
	}

	return
}
