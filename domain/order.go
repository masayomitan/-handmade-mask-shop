package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
)

type Order struct {
	ID              uint     `gorm:"primaryKey"`
	UserID          uint   `gorm:"user_id"`
	CartID          uint   `gorm:"cart_id"`

	PaymentID        uint    `json:"payment_id"`
	PreOrderID       string  `json:"pre_order_id"`
	Message          string  `json:"message"`
	Family_name      string  `json:"family_name"`
	First_name       string  `json:"first_name"`
	Family_name_kana string  `json:"family_name_kana"`
	First_name_kana  string  `json:"first_name_kana"`

	Email            *string   `json:"email"`
	Phone_number     int       `json:"phone_number"`
	Birthday         time.Time `json:"birthday"`
  SexId            uint      `json:"sex_id"`

	PrefIf           uint      `json:"pref_id"`
	Postal_code      int       `json:"postal_code"`
	Address_1        string    `json:"address_1"`
	Address_2        string    `json:"address_2"`

	SubTotal     decimal.Decimal `json:"sub_total"`
	Discount     decimal.Decimal `json:"discount"`
	Delivery_fee decimal.Decimal `json:"delivery_fee"`
	Tax          decimal.Decimal `json:"tax"`

	Total_price    decimal.Decimal `json:"total_price"`
	Total_Quantity int 						 `json:"total_quantity"`
	Currency_code  string  				 `json:"currency_code"`

	Add_point      int   `json:"add_point"`
  Use_point      int   `json:"use_point"`
	
	Order_date   time.Time  `form:"order_date" gorm:"NOT NULL"`
	
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`


	User User
	OrderItems []OrderItem


}

type Orders []Order