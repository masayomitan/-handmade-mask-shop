package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"time"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type ItemImageRepository struct {}

type ItemImage struct {
	domain.ItemImage
}

type ItemsItemImage struct {
	domain.ItemsItemImage
}

var itemImages domain.ItemImages


func GetAllItemImages() (*domain.ItemImages, error) {
	fmt.Println()
	db := database.GormConnect()
	if result := db.Table("item_images").
		Select("ID", "File_path", "orderby").
		Where("deleted_at IS NULL").
		Find(&itemImages); result.Error != nil {
		return &itemImages, result.Error
	}
	return &itemImages, nil
}


func (i *ItemImage) SaveItemImage(userId uint, fileName string) (*ItemImage, error) {
  filePath := "/public/img/"
	db := database.GormConnect()
  
	i.AdminUserID = userId
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


func (i * ItemsItemImage) SaveItemImageIds(itemId uint, imageId string) (*ItemsItemImage, error) {
	db := database.GormConnect()
	u64, _ := strconv.ParseUint(imageId, 10, 64)
	u := uint(u64)
	i.ItemImageID = u
	i.ItemID = itemId

	now := time.Now()
  i.CreatedAt = now
  i.UpdatedAt = now  
	if result := db.Create(&i); result.Error != nil {
		return i, result.Error
	}
	return i, nil
}