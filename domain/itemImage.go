package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type  ItemImage struct {
	ID           uint    `gorm:"primaryKey"`
	ItemId       uint    `gorm:"itemId"`
	Item Item

	File_name    string  `form:"file_name"`
	File_path    string  `form:"file_path"`

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}

type ItemImages []Item