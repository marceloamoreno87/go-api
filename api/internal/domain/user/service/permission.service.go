package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionService struct {
	DB config.SQLCInterface
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		DB: config.Sqcl,
	}
}

func (s *PermissionService) GetPermission(id int32) (output usecase.GetPermissionOutputDTO, err error) {

	input := usecase.GetPermissionInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetPermissionUseCase().Execute(input)
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

	output, err = usecase.NewGetPermissionsUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permissions found")
	return
}

func (s *PermissionService) CreatePermission(body io.ReadCloser) (output usecase.CreatePermissionOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.CreatePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := usecase.NewGetPermissionByInternalNameUseCase().Execute(usecase.GetPermissionByInternalNameInputDTO{
		InternalName: input.InternalName,
	})

	if check.ID != 0 {
		slog.Info("Permission already exists")
		return output, errors.New("permission already exists")
	}

	output, err = usecase.NewCreatePermissionUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Permission created")
	return
}

func (s *PermissionService) UpdatePermission(id int32, body io.ReadCloser) (output usecase.UpdatePermissionOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.UpdatePermissionInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}
	output, err = usecase.NewUpdatePermissionUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Permission updated")
	return
}

func (s *PermissionService) DeletePermission(id int32) (output usecase.DeletePermissionOutputDTO, err error) {
	s.DB.Begin()

	input := usecase.DeletePermissionInputDTO{
		ID: id,
	}

	output, err = usecase.NewDeletePermissionUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Permission deleted")
	return
}
