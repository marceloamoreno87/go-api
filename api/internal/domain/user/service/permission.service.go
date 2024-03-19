package service

import (
	"errors"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RequestGetPermissionInputDTO struct {
	ID int32 `json:"id"`
}

type RequestGetPermissionsInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type RequestCreatePermissionInputDTO struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type RequestUpdatePermissionInputDTO struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
	Description  string `json:"description"`
}

type RequestDeletePermissionInputDTO struct {
	ID int32 `json:"id"`
}

type PermissionService struct {
	NewGetPermissionUseCase               usecaseInterface.GetPermissionUseCaseInterface
	NewGetPermissionsUseCase              usecaseInterface.GetPermissionsUseCaseInterface
	NewCreatePermissionUseCase            usecaseInterface.CreatePermissionUseCaseInterface
	NewUpdatePermissionUseCase            usecaseInterface.UpdatePermissionUseCaseInterface
	NewDeletePermissionUseCase            usecaseInterface.DeletePermissionUseCaseInterface
	NewGetPermissionByInternalNameUseCase usecaseInterface.GetPermissionByInternalNameUseCaseInterface
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		NewGetPermissionUseCase:               usecase.NewGetPermissionUseCase(),
		NewGetPermissionsUseCase:              usecase.NewGetPermissionsUseCase(),
		NewCreatePermissionUseCase:            usecase.NewCreatePermissionUseCase(),
		NewUpdatePermissionUseCase:            usecase.NewUpdatePermissionUseCase(),
		NewDeletePermissionUseCase:            usecase.NewDeletePermissionUseCase(),
		NewGetPermissionByInternalNameUseCase: usecase.NewGetPermissionByInternalNameUseCase(),
	}
}

func (s *PermissionService) GetPermission(input RequestGetPermissionInputDTO) (output usecase.GetPermissionOutputDTO, err error) {
	output, err = s.NewGetPermissionUseCase.Execute(usecase.GetPermissionInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permission found")
	return
}

func (s *PermissionService) GetPermissions(input RequestGetPermissionsInputDTO) (output []usecase.GetPermissionsOutputDTO, err error) {
	output, err = s.NewGetPermissionsUseCase.Execute(usecase.GetPermissionsInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permissions found")
	return
}

func (s *PermissionService) CreatePermission(input RequestCreatePermissionInputDTO) (output usecase.CreatePermissionOutputDTO, err error) {
	check, _ := s.NewGetPermissionByInternalNameUseCase.Execute(usecase.GetPermissionByInternalNameInputDTO{
		InternalName: input.InternalName,
	})

	if check.ID != 0 {
		slog.Info("Permission already exists")
		return output, errors.New("permission already exists")
	}

	output, err = s.NewCreatePermissionUseCase.Execute(usecase.CreatePermissionInputDTO{
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permission created")
	return
}

func (s *PermissionService) UpdatePermission(input RequestUpdatePermissionInputDTO) (output usecase.UpdatePermissionOutputDTO, err error) {
	output, err = s.NewUpdatePermissionUseCase.Execute(usecase.UpdatePermissionInputDTO{
		ID:           input.ID,
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permission updated")
	return
}

func (s *PermissionService) DeletePermission(input RequestDeletePermissionInputDTO) (output usecase.DeletePermissionOutputDTO, err error) {
	output, err = s.NewDeletePermissionUseCase.Execute(usecase.DeletePermissionInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permission deleted")
	return
}
