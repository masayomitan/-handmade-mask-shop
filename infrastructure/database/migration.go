package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"handmade_mask_shop/domain"

)

var db *gorm.DB

func Migrations( db *gorm.DB ) {

		db.AutoMigrate( &domain.User{} )
		db.AutoMigrate( &domain.Item{})
		db.AutoMigrate( &domain.Category{})
		db.AutoMigrate( &domain.Purchase{})
		db.AutoMigrate( &domain.Review{})
		db.AutoMigrate( &domain.Contact{})
}


// func dbInsert(text string, status string) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/handmade_db?charset=utf8mb4&parseTime=True&loc=Local"
// 		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 			panic("データベース開けず！（dbInsert)")
// 	}
// 	// db.Create(&Item{Name: name})
// 	// defer db.Close()
// }