package roles

import (
	"context"

	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) Repository {
	return &roleRepo{db}
}

func (r *roleRepo) GetRolesByUserID(ctx context.Context, userID string) ([]Role, error) {
	var roles []Role
	err := r.db.WithContext(ctx).
		Table("roles").
		Select("roles.*").
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *roleRepo) GetPermissionsByUserID(ctx context.Context, userID string) ([]Permission, error) {
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
