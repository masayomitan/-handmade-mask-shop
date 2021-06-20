package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Contact struct {
	ID           uint              `gorm:"primaryKey"`
	User_id      uint
	Text         string

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Contacts []Contact