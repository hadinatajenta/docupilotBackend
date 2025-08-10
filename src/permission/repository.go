package permission

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionChecker {
	return &repository{
		db: db,
	}
}

func (r *repository) GetPermissionsByUserID(ctx context.Context, userID string) ([]Permission, error) {
	var perms []Permission
	err := r.db.WithContext(ctx).
		Table("permissions").
		Select("permissions.*").
		Joins("JOIN role_permissions rp ON rp.permission_id = permissions.id").
		Joins("JOIN user_roles ur ON ur.role_id = rp.role_id").
		Where("ur.user_id = ?", userID).
		Find(&perms).Error
	return perms, err
}

func (r *repository) CreatePermission(ctx context.Context, permission *Permission) error {
	return r.db.WithContext(ctx).Create(permission).Error
}
