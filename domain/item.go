package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type Item struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `json:"name"`
	Detail       string  `json:"detail"`
	Like         int     `json:"like"`
	Category_id  uint    `json:"category_id"`
	Display_flg  bool    `json:"display_flg"`

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Items []Item

