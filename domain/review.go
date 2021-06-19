package domain

import (
	_ "github.com/go-sql-driver/mysql"

	// "gorm.io/gorm"

	"time"
)

type Review struct {
	ID           uint             `gorm:"primaryKey"`
	Comment      string
	Reccomend    int

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Reviews []Review
