package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/permission/repository"
	"github.com/marceloamoreno/goapi/internal/domain/permission/usecase"
)

type PermissionServiceInterface interface {
	CreatePermission(body io.ReadCloser) (err error)
	GetPermission(id int32) (output usecase.GetPermissionOutputDTO, err error)
	GetPermissions(limit int32, offset int32) (output []usecase.GetPermissionsOutputDTO, err error)
	UpdatePermission(id int32, body io.ReadCloser) (err error)
	DeletePermission(id int32) (err error)
}

type PermissionService struct {
	repo repository.PermissionRepositoryInterface
}

func NewPermissionService(repo repository.PermissionRepositoryInterface) *PermissionService {
	return &PermissionService{
		repo: repo,
	}
}

func (s *PermissionService) GetPermission(id int32) (output usecase.GetPermissionOutputDTO, err error) {

	input := usecase.GetPermissionInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetPermissionUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permission found")
	return
}

func (s *PermissionService) GetPermissions(limit int32, offset int32) (output []usecase.GetPermissionsOutputDTO, err error) {

	input := usecase.GetPermissionsInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	output, err = usecase.NewGetPermissionsUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permissions found")
	return
}

func (s *PermissionService) CreatePermission(body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.CreatePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, _ := usecase.NewGetPermissionByInternalNameUseCase(s.repo).Execute(usecase.GetPermissionByInternalNameInputDTO{
		InternalName: input.InternalName,
	})

	if output.ID != 0 {
		slog.Info("Permission already exists")
		return errors.New("permission already exists")
	}

	if err = usecase.NewCreatePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Permission created")
	return
}

func (s *PermissionService) UpdatePermission(id int32, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdatePermissionInputDTO{
		ID: id,
	}
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
	slog.Info("Permission updated")
	return
}

func (s *PermissionService) DeletePermission(id int32) (err error) {
	s.repo.Begin()

	input := usecase.DeletePermissionInputDTO{
		ID: id,
	}

	if err = usecase.NewDeletePermissionUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	slog.Info("Permission deleted")
	return
}
