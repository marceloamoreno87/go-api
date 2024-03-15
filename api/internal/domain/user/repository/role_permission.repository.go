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

func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{
		DB: config.Sqcl,
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

// TODO: REFACTOR
func (repo *RolePermissionRepository) CreateRolePermission(rolePermission entityInterface.RolePermissionInterface) (output []entityInterface.RolePermissionInterface, err error) {

	errCh := make(chan error, len(rolePermission.GetPermissionIDs()))
	rpCh := make(chan entityInterface.RolePermissionInterface, len(rolePermission.GetPermissionIDs()))
	var wg sync.WaitGroup
	wg.Add(len(rolePermission.GetPermissionIDs()))

	for _, id := range rolePermission.GetPermissionIDs() {
		go func(permissionID int32) {
			defer wg.Done()
			rp, err := repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
				RoleID:       rolePermission.GetRoleID(),
				PermissionID: permissionID,
			})

			rpCh <- &entity.RolePermission{
				ID:     rp.ID,
				RoleID: rp.RoleID,
			}

			if err != nil {
				errCh <- err
			}
		}(id)
	}
	wg.Wait()
	close(errCh)
	close(rpCh)
	if len(errCh) > 0 {
		return nil, <-errCh
	}
	output = append(output, <-rpCh)
	return
}

func (repo *RolePermissionRepository) DeleteRolePermission(rolePermission entityInterface.RolePermissionInterface, id int32) (output entityInterface.RolePermissionInterface, err error) {
	rp, err := repo.DB.GetDbQueries().WithTx(repo.DB.GetTx()).DeleteRolePermission(context.Background(), id)
	if err != nil {
		return
	}

	output = &entity.RolePermission{
		ID:     rp.ID,
		RoleID: rp.RoleID,
	}

	return
}
