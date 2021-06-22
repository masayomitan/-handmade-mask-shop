package main

import (
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"

		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		// "net/http"
		// "html/template"
		"handmade_mask_shop/routes"
		"handmade_mask_shop/domain"

)


func main() {
		r := gin.Default()
		
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
		
		r.LoadHTMLGlob("./templates/top/*")

    r.Static("/src", "./src")

		r.GET("/", routes.Top)
		r.GET("/detail", routes.TopDetail)
		
		r.GET("/admin/dashboard", routes.AdminDashboard)
		r.GET("/admin/item", routes.AdminItems)
		


	r.Run(":80")

};