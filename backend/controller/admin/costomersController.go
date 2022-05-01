package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	// "handmade_mask_shop/service"
	"net/http"
	// "os"
	_ "github.com/go-sql-driver/mysql"
)


func CostomerIndex (c *gin.Context) {
	action := "costomers"
	costomers, err := repository.GetCostomers()
	fmt.Println(costomers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.HTML(http.StatusOK, "admin/costomers/index.html", gin.H{
		"costomers" :costomers,
		"action" :action,
	})
}