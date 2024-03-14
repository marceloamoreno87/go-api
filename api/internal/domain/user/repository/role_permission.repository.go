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
	config.SQLCInterface
}

func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (repo *RolePermissionRepository) GetRolePermissionsByRole(id int32) (output []entityInterface.RolePermissionInterface, err error) {
	rp, err := repo.GetDbQueries().GetRolePermissionsByRole(context.Background(), id)
	if err != nil {
		return
	}
	return
}

// TODO: REFACTOR
func (repo *RolePermissionRepository) CreateRolePermission(rolePermission entityInterface.RolePermissionInterface) (output []entityInterface.RolePermissionInterface, err error) {

	errCh := make(chan error, len(rolePermission.GetPermissionIDs()))
	var wg sync.WaitGroup
	wg.Add(len(rolePermission.GetPermissionIDs()))

	for _, id := range rolePermission.GetPermissionIDs() {
		go func(permissionID int32) {
			defer wg.Done()
			rp, err := repo.GetDbQueries().WithTx(repo.GetTx()).CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
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

func (repo *RolePermissionRepository) DeleteRolePermission(rolePermission entityInterface.RolePermissionInterface, id int32) (output entityInterface.RolePermissionInterface, err error) {
	rp, err := repo.GetDbQueries().WithTx(repo.GetTx()).DeleteRolePermission(context.Background(), id)
	if err != nil {
		return
	}

	return
}
