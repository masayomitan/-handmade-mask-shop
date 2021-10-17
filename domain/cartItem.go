package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type CartItem struct {
	ID            uint    				`gorm:"primaryKey"`
	ItemID        uint            `gorm:"item_id"`
	Item_code     int           	`json:"item_code"`
	CartID        uint    				`gorm:"cart_id"`
  Price         decimal.Decimal `json:"price"`
	Quantity      int             `json:"quantity"`
	Point         int             `json:"point"`
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`

	//belongsTo
	Item Item
	Cart Cart
}

type CartItems []CartItem