package controller

import (
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"net/http"
	// "strconv"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var costomer domain.Costomer

func GetByEmail(c *gin.Context) {
	err := c.ShouldBindJSON(&costomer)
  if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

	p := repository.Costomer{}
  res, err := p.GetCostomerByEmail(costomer.Email)
	fmt.Println(err)
	b, _ := json.Marshal(res)
	_ = json.Unmarshal([]byte(b), &costomer)
	// fmt.Println(&costomer)
	
	if err != nil {
		res, err = p.SaveCostomer(&costomer);
		fmt.Println(res)
	}

	c.JSON(http.StatusOK, 5)
}
