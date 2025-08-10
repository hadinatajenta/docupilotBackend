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

func (r *roleRepo) Create(ctx context.Context, role *Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepo) CheckRoleExist(ctx context.Context, role string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&Role{}).
		Where("name = ?", role).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
