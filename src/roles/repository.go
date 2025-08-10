package roles

import (
	"context"
	"gobackend/src/permission"

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

func (r *roleRepo) GetByID(ctx context.Context, id string) (*Role, error) {
	var role Role
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) PermissionsExist(ctx context.Context, ids []string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&permission.Permission{}).Where("id IN ?", ids).Count(&count).Error; err != nil {
		return false, err
	}
	return count == int64(len(ids)), nil
}

func (r *roleRepo) AssignPermissions(ctx context.Context, roleID string, permissionIDs []string) error {
	tx := r.db.WithContext(ctx).Begin()

	// Clear old permissions
	if err := tx.Where("role_id = ?", roleID).Delete(&RolePermission{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert new
	var rp []RolePermission
	for _, pid := range permissionIDs {
		rp = append(rp, RolePermission{
			RoleID:       roleID,
			PermissionID: pid,
		})
	}
	if err := tx.Create(&rp).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *roleRepo) GetPermissionByIDS(ctx context.Context, ids []string) ([]GetNameID, error) {
	var permissions []GetNameID
	err := r.db.WithContext(ctx).
		Table("permissions").
		Select("id, name").
		Where("id IN ?", ids).
		Find(&permissions).Error
	return permissions, err
}
