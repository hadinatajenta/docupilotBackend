package roles

import "time"

type RolePermission struct {
	RoleID       string `gorm:"primaryKey;column:role_id"`
	PermissionID string `gorm:"primaryKey;column:permission_id"`
}

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

type assignPermissionsReq struct {
	Permissions []string `json:"permissions" binding:"required,min=1,dive,uuid4"`
}

type RolePermissionRes struct {
	RoleID      string `json:"role_id"`
	RoleName    string `json:"role_name"`
	Permissions []any  `json:"permissions"`
}

type GetNameID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
