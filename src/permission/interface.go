package permission

import "context"

type PermissionChecker interface {
	GetPermissionsByUserID(ctx context.Context, userID string) ([]Permission, error)
	CreatePermission(ctx context.Context, permission *Permission) error
	GetAllPermissions(ctx context.Context) ([]Permission, error)
	GetPermissionByRoleId(ctx context.Context, roleId string) ([]Permission, error)
}

type Service interface {
	CreatePermission(ctx context.Context, permission *Permission) error
	GetAllPermissions(ctx context.Context) (PermissionResponse, error)
	GetPermissionByRoleId(ctx context.Context, userID string) (GetPermissionByUidRes, error)
}
