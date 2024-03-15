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
	DB config.SQLCInterface
}

func NewRolePermissionRepository(DB config.SQLCInterface) *RolePermissionRepository {
	return &RolePermissionRepository{
		DB: DB,
	}
}

func (repo *RolePermissionRepository) GetRolePermissionsByRole(id int32) (output []entityInterface.RolePermissionInterface, err error) {

	rp, err := repo.DB.GetDbQueries().GetRolePermissionsByRole(context.Background(), id)
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

func (repo *RolePermissionRepository) CreateRolePermission(rolePermission entityInterface.RolePermissionInterface) (err error) {

	errCh := make(chan error, len(rolePermission.GetPermissionIDs()))
	var wg sync.WaitGroup
	wg.Add(len(rolePermission.GetPermissionIDs()))

	for _, id := range rolePermission.GetPermissionIDs() {
		go func(permissionID int32) {
			defer wg.Done()
			err := repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
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

func (repo *RolePermissionRepository) DeleteRolePermission(id int32) (err error) {

	return repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).DeleteRolePermission(context.Background(), id)
}
