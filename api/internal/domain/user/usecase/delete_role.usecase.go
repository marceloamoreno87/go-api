package usecase

import (
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteRoleInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewDeleteRoleUseCase() *DeleteRoleUseCase {
	return &DeleteRoleUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *DeleteRoleUseCase) Execute(input DeleteRoleInputDTO) (err error) {

	role, err := uc.repo.GetRole(input.ID)
	if err != nil {
		return
	}

	if err = uc.repo.DeleteRole(role.GetID()); err != nil {
		return
	}

	return
}
