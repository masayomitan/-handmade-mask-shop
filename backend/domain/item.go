package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
	"time"
	"gorm.io/gorm"
)

type Item struct {
	ID            uint    					`gorm:"primaryKey"`
	AdminUserID   uint              `gorm:"admin_user_id"`
	Item_code     int               `form:"item_code"`
	Name          string  					`form:"name"`
	Detail        string  					`form:"detail"`
	Normal_price  decimal.Decimal   `form:"normal_price"`
	Special_price decimal.Decimal   `form:"special_price"`
  Stock         int     				  `form:"stock"`
	Add_point     int               `form:"add_point"`
	Like          int   					  `form:"like"`
	Display_flg   int   						`form:"display_flg" gorm:"NOT NULL"`
	CategoryID    uint  					  `form:"category_id"`
  Orderby       uint							`form:"orderby"`

	CreatedAt    time.Time  				`form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  				`form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt 		`json:"deleted_at" gorm:"default:'null'"`

	Category Category

	ItemImages []ItemImage `gorm:"many2many:items_item_images; association_autoupdate:false"`
	CartItems  []CartItem
	OrderItems []OrderItem

}

type Items []Item