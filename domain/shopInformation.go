package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ShopInformation struct {
	ID            uint   `gorm:"primaryKey"`
	Shop_name     string `form:"shop_name"`
	Shop_kana     string `form:"shop_kana"`
	Postal_code   int    `form:"Postal_code"`
	Address_1     string `form:"address_1"`
	Address_2     string `form:"address_2"`
  Phone_number  int    `form:"Phone_number"`
	Email         string `form:"Email"`
	
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}
