package main

import (
		"github.com/gin-gonic/gin"
		_ "github.com/go-sql-driver/mysql"
		"fmt"
		// "os"
		// "handmade_mask_shop/infrastructure/database"
		"handmade_mask_shop/routes"
		"github.com/gin-contrib/sessions"
		"github.com/gin-contrib/sessions/cookie"
)

func main() {
	fmt.Println()
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessionID", store))

	r = routes.GetAdminRoutes(r)
	r = routes.GetRoutes(r)

	// db := database.GormConnect()
	// database.Migrations(db)
	// database.Seeds(db)


		files := []string{ 
			"./templates/top/index.html", "./templates/front/item/detail.html",

			"./templates/admin/dashboards/index.html", "./templates/admin/items/index.html", "./templates/admin/items/detail.html", "./templates/admin/items/create.html",
			"./templates/admin/users/regist.html", 
			"./templates/admin/adminUsers/regist.html",
      "./templates/layout/dafault.html", "./templates/layout/admin_default.html",
			"./templates/admin/elements/header.html", "./templates/admin/elements/footer.html", "./templates/admin/elements/sidebar.html",
			"./templates/admin/login/index.html",
		}

		r.LoadHTMLFiles(files...)
		r.Static("/src", "./src")
		r.Static("/public", "./public")
		
	r.Run(":80")
};