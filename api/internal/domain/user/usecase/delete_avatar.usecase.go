package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteAvatarInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteAvatarOutputDTO struct {
	ID        int32     `json:"id"`
	SVG       string    `json:"svg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewDeleteAvatarUseCase() *DeleteAvatarUseCase {
	return &DeleteAvatarUseCase{
		repo: repository.NewAvatarRepository(),
	}
}

func (uc *DeleteAvatarUseCase) Execute(input DeleteAvatarInputDTO) (err error) {
	avatar, err := uc.repo.GetAvatar(input.ID)
	if err != nil {
		return
	}

	if err = uc.repo.DeleteAvatar(avatar.GetID()); err != nil {
		return
	}

	return
}
