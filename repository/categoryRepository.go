package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	// "handmade_mask_shop/component"
	// "github.com/shopspring/decimal"
	"time"
	_ "github.com/go-sql-driver/mysql"

)

type CategoryRepository struct {}

var category domain.Category
var categories domain.Categories
var now = time.Now()
var db = database.GormConnect()

func FindCategoryByID(id uint) (*domain.Category, error) {
	if result := db.Table("categories").
	Where("id = ? AND deleted_at IS NULL", id).
	First(&categories); result.Error != nil {
		return &category, result.Error
	}
	return &category, nil
}


func GetAllCategories() (domain.Categories, error) {
	if result := db.Table("categories").
		Where("deleted_at IS NULL").
		Find(&categories); result.Error != nil {
		return categories, result.Error
	}
	return categories, nil
} 


func SaveCategory(category *domain.Category) (*domain.Category, error) {
  category.CreatedAt = now
  category.UpdatedAt = now
	if result := db.Create(&category); result.Error != nil {
		return category, result.Error
	}
	return category, nil
}


func UpdateCategory(id uint, category *domain.Category) (*domain.Category, error) {
	category.UpdatedAt = now
	if result := db.Model(&categories).
		Where("id = ?", id).
		Update("name", category.Name); result.Error != nil {
		return category, result.Error
	}
	return category, nil
}


//logical delete
func DeleteCategory(id uint) (error)  {
  db.Table("categories").
	Where("id = ?", id).
	Delete(&category)
  return nil
}


func CheckExistsByCategoryName(name string) (bool) {
	if result := db.Table("categories").
	Where("name = ? AND deleted_at IS NULL", name).
	Find(&name); result.Error != nil {
		//record exists
		return true
	}
	//record not found
	return false
}
