package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Item struct {
	ID           uint              `gorm:"primaryKey"`
	Name         string
	category_id  uint
	comment      string
	like         int

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Items []Item

