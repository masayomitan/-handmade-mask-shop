package routes

import (
	// "net/http"
  // "fmt"
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/controller/admin"
)

func GetAdminRoutes() *gin.Engine {

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
  })

	r.GET("/admin/dashboard", controller.Dashboard)
	
	//商品グループ
	v1 := r.Group("/admin/item/")
		{	
			v1.GET("/", controller.Index)
			v1.GET("/detail/:id", controller.Detail)
			v1.GET("/create", controller.Create)
			v1.POST("/store", controller.Store)
		}
  return r
}
