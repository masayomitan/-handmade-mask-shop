package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Item struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `form:"name"`
	Detail       string  `form:"detail"`
	Like         int     `form:"like"`
	CategoryId  uint     `form:"categoryId"`
	Display_flg  bool    `form:"display_flg"`

	ItemImages []ItemImage 
	

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}

type Items []Item

