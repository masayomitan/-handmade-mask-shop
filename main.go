package main

import (
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		// "fmt"
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		"handmade_mask_shop/routes"
		"handmade_mask_shop/domain"
)

func dbInsert(text string, status string) {
	dsn := "root:@tcp(127.0.0.1:3306)/handmade_db?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			panic("データベース開けず！（dbInsert)")
	}
	db.Create(&Item{Name: name})
	defer db.Close()
}

func main() {
		r := gin.Default()
		
		//if we not found route
		r.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
		
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
		
		files := []string{ 
			"./templates/top/index.html", "./templates/top/detail.html",
			"./templates/admin/dashboard/index.html", "./templates/admin/item/index.html", "./templates/admin/item/detail.html", "./templates/admin/item/create.html",
		}

		r.LoadHTMLFiles(files...)
		r.Static("/src", "./src")

		r.GET("/", routes.Top)
		r.GET("/detail:id", routes.TopDetail)
		r.GET("/admin/dashboard", routes.AdminDashboard)
		
		v1 := r.Group("/admin/")
		{	
			v1.GET("/item", routes.AdminItem)
			v1.GET("/item/detail/:id", routes.AdminItemDetail)
			v1.GET("/item/create", routes.AdminItemCreate)
			v1.POST("/item/store", routes.AdminItemStore)
		}
	r.Run(":80")

};