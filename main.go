package main

import (
		"github.com/gin-gonic/gin"

		"net/http"
)

func main() {
	r := gin.Default()


	r.LoadHTMLGlob("./templates/*")


		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
		})
	})

		r.GET("/detail", func(c *gin.Context) {
			c.HTML(http.StatusOK, "detail.html", gin.H{
				
			})
	})	



	r.Run(":80")
};