package controller

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	// "handmade_mask_shop/repository"
	"net/http"
	_ "github.com/go-sql-driver/mysql"

)

// func Index(c *gin.Context) {

// 	fmt.Println()
// 	data := repository.GetAllDisplayItems()
// 	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{
// 		"data" : data,
// 	})
// }

func Detail(c *gin.Context) {
	// id := c.Param("id")
	// item := repository.GetDataById(id)
	c.HTML(http.StatusOK, "front/item/detail.html", gin.H{
		// "item" : item,
	})
}