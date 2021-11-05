package routes

import (
    // "net/http"
    "github.com/gin-gonic/gin"
    "handmade_mask_shop/controller/front"
)


func GetRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/", controller.Top)
    r.GET("/item/detail/:id", controller.Detail)
    return r
}