package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	// "handmade_mask_shop/domain"

)

func Migrations(db *gorm.DB) {

	// db.AutoMigrate(&domain.Cart{})
	// db.AutoMigrate(&domain.CartItem{})
	// db.AutoMigrate(&domain.Category{})
	// db.AutoMigrate(&domain.Contact{})
	// db.AutoMigrate(&domain.Delivery{} )
	// db.AutoMigrate(&domain.Item{})
	// db.AutoMigrate(&domain.ItemImage{})
	// db.AutoMigrate(&domain.Order{})
	// db.AutoMigrate(&domain.OrderItem{})
	// db.AutoMigrate(&domain.Review{})
	// db.AutoMigrate(&domain.ShopInformation{})
	// db.AutoMigrate(&domain.User{} )
}
