// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAvatarStmt, err = db.PrepareContext(ctx, createAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAvatar: %w", err)
	}
	if q.createPermissionStmt, err = db.PrepareContext(ctx, createPermission); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePermission: %w", err)
	}
	if q.createRoleStmt, err = db.PrepareContext(ctx, createRole); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRole: %w", err)
	}
	if q.createRolePermissionStmt, err = db.PrepareContext(ctx, createRolePermission); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRolePermission: %w", err)
	}
	if q.createTokenStmt, err = db.PrepareContext(ctx, createToken); err != nil {
		return nil, fmt.Errorf("error preparing query CreateToken: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.createValidationUserStmt, err = db.PrepareContext(ctx, createValidationUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateValidationUser: %w", err)
	}
	if q.deleteAvatarStmt, err = db.PrepareContext(ctx, deleteAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAvatar: %w", err)
	}
	if q.deletePermissionStmt, err = db.PrepareContext(ctx, deletePermission); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePermission: %w", err)
	}
	if q.deleteRoleStmt, err = db.PrepareContext(ctx, deleteRole); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRole: %w", err)
	}
	if q.deleteRolePermissionStmt, err = db.PrepareContext(ctx, deleteRolePermission); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRolePermission: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAvatarStmt, err = db.PrepareContext(ctx, getAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query GetAvatar: %w", err)
	}
	if q.getAvatarsStmt, err = db.PrepareContext(ctx, getAvatars); err != nil {
		return nil, fmt.Errorf("error preparing query GetAvatars: %w", err)
	}
	if q.getPermissionStmt, err = db.PrepareContext(ctx, getPermission); err != nil {
		return nil, fmt.Errorf("error preparing query GetPermission: %w", err)
	}
	if q.getPermissionByInternalNameStmt, err = db.PrepareContext(ctx, getPermissionByInternalName); err != nil {
		return nil, fmt.Errorf("error preparing query GetPermissionByInternalName: %w", err)
	}
	if q.getPermissionsStmt, err = db.PrepareContext(ctx, getPermissions); err != nil {
		return nil, fmt.Errorf("error preparing query GetPermissions: %w", err)
	}
	if q.getRoleStmt, err = db.PrepareContext(ctx, getRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetRole: %w", err)
	}
	if q.getRoleByInternalNameStmt, err = db.PrepareContext(ctx, getRoleByInternalName); err != nil {
		return nil, fmt.Errorf("error preparing query GetRoleByInternalName: %w", err)
	}
	if q.getRolePermissionStmt, err = db.PrepareContext(ctx, getRolePermission); err != nil {
		return nil, fmt.Errorf("error preparing query GetRolePermission: %w", err)
	}
	if q.getRolePermissionsByRoleStmt, err = db.PrepareContext(ctx, getRolePermissionsByRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetRolePermissionsByRole: %w", err)
	}
	if q.getRolesStmt, err = db.PrepareContext(ctx, getRoles); err != nil {
		return nil, fmt.Errorf("error preparing query GetRoles: %w", err)
	}
	if q.getTokenByUserStmt, err = db.PrepareContext(ctx, getTokenByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetTokenByUser: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserWithAvatarStmt, err = db.PrepareContext(ctx, getUserWithAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserWithAvatar: %w", err)
	}
	if q.getUserWithRoleStmt, err = db.PrepareContext(ctx, getUserWithRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserWithRole: %w", err)
	}
	if q.getUserWithRoleAndAvatarStmt, err = db.PrepareContext(ctx, getUserWithRoleAndAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserWithRoleAndAvatar: %w", err)
	}
	if q.getUserWithValidationUserStmt, err = db.PrepareContext(ctx, getUserWithValidationUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserWithValidationUser: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.getUsersWithAvatarStmt, err = db.PrepareContext(ctx, getUsersWithAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsersWithAvatar: %w", err)
	}
	if q.getUsersWithRoleStmt, err = db.PrepareContext(ctx, getUsersWithRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsersWithRole: %w", err)
	}
	if q.getUsersWithRoleAndAvatarStmt, err = db.PrepareContext(ctx, getUsersWithRoleAndAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsersWithRoleAndAvatar: %w", err)
	}
	if q.getValidationUserStmt, err = db.PrepareContext(ctx, getValidationUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetValidationUser: %w", err)
	}
	if q.getValidationUserByHashStmt, err = db.PrepareContext(ctx, getValidationUserByHash); err != nil {
		return nil, fmt.Errorf("error preparing query GetValidationUserByHash: %w", err)
	}
	if q.registerUserStmt, err = db.PrepareContext(ctx, registerUser); err != nil {
		return nil, fmt.Errorf("error preparing query RegisterUser: %w", err)
	}
	if q.revokeTokenByUserStmt, err = db.PrepareContext(ctx, revokeTokenByUser); err != nil {
		return nil, fmt.Errorf("error preparing query RevokeTokenByUser: %w", err)
	}
	if q.setUserValidationUsedStmt, err = db.PrepareContext(ctx, setUserValidationUsed); err != nil {
		return nil, fmt.Errorf("error preparing query SetUserValidationUsed: %w", err)
	}
	if q.updateAvatarStmt, err = db.PrepareContext(ctx, updateAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAvatar: %w", err)
	}
	if q.updatePasswordUserStmt, err = db.PrepareContext(ctx, updatePasswordUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePasswordUser: %w", err)
	}
	if q.updatePermissionStmt, err = db.PrepareContext(ctx, updatePermission); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePermission: %w", err)
	}
	if q.updateRoleStmt, err = db.PrepareContext(ctx, updateRole); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateRole: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAvatarStmt != nil {
		if cerr := q.createAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAvatarStmt: %w", cerr)
		}
	}
	if q.createPermissionStmt != nil {
		if cerr := q.createPermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPermissionStmt: %w", cerr)
		}
	}
	if q.createRoleStmt != nil {
		if cerr := q.createRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRoleStmt: %w", cerr)
		}
	}
	if q.createRolePermissionStmt != nil {
		if cerr := q.createRolePermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRolePermissionStmt: %w", cerr)
		}
	}
	if q.createTokenStmt != nil {
		if cerr := q.createTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTokenStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.createValidationUserStmt != nil {
		if cerr := q.createValidationUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createValidationUserStmt: %w", cerr)
		}
	}
	if q.deleteAvatarStmt != nil {
		if cerr := q.deleteAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAvatarStmt: %w", cerr)
		}
	}
	if q.deletePermissionStmt != nil {
		if cerr := q.deletePermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePermissionStmt: %w", cerr)
		}
	}
	if q.deleteRoleStmt != nil {
		if cerr := q.deleteRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRoleStmt: %w", cerr)
		}
	}
	if q.deleteRolePermissionStmt != nil {
		if cerr := q.deleteRolePermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRolePermissionStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAvatarStmt != nil {
		if cerr := q.getAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAvatarStmt: %w", cerr)
		}
	}
	if q.getAvatarsStmt != nil {
		if cerr := q.getAvatarsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAvatarsStmt: %w", cerr)
		}
	}
	if q.getPermissionStmt != nil {
		if cerr := q.getPermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPermissionStmt: %w", cerr)
		}
	}
	if q.getPermissionByInternalNameStmt != nil {
		if cerr := q.getPermissionByInternalNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPermissionByInternalNameStmt: %w", cerr)
		}
	}
	if q.getPermissionsStmt != nil {
		if cerr := q.getPermissionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPermissionsStmt: %w", cerr)
		}
	}
	if q.getRoleStmt != nil {
		if cerr := q.getRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRoleStmt: %w", cerr)
		}
	}
	if q.getRoleByInternalNameStmt != nil {
		if cerr := q.getRoleByInternalNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRoleByInternalNameStmt: %w", cerr)
		}
	}
	if q.getRolePermissionStmt != nil {
		if cerr := q.getRolePermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRolePermissionStmt: %w", cerr)
		}
	}
	if q.getRolePermissionsByRoleStmt != nil {
		if cerr := q.getRolePermissionsByRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRolePermissionsByRoleStmt: %w", cerr)
		}
	}
	if q.getRolesStmt != nil {
		if cerr := q.getRolesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRolesStmt: %w", cerr)
		}
	}
	if q.getTokenByUserStmt != nil {
		if cerr := q.getTokenByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTokenByUserStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserWithAvatarStmt != nil {
		if cerr := q.getUserWithAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserWithAvatarStmt: %w", cerr)
		}
	}
	if q.getUserWithRoleStmt != nil {
		if cerr := q.getUserWithRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserWithRoleStmt: %w", cerr)
		}
	}
	if q.getUserWithRoleAndAvatarStmt != nil {
		if cerr := q.getUserWithRoleAndAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserWithRoleAndAvatarStmt: %w", cerr)
		}
	}
	if q.getUserWithValidationUserStmt != nil {
		if cerr := q.getUserWithValidationUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserWithValidationUserStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.getUsersWithAvatarStmt != nil {
		if cerr := q.getUsersWithAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersWithAvatarStmt: %w", cerr)
		}
	}
	if q.getUsersWithRoleStmt != nil {
		if cerr := q.getUsersWithRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersWithRoleStmt: %w", cerr)
		}
	}
	if q.getUsersWithRoleAndAvatarStmt != nil {
		if cerr := q.getUsersWithRoleAndAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersWithRoleAndAvatarStmt: %w", cerr)
		}
	}
	if q.getValidationUserStmt != nil {
		if cerr := q.getValidationUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getValidationUserStmt: %w", cerr)
		}
	}
	if q.getValidationUserByHashStmt != nil {
		if cerr := q.getValidationUserByHashStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getValidationUserByHashStmt: %w", cerr)
		}
	}
	if q.registerUserStmt != nil {
		if cerr := q.registerUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing registerUserStmt: %w", cerr)
		}
	}
	if q.revokeTokenByUserStmt != nil {
		if cerr := q.revokeTokenByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing revokeTokenByUserStmt: %w", cerr)
		}
	}
	if q.setUserValidationUsedStmt != nil {
		if cerr := q.setUserValidationUsedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setUserValidationUsedStmt: %w", cerr)
		}
	}
	if q.updateAvatarStmt != nil {
		if cerr := q.updateAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAvatarStmt: %w", cerr)
		}
	}
	if q.updatePasswordUserStmt != nil {
		if cerr := q.updatePasswordUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePasswordUserStmt: %w", cerr)
		}
	}
	if q.updatePermissionStmt != nil {
		if cerr := q.updatePermissionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePermissionStmt: %w", cerr)
		}
	}
	if q.updateRoleStmt != nil {
		if cerr := q.updateRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateRoleStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                              DBTX
	tx                              *sql.Tx
	createAvatarStmt                *sql.Stmt
	createPermissionStmt            *sql.Stmt
	createRoleStmt                  *sql.Stmt
	createRolePermissionStmt        *sql.Stmt
	createTokenStmt                 *sql.Stmt
	createUserStmt                  *sql.Stmt
	createValidationUserStmt        *sql.Stmt
	deleteAvatarStmt                *sql.Stmt
	deletePermissionStmt            *sql.Stmt
	deleteRoleStmt                  *sql.Stmt
	deleteRolePermissionStmt        *sql.Stmt
	deleteUserStmt                  *sql.Stmt
	getAvatarStmt                   *sql.Stmt
	getAvatarsStmt                  *sql.Stmt
	getPermissionStmt               *sql.Stmt
	getPermissionByInternalNameStmt *sql.Stmt
	getPermissionsStmt              *sql.Stmt
	getRoleStmt                     *sql.Stmt
	getRoleByInternalNameStmt       *sql.Stmt
	getRolePermissionStmt           *sql.Stmt
	getRolePermissionsByRoleStmt    *sql.Stmt
	getRolesStmt                    *sql.Stmt
	getTokenByUserStmt              *sql.Stmt
	getUserStmt                     *sql.Stmt
	getUserByEmailStmt              *sql.Stmt
	getUserWithAvatarStmt           *sql.Stmt
	getUserWithRoleStmt             *sql.Stmt
	getUserWithRoleAndAvatarStmt    *sql.Stmt
	getUserWithValidationUserStmt   *sql.Stmt
	getUsersStmt                    *sql.Stmt
	getUsersWithAvatarStmt          *sql.Stmt
	getUsersWithRoleStmt            *sql.Stmt
	getUsersWithRoleAndAvatarStmt   *sql.Stmt
	getValidationUserStmt           *sql.Stmt
	getValidationUserByHashStmt     *sql.Stmt
	registerUserStmt                *sql.Stmt
	revokeTokenByUserStmt           *sql.Stmt
	setUserValidationUsedStmt       *sql.Stmt
	updateAvatarStmt                *sql.Stmt
	updatePasswordUserStmt          *sql.Stmt
	updatePermissionStmt            *sql.Stmt
	updateRoleStmt                  *sql.Stmt
	updateUserStmt                  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                              tx,
		tx:                              tx,
		createAvatarStmt:                q.createAvatarStmt,
		createPermissionStmt:            q.createPermissionStmt,
		createRoleStmt:                  q.createRoleStmt,
		createRolePermissionStmt:        q.createRolePermissionStmt,
		createTokenStmt:                 q.createTokenStmt,
		createUserStmt:                  q.createUserStmt,
		createValidationUserStmt:        q.createValidationUserStmt,
		deleteAvatarStmt:                q.deleteAvatarStmt,
		deletePermissionStmt:            q.deletePermissionStmt,
		deleteRoleStmt:                  q.deleteRoleStmt,
		deleteRolePermissionStmt:        q.deleteRolePermissionStmt,
		deleteUserStmt:                  q.deleteUserStmt,
		getAvatarStmt:                   q.getAvatarStmt,
		getAvatarsStmt:                  q.getAvatarsStmt,
		getPermissionStmt:               q.getPermissionStmt,
		getPermissionByInternalNameStmt: q.getPermissionByInternalNameStmt,
		getPermissionsStmt:              q.getPermissionsStmt,
		getRoleStmt:                     q.getRoleStmt,
		getRoleByInternalNameStmt:       q.getRoleByInternalNameStmt,
		getRolePermissionStmt:           q.getRolePermissionStmt,
		getRolePermissionsByRoleStmt:    q.getRolePermissionsByRoleStmt,
		getRolesStmt:                    q.getRolesStmt,
		getTokenByUserStmt:              q.getTokenByUserStmt,
		getUserStmt:                     q.getUserStmt,
		getUserByEmailStmt:              q.getUserByEmailStmt,
		getUserWithAvatarStmt:           q.getUserWithAvatarStmt,
		getUserWithRoleStmt:             q.getUserWithRoleStmt,
		getUserWithRoleAndAvatarStmt:    q.getUserWithRoleAndAvatarStmt,
		getUserWithValidationUserStmt:   q.getUserWithValidationUserStmt,
		getUsersStmt:                    q.getUsersStmt,
		getUsersWithAvatarStmt:          q.getUsersWithAvatarStmt,
		getUsersWithRoleStmt:            q.getUsersWithRoleStmt,
		getUsersWithRoleAndAvatarStmt:   q.getUsersWithRoleAndAvatarStmt,
		getValidationUserStmt:           q.getValidationUserStmt,
		getValidationUserByHashStmt:     q.getValidationUserByHashStmt,
		registerUserStmt:                q.registerUserStmt,
		revokeTokenByUserStmt:           q.revokeTokenByUserStmt,
		setUserValidationUsedStmt:       q.setUserValidationUsedStmt,
		updateAvatarStmt:                q.updateAvatarStmt,
		updatePasswordUserStmt:          q.updatePasswordUserStmt,
		updatePermissionStmt:            q.updatePermissionStmt,
		updateRoleStmt:                  q.updateRoleStmt,
		updateUserStmt:                  q.updateUserStmt,
	}
}
