package controller

import (
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"net/http"
  // "fmt"
	_ "github.com/go-sql-driver/mysql"
)


func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{})
}

//create item with files
func Create(c *gin.Context) {

  c.HTML(http.StatusOK, "admin/item/create.html", gin.H{})
}

func Store(c *gin.Context) {
  
	filePath := "../../public/images"
	// image, header, _ := c.Request.FormFile("image")
	// saveFile, _ := os.Create("./images/" + header.Filename)

  //saving formData
	var item domain.Item		
	err := c.Bind(&item)
  if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

	repository.SaveItem(&item)

	// repository.SaveItemImage(files)

	c.JSON(http.StatusCreated, gin.H{
		"name": item.Name,
		"detail": item.Detail,
  })
	c.Redirect(http.StatusFound, "/admin/item/create")
}

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/detail.html", gin.H{})
}