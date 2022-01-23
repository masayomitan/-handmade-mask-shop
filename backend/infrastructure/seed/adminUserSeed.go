package seed

import (
	_ "github.com/go-sql-driver/mysql"
	"handmade_mask_shop/domain"
	"fmt"
	"time"
	"gorm.io/gorm"
	"handmade_mask_shop/component"
)
var adminUser domain.AdminUser


func AdminUserSeed(db *gorm.DB) {
	fmt.Println()
	var now = time.Now()
  pass := component.HashPassword("test")
	adminUser.Username = "hakuto"
	adminUser.Password = pass
  adminUser.CreatedAt = now
  adminUser.UpdatedAt = now
	if err := db.Create(&adminUser).Error; err != nil {
		fmt.Printf("%+v", err)
	}
}

