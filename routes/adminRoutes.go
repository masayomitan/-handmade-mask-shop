package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"github.com/gin-contrib/sessions"
	"handmade_mask_shop/controller/admin"
  API "handmade_mask_shop/controller/admin/api"
	"handmade_mask_shop/domain"
	"github.com/koron/go-dproxy"
	"encoding/json"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
)

var csrfMd func(http.Handler) http.Handler

func GetAdminRoutes(r *gin.Engine) *gin.Engine {
	fmt.Println()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
  })
  
	r.GET("/admin/dashboards", controller.Dashboard)
	login := r.Group("/admin/")
		{
			login.GET("/", controller.LoginTop)
			login.POST("/login", controller.Login)
			login.GET("/logout", controller.Logout)
		}

	user := r.Group("/admin/users/")
		{
			user.GET("/regist", controller.UserRegist)
			// user.GET("/detail/:id", controller.UserDetail)
		}
		
	//管理ユーザーグループ
	admin := r.Group("/admin/admin-users/", LoginCheckMiddleware())
		{	
			admin.GET("/regist", controller.AdminRegist)
			admin.GET("/edit", controller.AdminEdit)
			admin.POST("/update", controller.AdminUpdate)
		}

	//商品グループ
	item := r.Group("/admin/items/", LoginCheckMiddleware())
		{	
			item.Use(CSRF())
			item.GET("/", controller.Index)
			item.GET("/detail/:id", controller.ItemDetail)
			item.GET("/create", controller.Create)
			item.POST("/store", controller.Store)
			item.GET("/category", controller.Category)
		}

	api := r.Group("/admin/api/")
		{
			api.GET("/get-categories", API.GetCategories)
			api.POST("/post-categories", API.PostCategories)
			api.GET("/delete-category/:id", API.DeleteCategory)
		}
  return r
}

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			session := sessions.Default(c)
			var setAdminUser domain.SetAdminUser
			loginAdminUser, _ := dproxy.New(session.Get("adminUser")).String()
			err := json.Unmarshal([]byte(loginAdminUser), &setAdminUser)
			// fmt.Println(setAdminUser.ID)

			if err != nil {
					c.Redirect(http.StatusSeeOther, "/admin/")
					c.Abort()
			} else {
					c.Set("adminUser", string(loginAdminUser))
					c.Set("id", setAdminUser.ID)
					c.Next()
			}
			log.Println("session is alive")
	}
}

func CSRF() gin.HandlerFunc {
	csrfMd = csrf.Protect(
		[]byte("a-32-byte-long-key"),
		csrf.MaxAge(0),
		csrf.Secure(true),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message": "Forbidden - CSRF token invalid"}`))
		})),
	)
	return adapter.Wrap(csrfMd)
}