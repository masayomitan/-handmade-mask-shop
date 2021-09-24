package service

import (
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"net/http"
	"time"
  // "fmt"
	_ "github.com/go-sql-driver/mysql"
)

type ItemService struct {}

func (ItemService) InsertData(data *domain.Item) {

	db := database.GormConnect()
  
	now := time.Now()
  data.CreatedAt = now
  data.UpdatedAt = now
  
	res := db.Create(&data)
	return res
}