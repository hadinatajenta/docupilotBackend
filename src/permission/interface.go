package permission

import "context"

type PermissionChecker interface {
	GetPermissionsByUserID(ctx context.Context, userID string) ([]Permission, error)
}
