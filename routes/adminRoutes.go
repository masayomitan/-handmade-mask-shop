package routes

import (
	"net/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	"log"
	"handmade_mask_shop/controller/admin"
	"handmade_mask_shop/domain"

	"github.com/gin-gonic/gin"
)

var adminUser domain.AdminUser

func GetAdminRoutes(r *gin.Engine) *gin.Engine {
	fmt.Println()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
  })

	r.GET("/admin/dashboard", controller.Dashboard)
	login := r.Group("/admin/login-top/")
		{
			login.GET("/", controller.LoginTop)
			login.POST("/login", controller.Login)
			login.GET("/logout", controller.Logout)
		}

	user := r.Group("/admin/user/")
		{
			user.GET("/regist", controller.UserRegist)
			user.GET("/detail/:id", controller.Detail)
		}
		
	//管理ユーザーグループ
	admin := r.Group("/admin/admin-user/")
		{	
			admin.GET("/regist", controller.AdminRegist)
			// user.GET("/detail/:id", controller.Detail)
		}

	//商品グループ
	item := r.Group("/admin/item/")
	item.Use(sessionCheck())
		{	
			item.GET("/", controller.Index)
			item.GET("/detail/:id", controller.Detail)
			item.GET("/create", controller.Create)
			item.POST("/store", controller.Store)
		}
  return r
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

			session := sessions.Default(c)
			
			id := session.Get("adminUserID")
fmt.Println(id)

			if id == nil {
          // c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
					c.Redirect(http.StatusFound, "/admin/login-top/")
					c.Abort() // これがないと続けて処理されてしまう
			} else {
					c.Set("UserId", id) // ユーザidをセット
					c.Next()
			}
			log.Println("done")
	}
}