package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "handmade_mask_shop/domain"
	// "handmade_mask_shop/repository"
	// "handmade_mask_shop/service"
	"net/http"
	// "os"
	// _ "github.com/go-sql-driver/mysql"

)

func Top(c  *gin.Context) {
	fmt.Println()

	c.HTML(http.StatusOK, "top/index.html", gin.H{})
}