package users

import (
	"context"
	"errors"
	"gobackend/shared/utils"
	"gobackend/src/roles"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) URepository {
	return &userRepo{db}
}

func (r *userRepo) GetByUserID(ctx context.Context, uid string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *userRepo) Create(ctx context.Context, tx *gorm.DB, user *User) error {
	return tx.WithContext(ctx).Create(user).Error
}

func (r *userRepo) UpdateLastLogin(ctx context.Context, tx *gorm.DB, firebaseUid string, t time.Time) error {
	return tx.WithContext(ctx).
		Model(&User{}).
		Where("firebase_uid = ?", firebaseUid).
		Update("last_login", t).
		Error
}

func (r *userRepo) CheckEmailExist(ctx context.Context, email string) (*User, error) {
	var usr User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&usr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &usr, nil
}

func (r *userRepo) GetRolesByUserIds(ctx context.Context, ids []string) ([]*roles.Role, error) {
	var roles []*roles.Role
	err := r.db.WithContext(ctx).
		Table("roles").
		Select("roles.*").
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id IN ?", ids).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *userRepo) CheckRolesExist(ctx context.Context, roleIDs []string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&roles.Role{}).Where("id IN ?", roleIDs).Count(&count).Error; err != nil {
		return false, err
	}
	return count == int64(len(roleIDs)), nil

}

func (r *userRepo) AssignRolesToUser(ctx context.Context, userID string, roleIDs []string) error {
	panic("unimplemented")
}

func (r *userRepo) GetUsers(ctx context.Context, p utils.Params) ([]GetUsers, int, error) {
	var (
		users []GetUsers
		total int64
	)

	q := r.db.WithContext(ctx).Table("users")

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sortCol := map[string]string{
		"created_at": "created_at",
		"email":      "email",
		"name":       "name",
		"last_login": "last_login",
	}[p.Sort]
	if sortCol == "" {
		sortCol = "created_at"
	}

	desc := p.Order == "desc"

	err := q.
		Select(`firebase_uid, email, name, avatar_url, roles::text[] AS roles, last_login`).
		Order(clause.OrderByColumn{Column: clause.Column{Name: sortCol}, Desc: desc}).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false}).
		Limit(p.Limit).
		Offset(p.Offset).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
