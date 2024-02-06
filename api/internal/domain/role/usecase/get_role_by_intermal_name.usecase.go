package usecase

import (
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
)

type GetRoleByInternalNameInputDTO struct {
	InternalName string `json:"internal_name"`
}

type GetRoleByInternalNameOutputDTO struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	InternalName string    `json:"internal_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetRoleByInternalNameUseCase struct {
	RoleRepository repository.RoleRepositoryInterface
}

func NewGetRoleByInternalNameUseCase(roleRepository repository.RoleRepositoryInterface) *GetRoleByInternalNameUseCase {
	return &GetRoleByInternalNameUseCase{
		RoleRepository: roleRepository,
	}
}

func (uc *GetRoleByInternalNameUseCase) Execute(input GetRoleByInternalNameInputDTO) (output GetRoleByInternalNameOutputDTO, err error) {

	role, err := uc.RoleRepository.GetRoleByInternalName(input.InternalName)
	if err != nil {
		return GetRoleByInternalNameOutputDTO{}, err
	}

	output = GetRoleByInternalNameOutputDTO{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
		CreatedAt:    role.CreatedAt,
		UpdatedAt:    role.UpdatedAt,
	}

	return
}
