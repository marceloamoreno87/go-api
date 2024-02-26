package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

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

	input := usecase.CreatePermissionInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewCreatePermissionUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *PermissionService) UpdatePermission(id string, body io.ReadCloser) (err error) {

	input := usecase.UpdatePermissionInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewUpdatePermissionUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *PermissionService) DeletePermission(id string) (err error) {

	input := usecase.DeletePermissionInputDTO{
		ID: helper.StrToInt32(id),
	}

	err = usecase.NewDeletePermissionUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}
