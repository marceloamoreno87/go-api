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

func (repo *RolePermissionRepository) GetRolePermissionsByRole(id int32) (rolePermissions entityInterface.RolePermissionInterface, err error) {
	rps, err := repo.GetDbQueries().GetRolePermissionsByRole(context.Background(), id)
	if err != nil {
		return
	}
	rolePermissions = &entity.RolePermission{
		Role: &entity.Role{
			ID:           rps[0].RoleID,
			Name:         rps[0].Name,
			InternalName: rps[0].InternalName,
			Description:  rps[0].Description,
			CreatedAt:    rps[0].CreatedAt,
			UpdatedAt:    rps[0].UpdatedAt,
		},
	}

	for _, rp := range rps {
		rolePermissions.SetPermissions(append(rolePermissions.GetPermissions(), &entity.Permission{
			ID:           rp.PermissionID,
			Name:         rp.Name_2,
			InternalName: rp.InternalName_2,
			Description:  rp.Description_2,
			CreatedAt:    rp.CreatedAt_2,
			UpdatedAt:    rp.UpdatedAt_2,
		}))
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
			err := repo.GetDbQueries().WithTx(repo.GetTx()).CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
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

func (repo *RolePermissionRepository) UpdateRolePermission(rolePermission entityInterface.RolePermissionInterface, id int32) (err error) {
	if err = repo.GetDbQueries().WithTx(repo.GetTx()).DeleteRolePermission(context.Background(), id); err != nil {
		return
	}
	return repo.CreateRolePermission(rolePermission)
}
