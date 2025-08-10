package permission

import "time"

type Permission struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:100;unique;not null"`
	Description string
	CreatedAt   time.Time
}
