package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "fmt"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/domain"
)

var category domain.Category

func GetCategories(c *gin.Context) {
  fmt.Println()
	data, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, data)
}


func PostCategories(c *gin.Context) {
  err :=  c.Bind(&category)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

  //if record alredy found then not save at all
  exists := repository.CheckExistsByCategoryName(category.Name)
	if exists == true {
		fmt.Println("record already exists")
		c.String(http.StatusBadRequest, "record already exists: "+err.Error())
		return
	}
	
  _, err = repository.SaveCategory(&category)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteCategory(c *gin.Context) {
	err := c.Bind(&category)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

  err = repository.DeleteCategory(category.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}