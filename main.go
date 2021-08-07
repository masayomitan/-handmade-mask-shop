package main

import (
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		// "fmt"
		"handmade_mask_shop/infrastructure/database"
		"handmade_mask_shop/routes"
)

func main() {
		r := gin.Default()
		
		//if we not found route
		r.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
		
	db := database.GormConnect()
	database.Migrations(db)
		
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