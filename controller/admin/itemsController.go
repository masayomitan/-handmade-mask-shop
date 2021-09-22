package controller

import (
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"net/http"
  "fmt"
	_ "github.com/go-sql-driver/mysql"

)


func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{})
}

//create item
func Create(c *gin.Context) {
	// item := domain.Item{}
	// err := c.Bind(&item)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
// 	}
  c.HTML(http.StatusOK, "admin/item/create.html", gin.H{})
}

func Store(c *gin.Context) {
	database.GormConnect()
	var json domain.Item

	name := c.PostForm("")
	fmt.Println("name:", name,)

  // var item domain.Item
	// err := c.BindJSON(&item)
	// if err != nil{
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
  // }
	c.JSON(http.StatusCreated, gin.H{
		"status": name,
		"str": json.Name,
  })
	c.Redirect(http.StatusFound, "/admin/item/create")
}

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/detail.html", gin.H{})
}