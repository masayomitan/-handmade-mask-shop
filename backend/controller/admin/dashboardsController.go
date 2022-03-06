package controller

import (
	"github.com/gin-gonic/gin"
	// "handmade_mask_shop/domain"
	// "handmade_mask_shop/infrastructure/database"
	"net/http"
  "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-contrib/sessions"

)


func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	action := "dashboard"
	
	fmt.Println(session.Get("adminUser"))

	c.HTML(http.StatusOK, "admin/dashboards", gin.H{
		"param1": "ダッシュボードです",
    "param2": "layoutです",
		"action": action,
	})
}

