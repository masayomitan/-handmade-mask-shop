package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"gorm.io/gorm"
)

type Category struct {
	ID           uint    		`gorm:"primaryKey"`
	Name         string  		`json:"name"`
	Orderby      int     		`json:"orderby"`

	CreatedAt    time.Time  `json:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"default:'null'"`
	
	Items []Item
}

type Categories []Category