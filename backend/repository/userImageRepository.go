package repository

import (
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"time"
	_ "github.com/go-sql-driver/mysql"
)



func SaveAdminUserImage(fileName string) (id uint){
	fmt.Println()
  filePath := "/public/images/"
	db := database.GormConnect()

	var data domain.CostomerImage

	data.File_name = fileName
	data.File_path = filePath + fileName

	now := time.Now()
  data.CreatedAt = now
  data.UpdatedAt = now  
	db.Create(&data)
	return data.ID
}
