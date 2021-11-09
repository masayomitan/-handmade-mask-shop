package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"

)

type AdminUserRepository struct {}

func GetAdminUser(request []string) domain.AdminUser {

	var adminUser domain.AdminUser
	db := database.GormConnect()
	db.Table("adminUsers").Where("deleted_at IS NULL").First(&adminUser)
	return adminUser
} 