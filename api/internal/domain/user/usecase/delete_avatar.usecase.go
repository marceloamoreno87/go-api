package usecase

import (
	"github.com/marceloamoreno/goapi/config"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteAvatarInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteAvatarOutputDTO struct {
	ID int32 `json:"id"`
}

type DeleteAvatarUseCase struct {
	repo repositoryInterface.AvatarRepositoryInterface
}

func NewDeleteAvatarUseCase(db config.SQLCInterface) *DeleteAvatarUseCase {
	return &DeleteAvatarUseCase{
		repo: repository.NewAvatarRepository(db),
	}
}

func (uc *DeleteAvatarUseCase) Execute(input DeleteAvatarInputDTO) (output DeleteAvatarOutputDTO, err error) {
	err = uc.repo.DeleteAvatar(input.ID)
	output = DeleteAvatarOutputDTO{
		ID: input.ID,
	}
	return
}
