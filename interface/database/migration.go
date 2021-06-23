package migration

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"handmade_mask_shop/domain"

)

func Migration(c *gin.Context) {
		
		dsn := "root:@tcp(127.0.0.1:3306)/handmade_db?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic ("データベースとの通信に失敗しました。")
		}
		// Migrate the schema
		db.AutoMigrate( &domain.User{} )
		db.AutoMigrate( &domain.Item{})
		db.AutoMigrate( &domain.Category{})
		db.AutoMigrate( &domain.Purchase{})
		db.AutoMigrate( &domain.Review{})
		db.AutoMigrate( &domain.Contact{})
}