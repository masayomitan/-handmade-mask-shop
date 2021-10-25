package main

import (
		// "github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		// "fmt"
		// "os"
		// "handmade_mask_shop/infrastructure/database"
		"handmade_mask_shop/routes"
		
)

func main() {

	// db := database.GormConnect()
	// database.Migrations(db)
	// database.Seeds(db)
	r := routes.GetAdminRoutes()

		files := []string{ 
			"./templates/top/index.html", "./templates/top/detail.html",
			"./templates/admin/dashboard/index.html", "./templates/admin/item/index.html", "./templates/admin/item/detail.html", "./templates/admin/item/create.html",
			"./templates/admin/user/regist.html", 
			"./templates/admin/adminUser/regist.html",
		  "./templates/layout/dafault.html", "./templates/layout/admin_default.html",
			"./templates/admin/element/header.html", "./templates/admin/element/footer.html", "./templates/admin/element/sidebar.html",
		}

		r.LoadHTMLFiles(files...)
		r.Static("/src", "./src")
		r.Static("/public", "./public")

		r.GET("/", routes.Top)
		r.GET("/detail:id", routes.TopDetail)
		
	r.Run(":80")
};