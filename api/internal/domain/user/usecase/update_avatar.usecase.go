package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateAvatarInputDTO struct {
	ID  int32  `json:"id"`
	SVG string `json:"svg"`
}

type UpdateAvatarOutputDTO struct {
	ID        int32     `json:"id"`
	SVG       string    `json:"svg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewUpdateAvatarUseCase() *UpdateAvatarUseCase {
	return &UpdateAvatarUseCase{
		repo: repository.NewAvatarRepository(),
	}
}

func (uc *UpdateAvatarUseCase) Execute(input UpdateAvatarInputDTO) (output UpdateAvatarOutputDTO, err error) {
	avatar, err := entity.NewAvatar(input.SVG)
	if err != nil {
		return
	}

	a, err := uc.repo.UpdateAvatar(avatar, input.ID)
	if err != nil {
		return
	}
	output = UpdateAvatarOutputDTO{
		ID:        a.GetID(),
		SVG:       a.GetSVG(),
		CreatedAt: a.GetCreatedAt(),
		UpdatedAt: a.GetUpdatedAt(),
	}
	return
}
