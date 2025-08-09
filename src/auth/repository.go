package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Repository {
	return &authRepo{db}
}

func (r *authRepo) StoreRefreshToken(ctx context.Context, tx *gorm.DB, userId, refreshToken string, expAt time.Time) error {
	rt := RefreshToken{
		ID:        uuid.NewString(),
		UserID:    userId,
		ExpiresAt: expAt,
		CreatedAt: time.Now(),
		Token:     refreshToken,
	}
	return tx.WithContext(ctx).Create(&rt).Error
}

func (r *authRepo) FindRefreshToken(ctx context.Context, refreshToken string) (*RefreshToken, error) {
	var rt RefreshToken
	if err := r.db.WithContext(ctx).Where("token = ?", refreshToken).First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *authRepo) DeleteRefreshToken(ctx context.Context, tx *gorm.DB, refreshToken string) error {
	return tx.WithContext(ctx).Where("token = ?", refreshToken).Delete(&RefreshToken{}).Error
}
