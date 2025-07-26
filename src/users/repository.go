package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) URepository {
	return &userRepo{db}
}

func (r *userRepo) GetByFirebaseUID(ctx context.Context, uid string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("firebase_uid = ?", uid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, nil
}

func (r *userRepo) Create(ctx context.Context, tx *gorm.DB, user *User) error {
	return tx.WithContext(ctx).Create(user).Error
}

func (r *userRepo) UpdateLastLogin(ctx context.Context, tx *gorm.DB, firebaseUid string, t time.Time) error {
	return tx.WithContext(ctx).Model(&User{}).Where("firebase_uid = ? ", firebaseUid).Update("last_login", t).Error
}

func (r *userRepo) CheckEmailExist(ctx context.Context, email string) (*User, error) {
	var usr User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&usr).Error

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *userRepo) StoreRefreshToken(ctx context.Context, tx *gorm.DB, userId, refreshToken string, expAt time.Time) error {
	refreshtoken := RefreshToken{
		ID:        uuid.NewString(),
		UserID:    userId,
		ExpiresAt: expAt,
		CreatedAt: time.Now(),
		Token:     refreshToken,
	}
	return tx.WithContext(ctx).Create(refreshtoken).Error
}

func (r *userRepo) FindRefreshToken(ctx context.Context, refreshToken string) (*RefreshToken, error) {
	var rt RefreshToken
	if err := r.db.WithContext(ctx).Where("token = ? ", refreshToken).First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *userRepo) DeleteRefreshToken(ctx context.Context, tx *gorm.DB, refreshToken string) error {
	return r.db.WithContext(ctx).Where("token = ?", refreshToken).Delete(&RefreshToken{}).Error
}
