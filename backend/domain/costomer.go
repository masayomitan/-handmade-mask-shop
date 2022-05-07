package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Costomer struct {
	ID        		  uint      `json:"id" gorm:"primaryKey"`
	StatusID        uint
	Password        string    `json:"password"`
	Handlename      	string		`json:"handle_name"`
	Family_name      	string		`json:"family_name"`
	First_name       	string 	`json:"first_name"`
	Family_name_kana  string	`json:"first_name_kana"`
	First_name_kana   string  `json:"family_name_kana"`
	Email           	string  `json:"email" gorm:"uniqueIndex"`
	Phone_number    	int   
	Birthday        time.Time
  SexId           uint
	PrefID          uint
	Postal_code     int
	Address_1       string
	Address_2       string
	First_Purchase_date	time.Time
	Last_Purchase_date  	time.Time
	Purchase_times  int                 `json:"purchase_times" gorm:"default:0"`
	Purchase_total  int             		`json:"purchase_total" gorm:"default:0"`
	Purchase_average    decimal.Decimal `json:"purchase_average" gorm:"default:0"`
	Point          int                  `json:"point" gorm:"default:0"`
	Reset_key      string
	Reset_expire   string
	Session_key    string
	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt  `form:"deleted_at" gorm:"default:'null'"`
	
	// Carts []Cart
	Orders []Order  `gorm:"foreignKey:CostomerID"`
	CostomerImages []CostomerImage  `gorm:"foreignKey:CostomerID"`
}

type Costomers []Costomer
