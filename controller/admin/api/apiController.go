package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
	"handmade_mask_shop/infrastructure/database"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/domain"
)


func GetCategories(c *gin.Context) {
  fmt.Println("ok")
	data, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return 
	}
	c.JSON(http.StatusOK, data)
}

func PostCategories(c *gin.Context) {
	// fmt.Println(name)
	var category domain.Category
  err :=  c.Bind(&category)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}
	db := database.GormConnect()
	now := time.Now()
  category.CreatedAt = now
  category.UpdatedAt = now

	db.Create(&category)
	fmt.Println(&category)
	
}