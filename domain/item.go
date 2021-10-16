package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type Item struct {
	ID            uint    					`gorm:"primaryKey"`
	Name          string  					`form:"name"`
	Detail        string  					`form:"detail"`
	Normal_price  decimal.Decimal  `form:"normal_price"`
	Special_price decimal.Decimal  `form:"special_price"`
  stock         int     				  `form:"stock"`
	TypeId        uint    					`form:"typeId"`
	Like          int   					  `form:"like"`
	CategoryId    uint  					  `form:"categoryId"`
	Display_flg   int   						`form:"display_flg" gorm:"NOT NULL"`

	ItemImages []ItemImage
	
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}

type Items []Item