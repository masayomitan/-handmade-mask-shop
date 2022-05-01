package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	admin "handmade_mask_shop/controller/admin"
  API "handmade_mask_shop/controller/admin/api"
	"handmade_mask_shop/domain"
	"github.com/koron/go-dproxy"
	"encoding/json"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
)

var csrfMd func(http.Handler) http.Handler

func GetAdminRoutes(r *gin.Engine) *gin.Engine{

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
  })
  
	r.GET("/admin/dashboards", admin.Dashboard)

	login := r.Group("/admin/")
		{
			login.GET("/", admin.LoginTop)
			login.POST("/login", admin.Login)
			login.GET("/logout", admin.Logout)
			login.GET("/reset-password", admin.ResetPassword)
			login.POST("/send-email-reset-password", admin.SendEmailResetPassword)
			login.GET("/reset-password-complete", admin.ResetPasswordComplete)
		}

	user := r.Group("/admin/costomers/")
		{
			user.GET("/", admin.CostomerIndex)
		}
		
	//管理ユーザーグループ
	adminUser := r.Group("/admin/admin-users/", LoginCheckMiddleware())
		{	
			adminUser.GET("/regist", admin.AdminRegist)
			adminUser.GET("/edit", admin.AdminEdit)
			adminUser.POST("/update", admin.AdminUpdate)
		}

	//商品グループ
	item := r.Group("/admin/items/", LoginCheckMiddleware())
		{	
			item.Use(CSRF())
			item.GET("/", admin.ItemIndex)
			item.GET("/detail/:id", admin.ItemDetail)
			item.GET("/create", admin.ItemCreate)
			item.GET("/edit/:id", admin.ItemEdit)
			item.GET("/complete", admin.Complete)
			item.GET("/category", admin.Category)

		}

	api := r.Group("/admin/api/")
		{
			api.POST("/post-item", API.PostItem)
			api.POST("/update-item/:id", API.UpdateItem)
			api.GET("/get-item-images", API.GetItemImages)
			api.POST("/post-item-image", API.PostItemImage)

			api.GET("/get-categories", API.GetCategories)
			api.POST("/post-category", API.PostCategory)
			api.POST("/update-category/:id", API.UpdateCategory)
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
			fmt.Println(setAdminUser.ID)

			if err != nil {
					c.Redirect(http.StatusSeeOther, "/admin/")
					c.Abort()
			} else {
					c.Set("adminUser", string(loginAdminUser))
					c.Set("id", setAdminUser.ID)
					c.Next()
			}
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