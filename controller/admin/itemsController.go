package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

)


func Index(c *gin.Context) {
	fmt.Println()
	data := repository.GetAllDisplayItems()

	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{
		"data" : data,
	})
}

//create item with files
func Create(c *gin.Context) {

  c.HTML(http.StatusOK, "admin/item/create.html", gin.H{})
}

func Store(c *gin.Context) {

  //saving formData
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
		})
		return
  }
  
	// saving fileData
	service.ResizeFile(filePath)
	repository.SaveItemImage(newFileName, itemId)
}

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/detail.html", gin.H{})
}