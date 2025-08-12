package users

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID          string `gorm:"primaryKey"`
	FirebaseUID string `gorm:"uniqueIndex"`
	Email       string
	Name        string
	AvatarURL   string
	CreatedAt   time.Time
	LastLogin   time.Time
	Roles       pq.StringArray `json:"roles" gorm:"type:text[]"`
	Password    string         `json:"-"`
}

func (User) TableName() string {
	return "users"
}
