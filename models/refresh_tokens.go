package models

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	Token     string    `gorm:"uniqueIndex;size:512"`
	UserID    uint
	User      User
	ExpiresAt time.Time
	CreatedAt time.Time
}
