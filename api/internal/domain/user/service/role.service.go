package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/repository"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleServiceInterface interface {
	CreateRole(body io.ReadCloser) (err error)
	GetRole(id int32) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(limit int32, offset int32) (output []usecase.GetRolesOutputDTO, err error)
	UpdateRole(id int32, body io.ReadCloser) (err error)
	DeleteRole(id int32) (err error)
}

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{
	}
}

func (s *RoleService) GetRole(id int32) (output usecase.GetRoleOutputDTO, err error) {

	input := usecase.GetRoleInputDTO{
		ID: id,
	}

	output, err = usecase.NewGetRoleUseCase().Execute(input)
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

	output, err = usecase.NewGetRolesUseCase().Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Roles found")
	return
}

func (s *RoleService) CreateRole(body io.ReadCloser) (err error) {
	s.roleRepo.Begin()
	input := usecase.CreateRoleInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, _ := usecase.NewGetRoleByInternalNameUseCase().Execute(usecase.GetRoleByInternalNameInputDTO{InternalName: input.InternalName})
	if output.ID != 0 {
		slog.Info("role already exists")
		return errors.New("role already exists")
	}

	if err = usecase.NewCreateRoleUseCase().Execute(input); err != nil {
		s.roleRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.roleRepo.Commit()
	slog.Info("Role created")
	return
}

func (s *RoleService) UpdateRole(id int32, body io.ReadCloser) (err error) {
	s.roleRepo.Begin()
	input := usecase.UpdateRoleInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdateRoleUseCase().Execute(input); err != nil {
		s.roleRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.roleRepo.Commit()
	slog.Info("Role updated")
	return
}

func (s *RoleService) DeleteRole(id int32) (err error) {
	s.roleRepo.Begin()
	input := usecase.DeleteRoleInputDTO{
		ID: id,
	}

	if err = usecase.NewDeleteRoleUseCase().Execute(input); err != nil {
		s.roleRepo.Rollback()
		slog.Info("err", err)
		return
	}
	s.roleRepo.Commit()
	slog.Info("Role deleted")
	return
}
