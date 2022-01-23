package controller

import (
	// "fmt"
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"

	// "handmade_mask_shop/service"
	// "github.com/gin-contrib/sessions"
	"net/http"
	// "os"
	// "strconv"
	"github.com/gin-gonic/gin"
)


var item domain.Item
var itemImage domain.ItemImage
var category domain.Category


func GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := repository.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adminUser not found"})
	}
	c.JSON(http.StatusOK, item)
}


func GetDisplayItems(c *gin.Context) {

  items, err := repository.GetDisplayItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items not found"})
	}
	fmt.Println(items)
	c.JSON(http.StatusOK, items)
}

func GetItemImages (c * gin.Context) {
	itemImages, err := repository.GetAllItemImages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, itemImages)
}


func GetCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, categories)
}