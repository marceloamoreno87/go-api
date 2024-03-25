package service

import (
	"context"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/response"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type RolePermissionService struct {
	db                                  config.SQLCInterface
	GetRolePermissionsUseCase           usecase.GetRolePermissionsUseCase
	CreateRolePermissionUseCase         usecase.CreateRolePermissionUseCase
	DeleteRolePermissionByRoleIDUseCase usecase.DeleteRolePermissionByRoleIDUseCase
}

func NewRolePermissionService() *RolePermissionService {
	db := config.NewSqlc(config.DB)
	return &RolePermissionService{
		db:                                  db,
		GetRolePermissionsUseCase:           *usecase.NewGetRolePermissionsUseCase(db),
		CreateRolePermissionUseCase:         *usecase.NewCreateRolePermissionUseCase(db),
		DeleteRolePermissionByRoleIDUseCase: *usecase.NewDeleteRolePermissionByRoleIDUseCase(db),
	}
}

func (s *RolePermissionService) GetRolePermissions(ctx context.Context, input request.GetRolePermissionRequest) (output []response.GetRolePermissionsResponse, err error) {
	rolePermission, err := s.GetRolePermissionsUseCase.Execute(ctx, usecase.GetRolePermissionsInputDTO{RoleID: input.RoleID})
	if err != nil {
		slog.Info("err", err)
		return
	}

	for _, rolePermission := range rolePermission {
		output = append(output, response.GetRolePermissionsResponse{
			RoleID:        rolePermission.RoleID,
			PermissionIDs: rolePermission.PermissionIDs,
		})
	}
	slog.Info("Role permission found")
	return
}

func (s *RolePermissionService) CreateRolePermission(ctx context.Context, input request.CreateRolePermissionRequest) (output response.CreateRolePermissionResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	created, err := s.CreateRolePermissionUseCase.Execute(ctx, usecase.CreateRolePermissionInputDTO{
		RoleID:        input.RoleID,
		PermissionIDs: input.PermissionIDs,
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
	output = response.CreateRolePermissionResponse{
		RoleID:        created.RoleID,
		PermissionIDs: created.PermissionIDs,
	}

	slog.Info("Role permission created")
	return
}

func (s *RolePermissionService) UpdateRolePermission(ctx context.Context, input request.UpdateRolePermissionRequest) (output response.CreateRolePermissionResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	_, err = s.DeleteRolePermissionByRoleIDUseCase.Execute(ctx, usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
	if err != nil {
		errtx := tx.Rollback()
		if errtx != nil {
			slog.Info("errtx", errtx)
			return
		}
		slog.Info("err", err)
		return
	}

	updated, err := s.CreateRolePermissionUseCase.Execute(ctx, usecase.CreateRolePermissionInputDTO{
		RoleID:        input.RoleID,
		PermissionIDs: input.PermissionIDs,
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
	output = response.CreateRolePermissionResponse{
		RoleID:        updated.RoleID,
		PermissionIDs: updated.PermissionIDs,
	}
	slog.Info("Role permission updated")
	return
}

func (s *RolePermissionService) DeleteRolePermissionByRoleID(ctx context.Context, input request.DeleteRolePermissionByRoleIDRequest) (output response.DeleteRolePermissionByRoleIDResponse, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	deleted, err := s.DeleteRolePermissionByRoleIDUseCase.Execute(ctx, usecase.DeleteRolePermissionByRoleIDInputDTO{RoleID: input.RoleID})
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
	output = response.DeleteRolePermissionByRoleIDResponse{
		RoleID: deleted.RoleID,
	}
	slog.Info("Role permission deleted")
	return
}
