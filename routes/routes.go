package routes

import (
    // "net/http"
    "github.com/gin-gonic/gin"
    "handmade_mask_shop/controller/front"
)

// func Top(c *gin.Context) {
//     c.HTML(http.StatusOK, "top/index.html", gin.H{})
// }

// func TopDetail(c *gin.Context){
// 	c.HTML(http.StatusOK, "top/detail.html", gin.H{})
// }

func GetRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/", controller.Top)
    r.GET("/item/detail/:id", controller.Detail)
    return r
}