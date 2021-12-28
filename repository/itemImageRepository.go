package repository

import (
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type ItemImageRepository struct {}

type ItemImage struct {
	domain.ItemImage
}
var image domain.ItemImage

func (i *ItemImage) SaveItemImage (fileName string) (*ItemImage, error) {
	fmt.Println(fileName)
  filePath := "/public/img/"
	db := database.GormConnect()

	i.File_name = fileName
	i.File_path = filePath + fileName

	now := time.Now()
  i.CreatedAt = now
  i.UpdatedAt = now  
	if result := db.Create(&i); result.Error != nil {
		return i, result.Error
	}
	return i, nil
}