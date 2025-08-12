package users

import (
	"time"

	"github.com/lib/pq"
)

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type RefreshToken struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	UserID    string `gorm:"type:uuid;index"`
	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

type UpdateProfile struct {
	Name     string `json:"name" gorm:"column:name" binding:"required"`
	Email    string `json:"email" gorm:"column:email" binding:"required"`
	Password string `json:"password" gorm:"column:password" binding:"required"`
}

type assignRolesReq struct {
	RolesIDS []string `json:"role_ids" binding:"required,min=1,dive,uuid4"`
}

type assignRoleRes struct {
	RoleID      string `json:"user_id"`
	RoleName    string `json:"user_name"`
	Permissions []any  `json:"roles"`
}

type GetRoles struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetUsers struct {
	FirebaseUid string         `json:"firebase_uid"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	AvatarUrl   string         `json:"avatar_url"`
	Roles       pq.StringArray `gorm:"type:text[]" json:"roles"`
	LastLogin   time.Time      `json:"last_login"`
}

type CreateUserRequest struct {
	Email    string   `json:"email" binding:"required,email"`
	Roles    []string `json:"roles" binding:"required"`
	Password string   `json:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	CreatedAt string   `json:"created_at"`
}
