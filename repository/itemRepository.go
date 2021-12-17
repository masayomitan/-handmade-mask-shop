package repository

import (
	"fmt"
	"time"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"

)

type ItemRepository struct {}


func GetItemByID(id string) domain.Item {
	var item domain.Item
	db := database.GormConnect()
	db.Where("ID = ? AND items.deleted_at IS NULL", id, "NULL").
	Preload("ItemImages").
	First(&item)
	return item
} 

func GetAllItems() domain.Items {
	var items domain.Items
  db := database.GormConnect()
	db.Table("items").
	Where("items.deleted_at IS NULL").
	Preload("ItemImages").Find(&items)
	return items
}


func SaveItem(data *domain.Item) uint {
	fmt.Println(data)

	db := database.GormConnect()
	now := time.Now()
  data.CreatedAt = now
  data.UpdatedAt = now

	db.Create(&data)
	return data.ID
}