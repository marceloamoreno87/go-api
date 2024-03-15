package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	repositoryInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
)

type UpdateRoleInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type UpdateRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateRoleUseCase struct {
	repo repositoryInterface.RoleRepositoryInterface
}

func NewUpdateRoleUseCase() *UpdateRoleUseCase {
	return &UpdateRoleUseCase{
		repo: repository.NewRoleRepository(),
	}
}

func (uc *UpdateRoleUseCase) Execute(input UpdateRoleInputDTO) (output UpdateRoleOutputDTO, err error) {
	role, err := entity.NewRole(input.Name, input.InternalName, input.Description)
	if err != nil {
		return
	}
	r, err := uc.repo.UpdateRole(role, input.ID)
	if err != nil {
		return
	}

	output = UpdateRoleOutputDTO{
		ID:           r.GetID(),
		Name:         r.GetName(),
		InternalName: r.GetInternalName(),
		Description:  r.GetDescription(),
		CreatedAt:    r.GetCreatedAt(),
		UpdatedAt:    r.GetUpdatedAt(),
	}
	return
}
