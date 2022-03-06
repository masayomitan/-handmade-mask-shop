package repository

import (
	"fmt"
	"time"
	// "strconv"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"

)

// type ItemRepository struct {}

type Item struct {
	*domain.Item
}

type Items struct {
	*domain.Items
}


func GetItemByID(id uint) (*domain.Item, error) {
	var item domain.Item
	db := database.GormConnect()
	if result := db.Where("ID = ? AND items.deleted_at IS NULL", id, "NULL").
	Preload("ItemImages").
	Preload("Category").
	First(&item); result.Error != nil {
		fmt.Println(result.Error)
		return &item, result.Error
	}
	return &item, nil
} 


func GetAllItems() (*domain.Items, error) {
	var items domain.Items
  db := database.GormConnect()
	if result := db.Table("items").
		Where("items.deleted_at IS NULL").
		Preload("ItemImages").
		Preload("Category").
		Find(&items);  result.Error != nil {
			fmt.Println(result.Error)
			return &items, result.Error
		}
		return &items, nil
	}


func GetAllItemsByDisplayFlg(flg uint) (*domain.Items, error) {
	var items domain.Items
  db := database.GormConnect()
	if result := db.Table("items").
		Where("items.display_flg = ? AND items.deleted_at IS NULL", flg).
		Preload("ItemImages").
		Preload("Category").
		Find(&items); result.Error != nil {
			fmt.Println(result.Error)
			return &items, result.Error
		}
		return &items, nil
}


func (i *Item) GetDisplayItem (id uint) (*Item, error) {
	db := database.GormConnect()
	if result := db.Where("ID = ? AND items.display_flg = ?", id, 1).
		Preload("ItemImages").
		Preload("Category").
		First(&i); result.Error != nil {
			fmt.Println(result.Error)
			return i, result.Error
	}
	return i, nil
} 


func GetDisplayItems() (*domain.Items, error) {
	var items domain.Items
  db := database.GormConnect()
	if result := db.Table("items").
		Where("items.display_flg = ?", 1).
		Preload("ItemImages").
    Find(&items); result.Error != nil {
		fmt.Println(result.Error)
		return &items, result.Error
	}
	return &items, nil
}


func GetDisplayItemsCategoryId(id uint) (*domain.Items, error) {
	var items domain.Items
  db := database.GormConnect()
	if result := db.Table("items").
		Where("items.display_flg = ? AND items.category_id = ?", 1, id).
		Preload("ItemImages").
    Find(&items); result.Error != nil {
		fmt.Println(result.Error)
		return &items, result.Error
	}
	return &items, nil
}



func SaveItem(item *domain.Item) (*domain.Item, error) {

	db := database.GormConnect()
	now := time.Now()
  item.CreatedAt = now
  item.UpdatedAt = now

	if result := db.Create(&item); result.Error != nil {
		return item, result.Error
	}
	return item, nil
}



func UpdateItem(id uint, item *domain.Item) (*domain.Item, error) {

	db := database.GormConnect()
	now := time.Now()
	item.ID = id
  item.UpdatedAt = now

	if result := db.Updates(&item).
		Where("id deleted IS NUll", id); result.Error != nil {
		return item, result.Error
	}
	return item, nil
}