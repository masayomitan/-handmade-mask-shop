package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "fmt"
	"strconv"
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


func PostCategory(c *gin.Context) {
  err :=  c.ShouldBindJSON(&category)
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
	c.JSON(http.StatusOK, category)
}


func UpdateCategory(c *gin.Context) {
	err :=  c.ShouldBindJSON(&category)
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
	id := c.Param("id")
	if id != "" {
		u64, _ := strconv.ParseUint(id, 10, 32)
		id := uint(u64)
			_, err := repository.UpdateCategory(id, &category)
			if err != nil {
				fmt.Println(err)
				return
			}
		return 
	}
}


func DeleteCategory(c *gin.Context) {
	u64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(u64)
  _, err := repository.FindCategoryByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
  err2 := repository.DeleteCategory(id)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}