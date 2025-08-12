package permission

import "time"

type Permission struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:100;unique;not null" binding:"required"`
	Description string `binding:"required"`
	CreatedAt   time.Time
	Feature     string `gorm:"size:100;not null" binding:"required"`
}

type PermissionResponse struct {
	Data map[string][]Permission `json:"permissions"`
}

type GetPermissionByUidRes struct {
	Role       string       `json:"role_id"`
	Permission []Permission `json:"permissions"`
}
