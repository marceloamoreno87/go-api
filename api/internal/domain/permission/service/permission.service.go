package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type PermissionServiceInterface interface {
	CreatePermission(body io.ReadCloser) (err error)
	GetPermission(id string) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(limit string, offset string) (output []usecase.GetPermissionsOutputDTO, err error)
	UpdatePermission(id string, body io.ReadCloser) (err error)
	DeletePermission(id string) (err error)
}

type PermissionService struct {
	repo repository.PermissionRepositoryInterface
}

func NewPermissionService(repo repository.PermissionRepositoryInterface) *PermissionService {
	return &PermissionService{
		repo: repo,
	}
}

func (s *PermissionService) GetPermission(id string) (output usecase.GetPermissionOutputDTO, err error) {

	input := usecase.GetPermissionInputDTO{
		ID: helper.StrToInt32(id),
	}

	output, err = usecase.NewGetPermissionUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *PermissionService) GetPermissions(limit string, offset string) (output []usecase.GetPermissionsOutputDTO, err error) {

	input := usecase.GetPermissionsInputDTO{
		Limit:  helper.StrToInt32(limit),
		Offset: helper.StrToInt32(offset),
	}

	output, err = usecase.NewGetPermissionsUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *PermissionService) CreatePermission(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.CreatePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {

		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreatePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()

	return
}

func (s *PermissionService) UpdatePermission(id string, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdatePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdatePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()

	return
}

func (s *PermissionService) DeletePermission(id string) (err error) {
	s.repo.Begin()

	input := usecase.DeletePermissionInputDTO{
		ID: helper.StrToInt32(id),
	}

	if err = usecase.NewDeletePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()

	return
}
