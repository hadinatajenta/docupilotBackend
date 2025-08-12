package users

import (
	"context"
	"gobackend/shared/utils"
	"gobackend/src/roles"
	"time"

	"gorm.io/gorm"
)

type URepository interface {
	Create(ctx context.Context, tx *gorm.DB, user *User) error
	GetByUserID(ctx context.Context, uid string) (*User, error)
	CheckEmailExist(ctx context.Context, email string) (*User, error)
	UpdateLastLogin(ctx context.Context, tx *gorm.DB, firebaseUid string, t time.Time) error
	GetRolesByUserIds(ctx context.Context, ids []string) ([]*roles.Role, error)
	AssignRolesToUser(ctx context.Context, userID string, roleIDs []string) error
	CheckRolesExist(ctx context.Context, roleIDs []string) (bool, error)
	GetUsers(ctx context.Context, p utils.Params) ([]GetUsers, int, error)
	CreateUser(ctx context.Context, user *User) error
}

type UService interface {
	GetByUserID(ctx context.Context, firebaseUID string) (*User, error)
	AssignRolesToUser(ctx context.Context, userId string, roleIDs []string) (assignRoleRes, error)
	GetUsers(ctx context.Context, p utils.Params) ([]GetUsers, utils.Meta, error)
	CreateUser(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
}
