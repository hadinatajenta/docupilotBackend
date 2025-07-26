package users

import "time"

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	UserID    string `gorm:"type:uuid;index"`
	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

type TokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
