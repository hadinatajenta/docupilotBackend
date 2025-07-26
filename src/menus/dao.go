package menus

import "time"

type Menu struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Text      string    `gorm:"type:varchar(100);not null" json:"text"`
	Icon      string    `gorm:"type:varchar(100)" json:"icon"`
	Path      string    `gorm:"type:varchar(255);not null" json:"path"`
	UserRoles []string  `gorm:"type:text[];not null" json:"user_roles"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Menu) TableName() string {
	return "menus"
}
