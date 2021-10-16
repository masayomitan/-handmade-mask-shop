package controller

import (
	"github.com/gin-gonic/gin"
	// "handmade_mask_shop/domain"
	// "handmade_mask_shop/infrastructure/database"
	"net/http"
  // "fmt"
	_ "github.com/go-sql-driver/mysql"

)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/dashboard", gin.H{
		"param1": "ダッシュボードです",
    "param2": "layoutです",
	})
}

