package service

import (
	"encoding/json"
	"io"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionService struct {
	NewGetRolePermissionsUseCase   usecaseInterface.GetRolePermissionsUseCaseInterface
	NewCreateRolePermissionUseCase usecaseInterface.CreateRolePermissionUseCaseInterface
	NewDeleteRolePermissionUseCase usecaseInterface.DeleteRolePermissionUseCaseInterface
}

func NewRolePermissionService() *RolePermissionService {
	return &RolePermissionService{
		NewGetRolePermissionsUseCase:   usecase.NewGetRolePermissionsUseCase(),
		NewCreateRolePermissionUseCase: usecase.NewCreateRolePermissionUseCase(),
		NewDeleteRolePermissionUseCase: usecase.NewDeleteRolePermissionUseCase(),
	}
}

func (s *RolePermissionService) GetRolePermissions(id int32) (output []usecase.GetRolePermissionsOutputDTO, err error) {

	input := usecase.GetRolePermissionsInputDTO{
		RoleID: id,
	}

	output, err = s.NewGetRolePermissionsUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RolePermissionService) CreateRolePermission(body io.ReadCloser) (output usecase.CreateRolePermissionOutputDTO, err error) {
	input := usecase.CreateRolePermissionInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.NewCreateRolePermissionUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	return

}

func (s *RolePermissionService) DeleteRolePermission(id int32, body io.ReadCloser) (output usecase.DeleteRolePermissionOutputDTO, err error) {
	input := usecase.DeleteRolePermissionInputDTO{
		RoleID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.NewDeleteRolePermissionUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role permission deleted")
	return
}
