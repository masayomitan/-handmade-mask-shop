package component

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"time"
  // "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SaveData(data *domain.Item) {

	db := database.GormConnect()
  
	now := time.Now()
  data.CreatedAt = now
  data.UpdatedAt = now
  
	db.Create(&data)
	return
}