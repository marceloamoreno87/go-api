package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
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

func (s *RoleService) GetRole(ctx context.Context, input request.RequestGetRole) (output usecase.GetRoleOutputDTO, err error) {
	output, err = s.NewGetRoleUseCase.Execute(ctx, usecase.GetRoleInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Role found")
	return
}

func (s *RoleService) GetRoles(ctx context.Context, input request.RequestGetRoles) (output []usecase.GetRolesOutputDTO, err error) {
	output, err = s.NewGetRolesUseCase.Execute(ctx, usecase.GetRolesInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Roles found")
	return
}

func (s *RoleService) CreateRole(ctx context.Context, input request.RequestCreateRole) (output usecase.CreateRoleOutputDTO, err error) {
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
	output, err = s.NewCreateRoleUseCase.Execute(ctx, usecase.CreateRoleInputDTO{
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
	slog.Info("Role created")
	return
}

func (s *RoleService) UpdateRole(ctx context.Context, input request.RequestUpdateRole) (output usecase.UpdateRoleOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.NewUpdateRoleUseCase.Execute(ctx, usecase.UpdateRoleInputDTO{
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
	slog.Info("Role updated")
	return
}

func (s *RoleService) DeleteRole(ctx context.Context, input request.RequestDeleteRole) (output usecase.DeleteRoleOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.NewDeleteRoleUseCase.Execute(ctx, usecase.DeleteRoleInputDTO{ID: input.ID})
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
	slog.Info("Role deleted")
	return
}
