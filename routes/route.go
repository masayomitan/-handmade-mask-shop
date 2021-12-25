package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	front "handmade_mask_shop/controller/front"
  // API "handmade_mask_shop/controller/front/api"
	"handmade_mask_shop/domain"
	"github.com/koron/go-dproxy"
	"encoding/json"
)


func GetUserRoutes(r *gin.Engine) *gin.Engine {
	fmt.Println()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
  })

	login := r.Group("/admin/")
		{
			login.GET("/", front.LoginTop)
			login.POST("/login", front.Login)
			login.GET("/logout", front.Logout)
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
