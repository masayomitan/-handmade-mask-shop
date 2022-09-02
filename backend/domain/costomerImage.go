package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

type  CostomerImage struct {
	ID           uint       `gorm:"primaryKey"`
	CostomerID   uint       `gorm:"costomer_id"`
	File_name    string  		`json:"file_name"`
	File_path    string  		`json:"file_path"`

	CreatedAt    time.Time  `json:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt  `json:"deleted_at" gorm:"default:'null'"`
  
	Costomer Costomer
}

type CostomerImages []CostomerImage
