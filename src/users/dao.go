package users

import "time"

type User struct {
	ID          string `gorm:"primaryKey"`
	FirebaseUID string `gorm:"uniqueIndex"`
	Email       string
	Name        string
	AvatarURL   string
	CreatedAt   time.Time
	LastLogin   time.Time
	Role        string
	Password    string
}
