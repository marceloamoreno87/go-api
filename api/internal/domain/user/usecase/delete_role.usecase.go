package usecase

import (
	"time"

	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type DeleteRoleInputDTO struct {
	ID int32 `json:"id"`
}

type DeleteRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DeleteRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewDeleteRoleUseCase() *DeleteRoleUseCase {
	return &DeleteRoleUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *DeleteRoleUseCase) Execute(input DeleteRoleInputDTO) (output DeleteRoleOutputDTO, err error) {
	role, err := uc.repo.GetRole(input.ID)
	if err != nil {
		return
	}

	r, err := uc.repo.DeleteRole(role.GetID())
	if err != nil {
		return
	}
	output = DeleteRoleOutputDTO{
		ID:           r.GetID(),
		Name:         r.GetName(),
		InternalName: r.GetInternalName(),
		Description:  r.GetDescription(),
		CreatedAt:    r.GetCreatedAt(),
		UpdatedAt:    r.GetUpdatedAt(),
	}
	return
}
