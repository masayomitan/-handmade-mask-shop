package controller

import (
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
	"encoding/json"
	"net/http"
	"strconv"
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
	u64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(u64)
	item, err := repository.GetItemByID(id)
	jsonItem, _ := json.Marshal(item)
	jsonItemArr := []string{string(jsonItem)}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品が見つかりませんでした"})
	}
	categories := service.GetJsonAllCategories()
  c.HTML(http.StatusOK, "admin/items/edit.html", gin.H{
		"item" : jsonItemArr,
		"categories": categories,
		"CSRFField" :  csrf,
	})
}


func ItemDetail(c *gin.Context) {
	u64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(u64)
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
