package users

import (
	"context"
	"time"

	"gorm.io/gorm"
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
		return nil, err
	}
	return &usr, nil
}
