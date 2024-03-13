package repository

import (
	"context"
	"sync"

	"github.com/marceloamoreno/goapi/config"
	PermissionEntity "github.com/marceloamoreno/goapi/internal/domain/permission/entity"
	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	RolePermissionEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type RolePermissionRepositoryInterface interface {
	GetRolePermissionsByRole(id int32) (rolePermissions *entity.RolePermission, err error)
	CreateRolePermission(rolePermission *entity.RolePermission) (err error)
	UpdateRolePermission(rolePermission *entity.RolePermission, id int32) (err error)
	config.SQLCInterface
}

type RolePermissionRepository struct {
	config.SQLCInterface
}

func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (repo *RolePermissionRepository) GetRolePermissionsByRole(id int32) (rolePermissions *RolePermissionEntity.RolePermission, err error) {
	rps, err := repo.GetDbQueries().GetRolePermissionsByRole(context.Background(), id)
	if err != nil {
		return
	}
	rolePermissions = &RolePermissionEntity.RolePermission{
		Role: &RoleEntity.Role{
			ID:           rps[0].RoleID,
			Name:         rps[0].Name,
			InternalName: rps[0].InternalName,
			Description:  rps[0].Description,
			CreatedAt:    rps[0].CreatedAt,
			UpdatedAt:    rps[0].UpdatedAt,
		},
	}

	for _, rp := range rps {
		rolePermissions.Permissions = append(rolePermissions.Permissions, &PermissionEntity.Permission{
			ID:           rp.PermissionID,
			Name:         rp.Name_2,
			InternalName: rp.InternalName_2,
			Description:  rp.Description_2,
			CreatedAt:    rp.CreatedAt_2,
			UpdatedAt:    rp.UpdatedAt_2,
		})
	}
	return
}

func (repo *RolePermissionRepository) CreateRolePermission(rolePermission *RolePermissionEntity.RolePermission) (err error) {

	errCh := make(chan error, len(rolePermission.PermissionIDs))
	var wg sync.WaitGroup
	wg.Add(len(rolePermission.PermissionIDs))

	for _, id := range rolePermission.PermissionIDs {
		go func(permissionID int32) {
			defer wg.Done()
			err := repo.GetDbQueries().WithTx(repo.GetTx()).CreateRolePermission(context.Background(), db.CreateRolePermissionParams{
				RoleID:       rolePermission.RoleID,
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

func (repo *RolePermissionRepository) UpdateRolePermission(rolePermission *RolePermissionEntity.RolePermission, id int32) (err error) {
	if err = repo.GetDbQueries().WithTx(repo.GetTx()).DeleteRolePermission(context.Background(), id); err != nil {
		return
	}
	return repo.CreateRolePermission(rolePermission)
}
