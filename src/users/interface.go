package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type URepository interface {
	Create(ctx context.Context, tx *gorm.DB, user *User) error
	GetByFirebaseUID(ctx context.Context, uid string) (*User, error)
	CheckEmailExist(ctx context.Context, email string) (*User, error)
	UpdateLastLogin(ctx context.Context, tx *gorm.DB, firebaseUid string, t time.Time) error
}

type UService interface {
	GetByFirebaseUID(ctx context.Context, firebaseUID string) (*User, error)
}
