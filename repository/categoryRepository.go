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

type CategoryRepository struct {}

var category domain.Category
var categories domain.Categories

func GetAllCategories() (domain.Categories, error) {
	fmt.Println()
	db := database.GormConnect()
	if result := db.Table("categories").
		Where("deleted_at IS NULL").
		Find(&categories); result.Error != nil {
		return categories, result.Error
	}
	return categories, nil
} 


func SaveCategory(category *domain.Category) (*domain.Category, error) {

	db := database.GormConnect()
	// category.Name = name
	now := time.Now()
  category.CreatedAt = now
  category.UpdatedAt = now
	if result := db.Create(&category); result.Error != nil {
		return category, result.Error
	}
	return category, nil
}


func CheckExistsByCategoryName(name string) (bool) {
	db := database.GormConnect()

	if result := db.Table("categories").
	Where("name = ? AND deleted_at IS NULL", name).
	Find(&name); result.Error != nil {
		//record exists
		return true
	}
	//record not found
	return false
}