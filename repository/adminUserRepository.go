package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type AdminUserRepository struct {}

func GetLoginAdminUserByRequest(request map[string]string) (domain.AdminUser, error) {
  username := request["Username"]

	var adminUser domain.AdminUser
	db := database.GormConnect()
	if result := db.Table("admin_users").Where("username = ? AND deleted_at IS NULL", username, "NULL").First(&adminUser)
	result.Error != nil {
		fmt.Println("wrong or missing username")
		return adminUser, result.Error
	}
	return adminUser, nil
}

// func GetAdminUserById(id string) (domain.AdminUser){

// }