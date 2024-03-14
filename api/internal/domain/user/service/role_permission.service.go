package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(id int32) (output usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(body io.ReadCloser) (err error)
	UpdateRolePermission(id int32, body io.ReadCloser) (err error)
}

type RolePermissionService struct {
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{
	}
}

func (s *RolePermissionService) GetRolePermissions(id int32) (output usecase.GetRolePermissionsOutputDTO, err error) {

	input := usecase.GetRolePermissionsInputDTO{
		RoleID: id,
	}

	output, err = usecase.NewGetRolePermissionsUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(body io.ReadCloser) (err error) {
	s.rolePermissionRepo.Begin()
	input := usecase.CreateRolePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreateRolePermissionUseCase().Execute(input); err != nil {
		s.rolePermissionRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.rolePermissionRepo.Commit()
	return

}

func (s *RolePermissionService) UpdateRolePermission(id int32, body io.ReadCloser) (err error) {
	s.rolePermissionRepo.Begin()
	input := usecase.UpdateRolePermissionInputDTO{
		RoleID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdateRolePermissionUseCase().Execute(input); err != nil {
		s.rolePermissionRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.rolePermissionRepo.Commit()
	slog.Info("Role permission updated")
	return
}
