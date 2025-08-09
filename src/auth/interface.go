package auth

import (
	"context"
	"gobackend/src/users"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	StoreRefreshToken(ctx context.Context, tx *gorm.DB, userId, refreshToken string, expAt time.Time) error
	FindRefreshToken(ctx context.Context, refreshToken string) (*RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, tx *gorm.DB, refreshToken string) error
}
type Service interface {
	SyncFirebaseUser(ctx context.Context, uid, email, name, avatarURL string) (*users.User, error)
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, refreshToken string) error
}
