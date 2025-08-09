package roles

import "time"

type Role struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:50;unique;not null"`
	Description string
	CreatedAt   time.Time
}

type Permission struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:100;unique;not null"`
	Description string
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
