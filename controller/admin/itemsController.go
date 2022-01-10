package controller

import (
	"fmt"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"

	"github.com/gin-gonic/gin"

	// "github.com/gin-contrib/sessions"
	"net/http"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
)


func ItemIndex(c *gin.Context) {
	data := repository.GetAllItems()
	c.HTML(http.StatusOK, "admin/items/index.html", gin.H{
		"data" : data,
	})
}


func ItemCreate(c *gin.Context) {
	csrf := csrf.Token
	// session := sessions.Default(c)
  categories := service.GetJsonAllCategories()
  c.HTML(http.StatusOK, "admin/items/create.html", gin.H{
		"categories": categories,
		"CSRFField" :  csrf,
	})
}


func ItemEdit(c *gin.Context) {
	csrf := csrf.Token
	id := c.Param("id")
	item, err := repository.GetItemByID(id)
	jsonItem, _ := json.Marshal(item)
	jsonItemArr := []string{string(jsonItem)}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adminUser not found"})
	}
	categories := service.GetJsonAllCategories()
  c.HTML(http.StatusOK, "admin/items/edit.html", gin.H{
		"item" : jsonItemArr,
		"categories": categories,
		"CSRFField" :  csrf,
	})
}


func ItemDetail(c *gin.Context) {
	id := c.Param("id")
	item, err := repository.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adminUser not found"})
	}
	c.HTML(http.StatusOK, "admin/items/detail.html", gin.H{
		"item" : item,
	})
}


func Complete(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/items/complete.html", gin.H{})
}


func Category(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/items/category.html", gin.H{})
}
