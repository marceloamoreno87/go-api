package service

import (
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RequestCreateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type RequestGetRolePermissionInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type RequestDeleteRolePermissionByRoleIDInputDTO struct {
	RoleID int32 `json:"role_id"`
}

type RequestUpdateRolePermissionInputDTO struct {
	RoleID        int32   `json:"role_id"`
	PermissionIDs []int32 `json:"permission_ids"`
}

type RolePermissionService struct {
	GetRolePermissionsUseCase           usecaseInterface.GetRolePermissionsUseCaseInterface
	CreateRolePermissionUseCase         usecaseInterface.CreateRolePermissionUseCaseInterface
	DeleteRolePermissionByRoleIDUseCase usecaseInterface.DeleteRolePermissionByRoleIDUseCaseInterface
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{
		GetRolePermissionsUseCase:           usecase.NewGetRolePermissionsUseCase(),
		CreateRolePermissionUseCase:         usecase.NewCreateRolePermissionUseCase(),
		DeleteRolePermissionByRoleIDUseCase: usecase.NewDeleteRolePermissionByRoleIDUseCase(),
	}
}

func (s *RolePermissionService) GetRolePermissions(input RequestGetRolePermissionInputDTO) (output []usecase.GetRolePermissionsOutputDTO, err error) {
	output, err = s.GetRolePermissionsUseCase.Execute(usecase.GetRolePermissionsInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(input RequestCreateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error) {
	output, err = s.CreateRolePermissionUseCase.Execute(usecase.CreateRolePermissionInputDTO{
		RoleID:        input.RoleID,
		PermissionIDs: input.PermissionIDs,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	return
}

func (s *RolePermissionService) UpdateRolePermission(input RequestUpdateRolePermissionInputDTO) (output usecase.CreateRolePermissionOutputDTO, err error) {
	_, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.CreateRolePermissionUseCase.Execute(usecase.CreateRolePermissionInputDTO{
		RoleID:        input.RoleID,
		PermissionIDs: input.PermissionIDs,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) DeleteRolePermissionByRoleID(input RequestDeleteRolePermissionByRoleIDInputDTO) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error) {
	output, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role permission deleted")
	return
}
