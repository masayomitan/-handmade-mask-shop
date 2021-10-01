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
	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{})
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

	id := repository.SaveItem(&item)

	file, _ := c.FormFile("image")
	newFileName := service.RenameFile(file.Filename)
	filePath := "./public/images/"

	if f, exist := os.Stat(filePath); os.IsNotExist(exist) || 
	!f.IsDir() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "message": "error. could not find dir",
    })
  }

	er := c.SaveUploadedFile(file, filePath + newFileName)
	if er != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
		})
		return
  }
  
	// saving fileData
	repository.SaveItemImage(newFileName, id)


}

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/detail.html", gin.H{})
}