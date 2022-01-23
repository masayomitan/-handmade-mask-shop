package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type Cart struct {
	ID            uint 						`gorm:"primaryKey"`
	UserID      	uint				   	`gorm:"user_id"`

	Pre_orderID    string  				`json:"pre_order_id"`
	Total_price   decimal.Decimal `json:"total_price"`
	Add_point     decimal.Decimal `json:"add_point"`
  Use_point     int          		`json:"use_point"`
	Delivery_fee  decimal.Decimal `json:"delivary_fee"`

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`

	//belongsTo
	User User

	//hasMany
	CartItems []CartItem
}

type Carts []Cart