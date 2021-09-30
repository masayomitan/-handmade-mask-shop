package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"handmade_mask_shop/domain"

)

func Migrations(db *gorm.DB) {

		db.AutoMigrate(&domain.User{} )
		db.AutoMigrate(&domain.Item{})
		db.AutoMigrate(&domain.ItemImage{})
		db.AutoMigrate(&domain.Category{})
		db.AutoMigrate(&domain.Purchase{})
		db.AutoMigrate(&domain.Review{})
		db.AutoMigrate(&domain.Contact{})
}
