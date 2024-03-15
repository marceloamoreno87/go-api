package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionService struct {
	DB config.SQLCInterface
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{
		DB: config.Sqcl,
	}
}

func (s *RolePermissionService) GetRolePermissions(id int32) (output []usecase.GetRolePermissionsOutputDTO, err error) {

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

func (s *RolePermissionService) CreateRolePermission(body io.ReadCloser) (output usecase.CreateRolePermissionOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.CreateRolePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = usecase.NewCreateRolePermissionUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	return

}

func (s *RolePermissionService) DeleteRolePermission(id int32, body io.ReadCloser) (output usecase.DeleteRolePermissionOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.DeleteRolePermissionInputDTO{
		RoleID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = usecase.NewDeleteRolePermissionUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Role permission deleted")
	return
}
