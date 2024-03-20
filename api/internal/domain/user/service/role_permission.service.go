package service

import (
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

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

func (s *RolePermissionService) GetRolePermissions(input request.RequestGetRolePermission) (output []usecase.GetRolePermissionsOutputDTO, err error) {
	output, err = s.GetRolePermissionsUseCase.Execute(usecase.GetRolePermissionsInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(input request.RequestCreateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error) {
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

func (s *RolePermissionService) UpdateRolePermission(input request.RequestUpdateRolePermission) (output usecase.CreateRolePermissionOutputDTO, err error) {
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

func (s *RolePermissionService) DeleteRolePermissionByRoleID(input request.RequestDeleteRolePermissionByRoleID) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error) {
	output, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role permission deleted")
	return
}
