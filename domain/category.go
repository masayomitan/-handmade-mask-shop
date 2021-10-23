package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Category struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string  `form:"name"`
	Orderby      int

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
	
	Items []Item
}

type Categories []Category