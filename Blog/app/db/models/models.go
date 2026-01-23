package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Password  string
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

type Blog struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Body        string
	PostedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UserID      uint      `gorm:"foreignKey"`
	Likes       int       `gorm:"default:0"`
	Comments    int       `gorm:"default:0"`
	Views       int       `gorm:"default:0"`

	User User
}

type Like struct {
	ID      uint      `gorm:"primaryKey"`
	UserID  uint      `gorm:"foreignKey"`
	LikedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`

	User User
}

type Comment struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"foreignKey"`
	CommentedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CommentText string

	User User
}
