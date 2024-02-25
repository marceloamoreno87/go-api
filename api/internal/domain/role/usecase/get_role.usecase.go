package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type GetRoleInputDTO struct {
	ID int32 `json:"id"`
}

type GetRoleOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRoleUseCase struct {
	repo repository.RoleRepositoryInterface
}

func NewGetRoleUseCase(repo repository.RoleRepositoryInterface) *GetRoleUseCase {
	return &GetRoleUseCase{
		repo: repo,
	}
}

func (uc *GetRoleUseCase) Execute(input GetRoleInputDTO) (output GetRoleOutputDTO, err error) {

	role, err := uc.repo.GetRole(input.ID)
	if err != nil {
		return
	}

	output = GetRoleOutputDTO{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
		CreatedAt:    role.CreatedAt,
		UpdatedAt:    role.UpdatedAt,
	}

	return
}
