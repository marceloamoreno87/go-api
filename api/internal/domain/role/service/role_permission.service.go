package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(id string) (output usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(body io.ReadCloser) (err error)
	UpdateRolePermission(id string, body io.ReadCloser) (err error)
}

type RolePermissionService struct {
	repo repository.RolePermissionRepositoryInterface
}

func NewRolePermissionService(repo repository.RolePermissionRepositoryInterface) *RolePermissionService {
	return &RolePermissionService{
		repo: repo,
	}
}

func (s *RolePermissionService) GetRolePermissions(id string) (output usecase.GetRolePermissionsOutputDTO, err error) {

	input := usecase.GetRolePermissionsInputDTO{
		RoleID: helper.StrToInt32(id),
	}

	output, err = usecase.NewGetRolePermissionsUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.CreateRolePermissionInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewCreateRolePermissionUseCase(s.repo).Execute(input)
	if err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	return

}

func (s *RolePermissionService) UpdateRolePermission(id string, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdateRolePermissionInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewUpdateRolePermissionUseCase(s.repo).Execute(input)
	if err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	return
}
