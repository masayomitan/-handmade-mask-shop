package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"gorm.io/gorm"
)

type  ItemImage struct {
	ID           uint    `gorm:"primaryKey"`
	ItemID       uint    `gorm:"item_id"`
	File_name    string  `form:"file_name"`
	File_path    string  `form:"file_path"`

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt 		`form:"deleted_at" gorm:"default:'null'"`
  
	//belongsTo
	Item Item
}

type ItemImages []Item