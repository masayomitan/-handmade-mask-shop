package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Contact struct {
	ID           uint              `gorm:"primaryKey"`
	User_id      uint
	Text         string

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}

type Contacts []Contact