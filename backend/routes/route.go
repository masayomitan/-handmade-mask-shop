package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	front "handmade_mask_shop/controller/front"
  API "handmade_mask_shop/controller/front/api"
	"handmade_mask_shop/domain"
	"github.com/koron/go-dproxy"
	"encoding/json"
)


func GetRoutes(r *gin.Engine) *gin.Engine {
	fmt.Println()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
  })

	login := r.Group("/")
		{
			login.POST("/login", front.Login)
			login.GET("/logout", front.Logout)
		}

	api := r.Group("/front/api/")
		{
			api.GET("/get-display-items", API.GetDisplayItems)
			api.GET("/get-display-item/:id", API.GetDisplayItem)
			api.GET("/get-display-items-category/:id", API.GetDisplayItemsCategoryId)
			api.GET("/get-item-images", API.GetItemImages)

			api.GET("/get-categories", API.GetCategories)

			api.POST("/add-order", API.AddOrder)
		}
  return r
}


func LoginUserCheckMiddleware() gin.HandlerFunc {
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
	}
}
