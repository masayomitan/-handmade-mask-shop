package seed

import (
	_ "github.com/go-sql-driver/mysql"
	// "handmade_mask_shop/infrastructure/database"
	"handmade_mask_shop/domain"
	"fmt"
	"time"
	"gorm.io/gorm"
)


func CategorySeed(db *gorm.DB) error {
	var Category domain.Category
	now := time.Now()

	for i, v := range []string{"マスク", "ピアス", "リング", "その他"} {
		Category.ID += 1
		Category.Name = v
		Category.Orderby = (i + 1)
		Category.CreatedAt = now
		Category.UpdatedAt = now
		if err := db.Create(&Category).Error; err != nil {
      fmt.Printf("%+v", err)
    }
	}
	return nil

}