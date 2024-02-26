package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type RoleService struct {
	repo repository.RoleRepositoryInterface
}

func NewRoleService(repo repository.RoleRepositoryInterface) *RoleService {
	return &RoleService{
		repo: repo,
	}
}

func (s *RoleService) GetRole(id string) (output usecase.GetRoleOutputDTO, err error) {

	input := usecase.GetRoleInputDTO{
		ID: helper.StrToInt32(id),
	}

	output, err = usecase.NewGetRoleUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RoleService) GetRoles(limit string, offset string) (output []usecase.GetRolesOutputDTO, err error) {

	input := usecase.GetRolesInputDTO{
		Limit:  helper.StrToInt32(limit),
		Offset: helper.StrToInt32(offset),
	}

	output, err = usecase.NewGetRolesUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RoleService) CreateRole(body io.ReadCloser) (err error) {

	input := usecase.CreateRoleInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewCreateRoleUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RoleService) UpdateRole(id string, body io.ReadCloser) (err error) {

	input := usecase.UpdateRoleInputDTO{}
	err = json.NewDecoder(body).Decode(&input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	err = usecase.NewUpdateRoleUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}

func (s *RoleService) DeleteRole(id string) (err error) {

	input := usecase.DeleteRoleInputDTO{
		ID: helper.StrToInt32(id),
	}

	err = usecase.NewDeleteRoleUseCase(s.repo).Execute(input)
	if err != nil {
		slog.Info("err", err)
		return
	}

	return
}
