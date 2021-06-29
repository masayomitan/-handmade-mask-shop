package controller

import (
	"handmade_mask_shop/domain"
	"net/http"
  "fmt"
	"github.com/gin-gonic/gin"
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
	item := domain.Item{}
	s := c.PostForm("str")
		n := c.PostForm("num")
		b := c.PostForm("bool")
		message := fmt.Sprintf("s: %v, n: %v, b: %v, l: %v", s, n, b,)
		c.String(http.StatusOK, message)
	err := c.Bind(&item)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

}
