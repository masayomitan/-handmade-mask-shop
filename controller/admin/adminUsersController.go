package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "handmade_mask_shop/domain"
	// "handmade_mask_shop/repository"
	// "handmade_mask_shop/service"
	"net/http"
	// "os"
	_ "github.com/go-sql-driver/mysql"

)

func AdminRegist(c *gin.Context) {
	fmt.Println()
  // categories := service.GetJsonAllCategories()
  c.HTML(http.StatusOK, "admin/adminUser/regist.html", gin.H{
		// "categories": categories,
	})
}

// func AdminEdit(c *gin.Context) {
//   fmt.Println()
// 	var item domain.Item
// 	err := c.Bind(&item)
//   if err != nil {
// 		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
// 		return
// 	}

// 	itemId := repository.SaveItem(&item)
// 	file, _ := c.FormFile("image")
// 	newFileName := service.RenameFile(file.Filename)
// 	imageDir := "./public/images/" 
// 	filePath := imageDir + newFileName

// 	if f, exist := os.Stat(imageDir); os.IsNotExist(exist) || 
// 	!f.IsDir() {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//       "message": "error. could not find dir",
//     })
//   }
// 	er := c.SaveUploadedFile(file, filePath)
// 	if er != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"message": "Unable to save the file",
// 		})
// 		return
//   }
  
// 	// saving fileData
// 	// service.ResizeFile(filePath)
// 	repository.SaveItemImage(newFileName, itemId)
// }
