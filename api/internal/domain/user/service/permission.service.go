package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionService struct {
	DB                                    config.SQLCInterface
	NewGetPermissionUseCase               usecaseInterface.GetPermissionUseCaseInterface
	NewGetPermissionsUseCase              usecaseInterface.GetPermissionsUseCaseInterface
	NewCreatePermissionUseCase            usecaseInterface.CreatePermissionUseCaseInterface
	NewUpdatePermissionUseCase            usecaseInterface.UpdatePermissionUseCaseInterface
	NewDeletePermissionUseCase            usecaseInterface.DeletePermissionUseCaseInterface
	NewGetPermissionByInternalNameUseCase usecaseInterface.GetPermissionByInternalNameUseCaseInterface
}

func NewPermissionService(DB config.SQLCInterface) *PermissionService {
	return &PermissionService{
		DB:                                    DB,
		NewGetPermissionUseCase:               usecase.NewGetPermissionUseCase(DB),
		NewGetPermissionsUseCase:              usecase.NewGetPermissionsUseCase(DB),
		NewCreatePermissionUseCase:            usecase.NewCreatePermissionUseCase(DB),
		NewUpdatePermissionUseCase:            usecase.NewUpdatePermissionUseCase(DB),
		NewDeletePermissionUseCase:            usecase.NewDeletePermissionUseCase(DB),
		NewGetPermissionByInternalNameUseCase: usecase.NewGetPermissionByInternalNameUseCase(DB),
	}
}

func (s *PermissionService) GetPermission(id int32) (output usecase.GetPermissionOutputDTO, err error) {

	input := usecase.GetPermissionInputDTO{
		ID: id,
	}

	output, err = s.NewGetPermissionUseCase.Execute(input)
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

	output, err = s.NewGetPermissionsUseCase.Execute(input)
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

	check, _ := s.NewGetPermissionByInternalNameUseCase.Execute(usecase.GetPermissionByInternalNameInputDTO{
		InternalName: input.InternalName,
	})

	if check.ID != 0 {
		slog.Info("Permission already exists")
		return output, errors.New("permission already exists")
	}

	output, err = s.NewCreatePermissionUseCase.Execute(input)
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
	output, err = s.NewUpdatePermissionUseCase.Execute(input)
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

	output, err = s.NewDeletePermissionUseCase.Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Permission deleted")
	return
}
