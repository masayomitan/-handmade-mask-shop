package controller

import (
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"
  "fmt"
	// "github.com/gin-contrib/sessions"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
)


func Index(c *gin.Context) {
	data := repository.GetAllItems()
	c.HTML(http.StatusOK, "admin/items/index.html", gin.H{
		"data" : data,
	})
}


func Create(c *gin.Context) {
	csrf := csrf.Token
	// session := sessions.Default(c)
	fmt.Println()

  categories := service.GetJsonAllCategories()
  c.HTML(http.StatusOK, "admin/items/create.html", gin.H{
		"categories": categories,
		"CSRFField" :  csrf,
	})
}


func Store(c *gin.Context) {
	var item domain.Item
	err := c.Bind(&item)
  if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

	itemId := repository.SaveItem(&item)
	file, _ := c.FormFile("image")
	newFileName := service.RenameFile(file.Filename)
	imageDir := "./public/images/" 
	filePath := imageDir + newFileName

	if f, exist := os.Stat(imageDir); os.IsNotExist(exist) || 
	!f.IsDir() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "message": "error. could not find dir",
    })
  }

	er := c.SaveUploadedFile(file, filePath)
	if er != nil {
		defer 
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
		})
		return
  }
  
	// service.ResizeFile(filePath)
	repository.SaveItemImage(newFileName, itemId)
}


func ItemDetail(c *gin.Context) {
	id := c.Param("id")
	item := repository.GetItemByID(id)
	c.HTML(http.StatusOK, "admin/items/detail.html", gin.H{
		"item" : item,
	})
}


func Category(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/items/category.html", gin.H{})
}
