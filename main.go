package main

import (
		// "github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		// "fmt"
		// "os"
		"handmade_mask_shop/infrastructure/database"
		"handmade_mask_shop/routes"
		
)

func main() {

	db := database.GormConnect()
	database.Migrations(db)
	r := routes.GetAdminRoutes()

		files := []string{ 
			"./templates/top/index.html", "./templates/top/detail.html",
			"./templates/admin/dashboard/index.html", "./templates/admin/item/index.html", "./templates/admin/item/detail.html", "./templates/admin/item/create.html",
		}

		r.LoadHTMLFiles(files...)
		r.Static("/src", "./src")

		r.GET("/", routes.Top)
		r.GET("/detail:id", routes.TopDetail)
		

	r.Run(":80")
};