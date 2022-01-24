package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"gorm.io/gorm"
)

type ItemsItemImage struct {
	ID            uint    					`gorm:"primaryKey"`
	ItemID        uint              `gorm:"item_id"`
	ItemImageID   uint              `gorm:"item_image_id"`

	CreatedAt    time.Time  				`form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  				`form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt 		`form:"deleted_at" gorm:"default:'null'"`
}
