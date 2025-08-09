package menus

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetMenuByRole(ctx context.Context, role string) ([]Menus, error) {
	var menus []Menus
	err := r.db.WithContext(ctx).
		Model(&Menu{}).
		Where("? = ANY(user_roles)", role).
		Scan(&menus).Error

	if err != nil {
		return nil, err
	}

	return menus, nil
}
