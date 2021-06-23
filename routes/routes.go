package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Top(c *gin.Context) {
    c.HTML(http.StatusOK, "top/index.html", gin.H{})
}

func TopDetail(c *gin.Context){
	c.HTML(http.StatusOK, "top/detail.html", gin.H{})
}
