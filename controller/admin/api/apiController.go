package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
  "fmt"
	"os"
	"strconv"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"
	"handmade_mask_shop/domain"
)


var category domain.Category
var item domain.Item
var itemImage domain.ItemImage


func GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := repository.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adminUser not found"})
	}
	fmt.Println(item)
}

func PostItem(c *gin.Context) {

	err := c.Bind(&item)
  if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}
	fmt.Println(item.Detail)

	itemId, err := repository.SaveItem(&item)
	fmt.Println(itemId)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, item)
}


func PostItemImage (c * gin.Context) {
	file, _ := c.FormFile("file")

  if (file != nil) {
		newFileName := service.RenameFile(file.Filename)
		imageDir := "./public/img/"
		filePath := imageDir + newFileName

		if f, exist := os.Stat(imageDir); os.IsNotExist(exist) || 
		!f.IsDir() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "error. could not find dir",
			})
		}

		err2 := c.SaveUploadedFile(file, filePath)
		if err2 != nil {
			defer 
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
			})
			return
		}
		// service.ResizeFile(filePath)

		p := repository.ItemImage{}
		itemImage, _ := p.SaveItemImage(newFileName)
		c.JSON(http.StatusOK, itemImage.ID)
  }
}


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