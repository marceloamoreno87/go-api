package usecase

import "github.com/marceloamoreno/goapi/internal/domain/role/repository"

type DeleteRoleInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleUseCase struct {
	repo repository.RoleRepositoryInterface
}

func NewDeleteRoleUseCase(repo repository.RoleRepositoryInterface) *DeleteRoleUseCase {
	return &DeleteRoleUseCase{
		repo: repo,
	}
}

func (uc *DeleteRoleUseCase) Execute(input DeleteRoleInputDTO) (err error) {

	role, err := uc.repo.GetRole(input.ID)
	if err != nil {
		return
	}

	err = uc.repo.DeleteRole(role.GetID())
	if err != nil {
		return
	}

	return
}
