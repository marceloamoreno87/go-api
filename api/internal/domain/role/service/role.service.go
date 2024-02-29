package service

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/role/repository"
	"github.com/marceloamoreno/goapi/internal/domain/role/usecase"
	"github.com/marceloamoreno/goapi/internal/shared/helper"
)

type RoleServiceInterface interface {
	CreateRole(body io.ReadCloser) (err error)
	GetRole(id string) (output usecase.GetRoleOutputDTO, err error)
	GetRoles(limit string, offset string) (output []usecase.GetRolesOutputDTO, err error)
	UpdateRole(id string, body io.ReadCloser) (err error)
	DeleteRole(id string) (err error)
}

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
	s.repo.Begin()
	input := usecase.CreateRoleInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewCreateRoleUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	return
}

func (s *RoleService) UpdateRole(id string, body io.ReadCloser) (err error) {
	s.repo.Begin()
	input := usecase.UpdateRoleInputDTO{}
	if err = json.NewDecoder(body).Decode(&input); err != nil {
		slog.Info("err", err)
		return
	}

	if err = usecase.NewUpdateRoleUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	return
}

func (s *RoleService) DeleteRole(id string) (err error) {
	s.repo.Begin()
	input := usecase.DeleteRoleInputDTO{
		ID: helper.StrToInt32(id),
	}

	if err = usecase.NewDeleteRoleUseCase(s.repo).Execute(input); err != nil {
		s.repo.Rollback()
		slog.Info("err", err)
		return
	}
	s.repo.Commit()
	return
}
