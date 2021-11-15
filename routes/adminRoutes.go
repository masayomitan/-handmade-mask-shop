package routes

import (
	// "net/http"
  "fmt"
	// "log"
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/controller/admin"
	"github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
)

func GetAdminRoutes(r *gin.Engine) *gin.Engine {
	fmt.Println()
	
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

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
		{	
			item.GET("/", controller.Index)
			item.GET("/detail/:id", controller.Detail)
			item.GET("/create", controller.Create)
			item.POST("/store", controller.Store)
		}
  return r
}

// func sessionCheck() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 			session := sessions.Default(c)
// 			LoginInfo.UserId = session.Get("UserId")

// 			// セッションがない場合、ログインフォームをだす
// 			if LoginInfo.UserId == nil {
// 					log.Println("ログインしていません")
// 					c.Redirect(http.StatusFound, "/login")
// 					c.Abort() // これがないと続けて処理されてしまう
// 			} else {
// 					c.Set("UserId", LoginInfo.UserId) // ユーザidをセット
// 					c.Next()
// 			}
// 			log.Println("ログインチェック終わり")
// 	}
// }