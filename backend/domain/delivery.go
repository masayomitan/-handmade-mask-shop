package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type Delivery struct {
	ID            uint    					`gorm:"primaryKey"`
	Name          string  					`form:"name"`
	Service_name  string  					`form:"detail"`
  Description   string
  
  weight_min  decimal.Decimal
	weight_max  decimal.Decimal
	size_min    decimal.Decimal
	size_max    decimal.Decimal
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
}

type Deliveries []Delivery