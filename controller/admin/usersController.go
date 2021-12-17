package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "handmade_mask_shop/domain"
	// "handmade_mask_shop/repository"
	// "handmade_mask_shop/service"
	"net/http"
	// "os"
	_ "github.com/go-sql-driver/mysql"

)


func UserRegist(c *gin.Context) {
	fmt.Println()
  c.HTML(http.StatusOK, "admin/user/regist.html", gin.H{
	})
}