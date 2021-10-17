package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ShopInformation struct {
	ID            uint   `gorm:"primaryKey"`
	Shop_name     string `form:"name"`
	Shop_kana     string `form:"detail"`
	Name          string `form:"name"`
	Kana          string `form:"detail"`
	Postal_code   int    `form:"normal_price"`
	Address_1     string `form:"created_at"`
	Address_2     string `form:"special_price"`
  Phone_number  int    `form:"stock"`
	Email         string `form:"typeId"`
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}
