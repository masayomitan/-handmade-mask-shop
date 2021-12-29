package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"gorm.io/gorm"
)

type  ItemImage struct {
	ID           uint    	`gorm:"primaryKey"`
	UserID       uint    	`json:"user_id"`
	File_name    string  	`json:"file_name"`
	File_path    string  	`json:"file_path"`
	Orderby      int     	`json:"orderby"`


	CreatedAt    time.Time  `json:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt  `form:"deleted_at" gorm:"default:'null'"`

}

type ItemImages []ItemImage