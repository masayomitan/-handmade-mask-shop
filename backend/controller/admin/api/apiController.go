package controller

import (
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"
	"github.com/gin-contrib/sessions"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	// "encoding/json"
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


func PostItem(c *gin.Context) {
	userId := sessions.Default(c).Get("id").(uint)
	err := c.Bind(&item)
  if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}
	
	item.AdminUserID = userId
	item, err2 := repository.SaveItem(&item)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	//making insert into JOIN Table
	var imageIds []string
	for i := 1; i < 6; i++ {
		if c.PostForm("imageId" + strconv.Itoa(i)) != "" {
			imageId := append(imageIds, c.PostForm("imageId" + strconv.Itoa(i)))
			//appendで再生製したsliceを入れ直す
			imageIds = imageId
		}
  }
	if len(imageIds) > 0 {
		err3 := service.SetItemImageIds(item.ID, imageIds)
		if err3 != nil {
			fmt.Println(err3)
			c.String(http.StatusBadRequest, "Request is failed: "+err3.Error())
			return
		}
	}
	c.JSON(http.StatusOK, item)
}


func UpdateItem(c * gin.Context) {
	fmt.Println()
	fmt.Println(c.Param("id"))
}

func GetItemImages (c * gin.Context) {
	itemImages, err := repository.GetAllItemImages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, itemImages)
}


func PostItemImage (c * gin.Context) {
	userId := sessions.Default(c).Get("id").(uint)
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
		itemImage, _ := p.SaveItemImage(userId, newFileName)
		c.JSON(http.StatusOK, itemImage)
  }
}


func GetCategories(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}
  //Not save at all if record alredy found 
  exists := repository.CheckExistsByCategoryName(category.Name)
	if exists == true {
		fmt.Println("record already exists")
		c.JSON(http.StatusBadRequest, "record already exists: "+err.Error())
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
		c.JSON(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}
  //Not save at all if record alredy found 
	exists := repository.CheckExistsByCategoryName(category.Name)
	if exists == true {
		fmt.Println("record already exists")
		c.JSON(http.StatusBadRequest, "record already exists: "+err.Error())
		return
	}
	id := c.Param("id")
	if id != "" {
		u64, _ := strconv.ParseUint(id, 10, 32)
		_, err := repository.UpdateCategory(uint(u64), &category)
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