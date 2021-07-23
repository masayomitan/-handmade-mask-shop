package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Item struct {
	ID           uint              `gorm:"primaryKey"`
	Name         string
	Category_id  uint
	Comment      string
	Like         int

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Items []Item

