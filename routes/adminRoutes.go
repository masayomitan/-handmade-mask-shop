package routes

import (
	"net/http"
  // "fmt"
	"github.com/gin-gonic/gin"
)

func getRouteAdminItemGroup() *gin.Engine {
    
    // return routes/
}

func AdminDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/dashboard/index.html", gin.H{})
}

func AdminItem(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/index.html", gin.H{})
}

func AdminItemDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/detail.html", gin.H{})
}

func AdminItemCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/item/create.html", gin.H{})
}

func AdminItemStore(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
