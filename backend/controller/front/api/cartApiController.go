package controller

import (
	// "fmt"
	"fmt"
	"handmade_mask_shop/domain"
	// "handmade_mask_shop/repository"
	// "net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)

var order domain.Order
var orderItem domain.OrderItem

func AddOrder(c * gin.Context) {
	c.ShouldBindJSON(&orderItem)
	fmt.Println(&orderItem)

}