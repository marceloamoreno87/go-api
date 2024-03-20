package service

import (
	"errors"
	"log/slog"

	usecaseInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/usecase"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
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

func (s *RoleService) GetRole(input request.RequestGetRole) (output usecase.GetRoleOutputDTO, err error) {
	output, err = s.NewGetRoleUseCase.Execute(usecase.GetRoleInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role found")
	return
}

func (s *RoleService) GetRoles(input request.RequestGetRoles) (output []usecase.GetRolesOutputDTO, err error) {
	output, err = s.NewGetRolesUseCase.Execute(usecase.GetRolesInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Roles found")
	return
}

func (s *RoleService) CreateRole(input request.RequestCreateRole) (output usecase.CreateRoleOutputDTO, err error) {
	check, _ := s.NewGetRoleByInternalNameUseCase.Execute(usecase.GetRoleByInternalNameInputDTO{InternalName: input.InternalName})
	if check.ID != 0 {
		slog.Info("role already exists")
		return output, errors.New("role already exists")
	}
	output, err = s.NewCreateRoleUseCase.Execute(usecase.CreateRoleInputDTO{
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role created")
	return
}

func (s *RoleService) UpdateRole(input request.RequestUpdateRole) (output usecase.UpdateRoleOutputDTO, err error) {
	output, err = s.NewUpdateRoleUseCase.Execute(usecase.UpdateRoleInputDTO{
		ID:           input.ID,
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role updated")
	return
}

func (s *RoleService) DeleteRole(input request.RequestDeleteRole) (output usecase.DeleteRoleOutputDTO, err error) {
	output, err = s.NewDeleteRoleUseCase.Execute(usecase.DeleteRoleInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role deleted")
	return
}
