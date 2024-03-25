package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/request"
	"github.com/marceloamoreno/goapi/internal/domain/user/usecase"
)

type PermissionService struct {
	db                                    config.SQLCInterface
	NewGetPermissionUseCase               usecase.GetPermissionUseCase
	NewGetPermissionsUseCase              usecase.GetPermissionsUseCase
	NewCreatePermissionUseCase            usecase.CreatePermissionUseCase
	NewUpdatePermissionUseCase            usecase.UpdatePermissionUseCase
	NewDeletePermissionUseCase            usecase.DeletePermissionUseCase
	NewGetPermissionByInternalNameUseCase usecase.GetPermissionByInternalNameUseCase
}

func NewPermissionService() *PermissionService {
	db := config.NewSqlc(config.DB)
	return &PermissionService{
		db:                                    db,
		NewGetPermissionUseCase:               *usecase.NewGetPermissionUseCase(db),
		NewGetPermissionsUseCase:              *usecase.NewGetPermissionsUseCase(db),
		NewCreatePermissionUseCase:            *usecase.NewCreatePermissionUseCase(db),
		NewUpdatePermissionUseCase:            *usecase.NewUpdatePermissionUseCase(db),
		NewDeletePermissionUseCase:            *usecase.NewDeletePermissionUseCase(db),
		NewGetPermissionByInternalNameUseCase: *usecase.NewGetPermissionByInternalNameUseCase(db),
	}
}

func (s *PermissionService) GetPermission(ctx context.Context, input request.RequestGetPermission) (output usecase.GetPermissionOutputDTO, err error) {
	output, err = s.NewGetPermissionUseCase.Execute(ctx, usecase.GetPermissionInputDTO{ID: input.ID})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permission found")
	return
}

func (s *PermissionService) GetPermissions(ctx context.Context, input request.RequestGetPermissions) (output []usecase.GetPermissionsOutputDTO, err error) {
	output, err = s.NewGetPermissionsUseCase.Execute(ctx, usecase.GetPermissionsInputDTO{Limit: input.Limit, Offset: input.Offset})
	if err != nil {
		slog.Info("err", err)
		return
	}
	slog.Info("Permissions found")
	return
}

func (s *PermissionService) CreatePermission(ctx context.Context, input request.RequestCreatePermission) (output usecase.CreatePermissionOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	check, _ := s.NewGetPermissionByInternalNameUseCase.Execute(ctx, usecase.GetPermissionByInternalNameInputDTO{
		InternalName: input.InternalName,
	})

	if check.ID != 0 {
		slog.Info("Permission already exists")
		return output, errors.New("permission already exists")
	}

	output, err = s.NewCreatePermissionUseCase.Execute(ctx, usecase.CreatePermissionInputDTO{
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
	slog.Info("Permission created")
	return
}

func (s *PermissionService) UpdatePermission(ctx context.Context, input request.RequestUpdatePermission) (output usecase.UpdatePermissionOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.NewUpdatePermissionUseCase.Execute(ctx, usecase.UpdatePermissionInputDTO{
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
	slog.Info("Permission updated")
	return
}

func (s *PermissionService) DeletePermission(ctx context.Context, input request.RequestDeletePermission) (output usecase.DeletePermissionOutputDTO, err error) {
	tx, err := s.db.GetDbConn().Begin()
	if err != nil {
		slog.Info("err", err)
		return
	}
	s.db.SetTx(tx)
	output, err = s.NewDeletePermissionUseCase.Execute(ctx, usecase.DeletePermissionInputDTO{ID: input.ID})
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
	slog.Info("Permission deleted")
	return
}
