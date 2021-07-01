package controller

import (
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/domain"
	"net/http"
  "fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"

)

//create item
func ItemsCreate(c *gin.Context) {
	item := domain.Item{}
	err := c.Bind(&item)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

}

func ItemsStore(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/handmade_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	var item domain.Item
	c.BindJSON(&item)
	fmt.Println(item.Name)
	c.JSON(200, gin.H{
		"message": item.Name,
	})

	err := c.Bind(&item)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

}
