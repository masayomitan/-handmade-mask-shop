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


func UserIndex (c *gin.Context) {
	action := "userIndex"
	c.HTML(http.StatusOK, "admin/users/index.html", gin.H{
		"action" :action,
	})
}


func UserRegist(c *gin.Context) {
	fmt.Println()
  c.HTML(http.StatusOK, "admin/users/regist.html", gin.H{
	})
}