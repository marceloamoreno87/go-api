package repository

import (
	"context"
	"sync"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type RolePermissionRepository struct {
	db config.SQLCInterface
}

func NewRolePermissionRepository(db config.SQLCInterface) *RolePermissionRepository {
	return &RolePermissionRepository{
		db: db,
	}
}

func (repo *RolePermissionRepository) GetRolePermissionsByRole(ctx context.Context, id int32) (output []entityInterface.RolePermissionInterface, err error) {
	rp, err := repo.db.GetDbQueries().GetRolePermissionsByRole(ctx, id)
	if err != nil {
		return
	}

	for _, r := range rp {
		output = append(output, &entity.RolePermission{
			ID:     r.ID,
			RoleID: r.RoleID,
		})
	}
	return
}

func (repo *RolePermissionRepository) CreateRolePermission(ctx context.Context, rolePermission entityInterface.RolePermissionInterface) (err error) {
	errCh := make(chan error, len(rolePermission.GetPermissionIDs()))
	var wg sync.WaitGroup
	wg.Add(len(rolePermission.GetPermissionIDs()))

	for _, id := range rolePermission.GetPermissionIDs() {
		go func(permissionID int32) {
			defer wg.Done()
			err := repo.db.GetDbQueries().WithTx(repo.db.GetTx()).CreateRolePermission(ctx, db.CreateRolePermissionParams{
				RoleID:       rolePermission.GetRoleID(),
				PermissionID: permissionID,
			})

			if err != nil {
				errCh <- err
			}
		}(id)
	}
	wg.Wait()
	close(errCh)
	if len(errCh) > 0 {
		return <-errCh
	}
	return
}

func (repo *RolePermissionRepository) DeleteRolePermissionByRoleID(ctx context.Context, id int32) (err error) {
	return repo.db.GetDbQueries().WithTx(repo.db.GetTx()).DeleteRolePermissionByRoleID(ctx, id)
}
