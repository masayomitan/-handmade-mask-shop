package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	// "gorm.io/gorm"
	"time"
)

type User struct {
	ID        		  uint             `gorm:"primaryKey"`
	StatusID        uint

	Password        string    `gorm:"password"`
	PasswordConfirm string    `json:"password_confirm"`

	FamilyName      string
	firstName       string
	FamilyNameKana  string
	firstNameKana   string

	Email           *string
	Phone_number    int   
	Birthday        time.Time
  SexId           uint

	PrefIf          uint
	Postal_code     int
	Address_1       string
	Address_2       string

	First_buy_date time.Time
	Last_buy_date  time.Time
	Buy_times      int
	Buy_total      int
	Buy_average    decimal.Decimal
	Point          int
	
	Reset_key      string
	Reset_expire   string

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
	

	Carts []Cart
	Orders []Order
	OrderItems []OrderItem
}

type Users []User
