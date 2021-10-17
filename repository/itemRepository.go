package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	// "handmade_mask_shop/component"
	// "github.com/shopspring/decimal"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

)

type ItemRepository struct {}

func GetDataById(id string) domain.Item {

	var item domain.Item
	db := database.GormConnect()
	db.Where("ID =? AND items.deleted_at IS NULL",id, "NULL").Preload("ItemImages").First(&item)

	return item
} 

func GetAllDisplayItems() domain.Items {
	var items domain.Items
  db := database.GormConnect()

	db.Table("items").Where("display_flg = 1 AND items.deleted_at IS NULL").Preload("ItemImages").Find(&items)

	return items
}

func SaveItem(data *domain.Item) uint {
	fmt.Println()

	db := database.GormConnect()
	now := time.Now()
  data.CreatedAt = now
  data.UpdatedAt = now
  
	db.Create(&data)
	return data.ID
}