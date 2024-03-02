package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
)

type RolePermissionServiceInterface interface {
	GetRolePermissions(id int32) (output usecase.GetRolePermissionsOutputDTO, err error)
	CreateRolePermission(body io.ReadCloser) (err error)
	UpdateRolePermission(id int32, body io.ReadCloser) (err error)
}

type RolePermissionService struct {
	repo repository.RolePermissionRepositoryInterface
}

func NewRolePermissionService(repo repository.RolePermissionRepositoryInterface) *RolePermissionService {
	return &RolePermissionService{
		repo: repo,
	}
}

func (s *RolePermissionService) GetRolePermissions(id int32) (output usecase.GetRolePermissionsOutputDTO, err error) {

	input := usecase.GetRolePermissionsInputDTO{
		RoleID: id,
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
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreateRolePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	return

}

func (s *RolePermissionService) UpdateRolePermission(id int32, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdateRolePermissionInputDTO{
		RoleID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdateRolePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Role permission updated")
	return
}
