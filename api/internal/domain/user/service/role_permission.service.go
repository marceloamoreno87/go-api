package service

import (
	"encoding/json"
	"io"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
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

func (s *RolePermissionService) GetRolePermissions(id int32) (output []usecase.GetRolePermissionsOutputDTO, err error) {
	input := usecase.GetRolePermissionsInputDTO{
		RoleID: id,
	}
	output, err = s.GetRolePermissionsUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(body io.ReadCloser) (output usecase.CreateRolePermissionOutputDTO, err error) {
	input := usecase.CreateRolePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.CreateRolePermissionUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	return
}

func (s *RolePermissionService) UpdateRolePermission(roleId int32, body io.ReadCloser) (output usecase.CreateRolePermissionOutputDTO, err error) {
	inputDelete := usecase.DeleteRolePermissionByRoleIDInputDTO{
		RoleID: roleId,
	}

	if err = json.NewDecoder(body).Decode(&inputDelete); err != nil {
		slog.Info("err", err)
		return
	}
	_, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(inputDelete)
	if err != nil {
		slog.Info("err", err)
		return
	}

	inputCreate := usecase.CreateRolePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&inputCreate); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.CreateRolePermissionUseCase.Execute(inputCreate)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) DeleteRolePermissionByRoleID(id int32) (output usecase.DeleteRolePermissionByRoleIDOutputDTO, err error) {
	input := usecase.DeleteRolePermissionByRoleIDInputDTO{
		RoleID: id,
	}
	output, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role permission deleted")
	return
}
