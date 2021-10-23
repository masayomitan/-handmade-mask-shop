package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type OrderItem struct {
	ID           uint `gorm:"primaryKey"`
	UserId       uint `gorm:"name"`
  OrderId      uint `gorm:"order_id"`
	PaymentId    uint `gorm:"payment_id"`

  ItemId       uint   `gorm:"item_id"`
	Item_code    string `json:"item_code"`
	Item_name    string `json:"item_name"`
	CategoryID   uint   `gorm:"category_id"`

	Sub_total    decimal.Decimal `json:"sub_total"`
	Discount     decimal.Decimal `json:"discount"`
	Delivery_fee decimal.Decimal `json:"delivary_fee"`
	Tax          decimal.Decimal `json:"tax"`

	Price        decimal.Decimal `json:"price"`
	Quantity     int  					 `json:"quantity"`
	Add_point    int   					 `json:"add_point"`
	
	Order_date   time.Time  `form:"order_date" gorm:"NOT NULL"`
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`


	User User
	Item Item
	Order Order
	
}

type OrderItems []OrderItem