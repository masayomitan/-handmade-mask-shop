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
  
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Purcahses []Purchase

