package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

type CostomerAddress struct {
	ID        		  uint      `json:"id" gorm:"primaryKey"`
	CostomerID   		uint      `gorm:"costomer_id"`
	StatusID        uint
	PrefID          uint
	Postal_code     int
	Address_1       string
	Address_2       string
	CreatedAt   		time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt       time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt       gorm.DeletedAt  `form:"deleted_at" gorm:"default:'null'"`
	
	Costomer Costomer
}

type CostomerAdrdesses []CostomerAddress
