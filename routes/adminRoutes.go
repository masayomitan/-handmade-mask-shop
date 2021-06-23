package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
		
)

func AdminDashboard(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/dashboard/index.html", gin.H{})
}

func AdminItem(c *gin.Context){
	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{})
}

func AdminItemDetail(c *gin.Context){
	c.HTML(http.StatusOK, "admin/item/detail.html", gin.H{})
}

func AdminItemCreate(c *gin.Context){
	c.HTML(http.StatusOK, "admin/item/create.html", gin.H{})
}