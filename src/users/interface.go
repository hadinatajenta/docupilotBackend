package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type URepository interface {
	GetByFirebaseUID(ctx context.Context, uid string) (*User, error)
	Create(ctx context.Context, tx *gorm.DB, user *User) error
	UpdateLastLogin(ctx context.Context, tx *gorm.DB, firebaseUid string, t time.Time) error
	CheckEmailExist(ctx context.Context, email string) (*User, error)
	StoreRefreshToken(ctx context.Context, tx *gorm.DB, userId, refreshToken string, expAt time.Time) error
	FindRefreshToken(ctx context.Context, refreshToken string) (*RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, tx *gorm.DB, refreshToken string) error
	// UpdateProfile(ctx context.Context, id string) error
}

type UService interface {
	SyncFirebaseUser(ctx context.Context, uid, email, name, avatarURL string) (*User, error)
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, refreshToken string) error
	GetByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error)
	// UpdateProfile(ctx context.Context, id string) error
}
