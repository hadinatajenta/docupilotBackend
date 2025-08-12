package roles

import (
	"context"
)

type Repository interface {
	// role
	GetRolesByUserID(ctx context.Context, userID string) ([]Role, error)
	Create(ctx context.Context, role *Role) error
	CheckRoleExist(ctx context.Context, role string) (bool, error)
	GetAllRoles(ctx context.Context) ([]Role, error)

	// role permission
	GetByID(ctx context.Context, id string) (*Role, error)
	PermissionsExist(ctx context.Context, ids []string) (bool, error)
	AssignPermissions(ctx context.Context, roleID string, permissionIDs []string) error
	GetPermissionByIDS(ctx context.Context, ids []string) ([]GetNameID, error)
}

type Service interface {
	// role
	CreateRole(ctx context.Context, role *Role) error
	GetAllRoles(ctx context.Context) ([]Role, error)
	// role permission
	AssignPermissionsToRole(ctx context.Context, roleID string, permissionIDs []string) (RolePermissionRes, error)
}
