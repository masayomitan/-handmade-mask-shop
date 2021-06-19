package main

import (
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"

		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		"net/http"

		"handmade_mask_shop/domain"

)		
type User struct {
		gorm.Model
		ID     uint
		Name         string
		Email        *string
		Age          uint8
}

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
		
		r.LoadHTMLGlob("./templates/*")

		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
		})
	})

		r.GET("/detail", func(c *gin.Context) {
			c.HTML(http.StatusOK, "detail.html", gin.H{
				
			})
	})

	r.Run(":80")
};