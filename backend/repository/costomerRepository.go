package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)


type Costomer struct {
	*domain.Costomer
}

type Costomers struct {
	*domain.Costomers
}

var costomer domain.Costomer
var costomers domain.Costomers

func GetCostomers() (*domain.Costomers, error) {
	db := database.GormConnect()

	if result := db.Table("costomers").
	Where("costomers.deleted_at IS NULL").
	// Preload("CostomerImage").
	Find(&costomers); result.Error != nil {
		return &costomers, result.Error
	}
	return &costomers, nil
}
