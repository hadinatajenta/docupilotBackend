package roles

import (
	"context"
)

type Repository interface {
	GetRolesByUserID(ctx context.Context, userID string) ([]Role, error)

	// GetAll(ctx context.Context) ([]Role, error)
	// GetByID(ctx context.Context, id string) (*Role, error)
	Create(ctx context.Context, role *Role) error
	// Update(ctx context.Context, role *Role) error
	// Delete(ctx context.Context, id string) error
	// AssignPermissions(ctx context.Context, roleID string, permissionIDs []string) error
}

type Service interface {
	CreateRole(ctx context.Context, role *Role) error
}
