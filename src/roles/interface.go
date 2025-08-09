package roles

import "context"

type Repository interface {
	GetRolesByUserID(ctx context.Context, userID string) ([]Role, error)
	GetPermissionsByUserID(ctx context.Context, userID string) ([]Permission, error)
}
