package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Purchase struct {

	ID           uint             `gorm:"primaryKey"`
	Item_id      uint
	User_id      uint
	Item_name    string
  
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}

type Purcahses []Purchase