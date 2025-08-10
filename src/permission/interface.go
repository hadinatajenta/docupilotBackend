package permission

import "context"

type PermissionChecker interface {
	GetPermissionsByUserID(ctx context.Context, userID string) ([]Permission, error)
	CreatePermission(ctx context.Context, permission *Permission) error
}

type Service interface {
	CreatePermission(ctx context.Context, permission *Permission) error
}
