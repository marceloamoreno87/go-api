package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleService struct {
	NewGetRoleUseCase               usecaseInterface.GetRoleUseCaseInterface
	NewGetRolesUseCase              usecaseInterface.GetRolesUseCaseInterface
	NewCreateRoleUseCase            usecaseInterface.CreateRoleUseCaseInterface
	NewUpdateRoleUseCase            usecaseInterface.UpdateRoleUseCaseInterface
	NewDeleteRoleUseCase            usecaseInterface.DeleteRoleUseCaseInterface
	NewGetRoleByInternalNameUseCase usecaseInterface.NewGetRoleByInternalNameUseCaseInterface
}

func NewRoleService() *RoleService {
	return &RoleService{
		NewGetRoleUseCase:               usecase.NewGetRoleUseCase(),
		NewGetRolesUseCase:              usecase.NewGetRolesUseCase(),
		NewCreateRoleUseCase:            usecase.NewCreateRoleUseCase(),
		NewUpdateRoleUseCase:            usecase.NewUpdateRoleUseCase(),
		NewDeleteRoleUseCase:            usecase.NewDeleteRoleUseCase(),
		NewGetRoleByInternalNameUseCase: usecase.NewGetRoleByInternalNameUseCase(),
	}
}

func (s *RoleService) GetRole(id int32) (output usecase.GetRoleOutputDTO, err error) {
	input := usecase.GetRoleInputDTO{
		ID: id,
	}

	output, err = s.NewGetRoleUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role found")
	return
}

func (s *RoleService) GetRoles(limit int32, offset int32) (output []usecase.GetRolesOutputDTO, err error) {

	input := usecase.GetRolesInputDTO{
		Limit:  limit,
		Offset: offset,
	}

	output, err = s.NewGetRolesUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Roles found")
	return
}

func (s *RoleService) CreateRole(body io.ReadCloser) (output usecase.CreateRoleOutputDTO, err error) {
	input := usecase.CreateRoleInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := s.NewGetRoleByInternalNameUseCase.Execute(usecase.GetRoleByInternalNameInputDTO{InternalName: input.InternalName})
	if check.ID != 0 {
		slog.Info("role already exists")
		return output, errors.New("role already exists")
	}

	output, err = s.NewCreateRoleUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role created")
	return
}

func (s *RoleService) UpdateRole(id int32, body io.ReadCloser) (output usecase.UpdateRoleOutputDTO, err error) {
	input := usecase.UpdateRoleInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = s.NewUpdateRoleUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role updated")
	return
}

func (s *RoleService) DeleteRole(id int32) (output usecase.DeleteRoleOutputDTO, err error) {
	input := usecase.DeleteRoleInputDTO{
		ID: id,
	}

	output, err = s.NewDeleteRoleUseCase.Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role deleted")
	return
}
