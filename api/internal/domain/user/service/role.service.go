package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/response"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RoleService struct {
	db                              config.SQLCInterface
	NewGetRoleUseCase               usecase.GetRoleUseCase
	NewGetRolesUseCase              usecase.GetRolesUseCase
	NewCreateRoleUseCase            usecase.CreateRoleUseCase
	NewUpdateRoleUseCase            usecase.UpdateRoleUseCase
	NewDeleteRoleUseCase            usecase.DeleteRoleUseCase
	NewGetRoleByInternalNameUseCase usecase.GetRoleByInternalNameUseCase
}

func NewRoleService() *RoleService {
	db := config.NewSqlc(config.DB)
	return &RoleService{
		db:                              db,
		NewGetRoleUseCase:               *usecase.NewGetRoleUseCase(db),
		NewGetRolesUseCase:              *usecase.NewGetRolesUseCase(db),
		NewCreateRoleUseCase:            *usecase.NewCreateRoleUseCase(db),
		NewUpdateRoleUseCase:            *usecase.NewUpdateRoleUseCase(db),
		NewDeleteRoleUseCase:            *usecase.NewDeleteRoleUseCase(db),
		NewGetRoleByInternalNameUseCase: *usecase.NewGetRoleByInternalNameUseCase(db),
	}
}

func (s *RoleService) GetRole(ctx context.Context, input request.GetRoleRequest) (output response.GetRoleResponse, err error) {
	role, err := s.NewGetRoleUseCase.Execute(ctx, usecase.GetRoleInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	output = response.GetRoleResponse{
		ID:           role.ID,
		Name:         role.Name,
		InternalName: role.InternalName,
		Description:  role.Description,
	}
	slog.Info("Role found")
	return
}

func (s *RoleService) GetRoles(ctx context.Context, input request.GetRolesRequest) (output []response.GetRolesResponse, err error) {
	roles, err := s.NewGetRolesUseCase.Execute(ctx, usecase.GetRolesInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	for _, role := range roles {
		output = append(output, response.GetRolesResponse{
			ID:           role.ID,
			Name:         role.Name,
			InternalName: role.InternalName,
			Description:  role.Description,
		})
	}
	slog.Info("Roles found")
	return
}

func (s *RoleService) CreateRole(ctx context.Context, input request.CreateRoleRequest) (output response.CreateRoleResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	check, _ := s.NewGetRoleByInternalNameUseCase.Execute(ctx, usecase.GetRoleByInternalNameInputDTO{InternalName: input.InternalName})
	if check.ID != 0 {
		slog.Info("role already exists")
		return output, errors.New("role already exists")
	}
	created, err := s.NewCreateRoleUseCase.Execute(ctx, usecase.CreateRoleInputDTO{
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.CreateRoleResponse{
		Name:         created.Name,
		InternalName: created.InternalName,
		Description:  created.Description,
	}
	slog.Info("Role created")
	return
}

func (s *RoleService) UpdateRole(ctx context.Context, input request.UpdateRoleRequest) (output response.UpdateRoleResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	updated, err := s.NewUpdateRoleUseCase.Execute(ctx, usecase.UpdateRoleInputDTO{
		ID:           input.ID,
		Name:         input.Name,
		InternalName: input.InternalName,
		Description:  input.Description,
	})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.UpdateRoleResponse{
		ID:           updated.ID,
		Name:         updated.Name,
		InternalName: updated.InternalName,
		Description:  updated.Description,
	}
	slog.Info("Role updated")
	return
}

func (s *RoleService) DeleteRole(ctx context.Context, input request.DeleteRoleRequest) (output response.DeleteRoleResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	deleted, err := s.NewDeleteRoleUseCase.Execute(ctx, usecase.DeleteRoleInputDTO{ID: input.ID})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}
	errtx := tx.Commit()
	if errtx != nil {
		slog.Info("errtx", errtx)
		return
	}
	output = response.DeleteRoleResponse{
		ID: deleted.ID,
	}
	slog.Info("Role deleted")
	return
}
