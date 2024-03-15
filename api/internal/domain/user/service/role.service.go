package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleService struct {
	DB config.SQLCInterface
}

func NewRoleService() *RoleService {
	return &RoleService{
		DB: config.Sqcl,
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

func (s *RoleService) CreateRole(body io.ReadCloser) (output usecase.CreateRoleOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.CreateRoleInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	check, _ := usecase.NewGetRoleByInternalNameUseCase().Execute(usecase.GetRoleByInternalNameInputDTO{InternalName: input.InternalName})
	if check.ID != 0 {
		slog.Info("role already exists")
		return output, errors.New("role already exists")
	}

	output, err = usecase.NewCreateRoleUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Role created")
	return
}

func (s *RoleService) UpdateRole(id int32, body io.ReadCloser) (output usecase.UpdateRoleOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.UpdateRoleInputDTO{
		ID: id,
	}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	output, err = usecase.NewUpdateRoleUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Role updated")
	return
}

func (s *RoleService) DeleteRole(id int32) (output usecase.DeleteRoleOutputDTO, err error) {
	s.DB.Begin()
	input := usecase.DeleteRoleInputDTO{
		ID: id,
	}

	output, err = usecase.NewDeleteRoleUseCase().Execute(input)
	if err != nil {
		s.DB.Rollback()
		slog.Info("err", err)
		return
	}
	s.DB.Commit()
	slog.Info("Role deleted")
	return
}
