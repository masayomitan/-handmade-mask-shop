package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"handmade_mask_shop/domain"
	// "handmade_mask_shop/infrastructure/seed"

)

func Migrations(db *gorm.DB) {
	// db.AutoMigrate(&domain.Cart{})
	// db.AutoMigrate(&domain.CartItem{})
	// db.AutoMigrate(&domain.Category{})
	// db.AutoMigrate(&domain.Contact{})
	// db.AutoMigrate(&domain.Delivery{} )
	// db.AutoMigrate(&domain.Item{})
	// db.AutoMigrate(&domain.ItemImage{})
	// db.AutoMigrate(&domain.ItemsItemImage{})
	// db.AutoMigrate(&domain.Order{})
	// db.AutoMigrate(&domain.OrderItem{})
	// db.AutoMigrate(&domain.Review{})
	// db.AutoMigrate(&domain.ShopInformation{})
	db.AutoMigrate(&domain.Costomer{})
	// db.AutoMigrate(&domain.CostomerImage{})
	// db.AutoMigrate(&domain.AdminUser{})
	// db.AutoMigrate(&domain.UserImage{})
}

func Seeds(db *gorm.DB) {
  // seed.CategorySeed(db)
	// seed.AdminUserSeed(db)
}