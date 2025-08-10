package roles

import "time"

type Role struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:50;unique;not null" binding:"required"`
	Description string `gorm:"size:100" binding:"required"`
	CreatedAt   time.Time
}

type UserRole struct {
	UserID string `gorm:"type:uuid;primaryKey"`
	RoleID string `gorm:"type:uuid;primaryKey"`
}

type RolePermission struct {
	RoleID       string `gorm:"type:uuid;primaryKey"`
	PermissionID string `gorm:"type:uuid;primaryKey"`
}
