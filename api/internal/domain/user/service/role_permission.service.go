package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(id int32) (output usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(body io.ReadCloser) (err error)
	UpdateRolePermission(id int32, body io.ReadCloser) (err error)
	config.SQLCInterface
}

type RolePermissionService struct {
	config.SQLCInterface
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{}
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
	s.Begin()
	input := usecase.CreateRolePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreateRolePermissionUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	return

}

func (s *RolePermissionService) UpdateRolePermission(id int32, body io.ReadCloser) (err error) {
	s.Begin()
	input := usecase.UpdateRolePermissionInputDTO{
		RoleID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdateRolePermissionUseCase().Execute(input); err != nil {
		s.Rollback()
		slog.Info("err", err)
		return
	}
	s.Commit()
	slog.Info("Role permission updated")
	return
}
