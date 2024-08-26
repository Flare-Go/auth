// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	AssignPermissionToRole(ctx context.Context, arg AssignPermissionToRoleParams) error
	AssignRoleToUser(ctx context.Context, arg AssignRoleToUserParams) error
	CreatePermission(ctx context.Context, arg CreatePermissionParams) error
	CreateRole(ctx context.Context, arg CreateRoleParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (uint32, error)
	DeletePermission(ctx context.Context, id uint32) error
	DeleteRole(ctx context.Context, id uint32) error
	DeleteUser(ctx context.Context, id uint32) error
	GetPermissionByID(ctx context.Context, id uint32) (*GetPermissionByIDRow, error)
	GetRoleByID(ctx context.Context, id uint32) (*GetRoleByIDRow, error)
	GetRolePermissions(ctx context.Context, roleID uint32) ([]*Permission, error)
	GetUserByEmail(ctx context.Context, email string) (*GetUserByEmailRow, error)
	GetUserByID(ctx context.Context, id uint32) (*GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, username string) (*GetUserByUsernameRow, error)
	GetUserRoles(ctx context.Context, userID uint32) ([]*Role, error)
	ListRoles(ctx context.Context) ([]*ListRolesRow, error)
	ListUsers(ctx context.Context) ([]*ListUsersRow, error)
	RemovePermissionFromRole(ctx context.Context, arg RemovePermissionFromRoleParams) error
	RemoveRoleFromUser(ctx context.Context, arg RemoveRoleFromUserParams) error
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) error
}

var _ Querier = (*Queries)(nil)
