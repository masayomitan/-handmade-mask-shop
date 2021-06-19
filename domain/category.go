package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Category struct {
	ID           uint              `gorm:"primaryKey"`
	Name         string

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Categories []Category