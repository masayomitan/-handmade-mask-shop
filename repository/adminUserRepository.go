package repository

import (
	"fmt"
	"time"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

var adminUser domain.AdminUser

func GetLoginAdminUserByRequest(request map[string]string) (domain.AdminUser, error) {
  username := request["username"]
	db := database.GormConnect()
	if result := db.Table("admin_users").Where("username = ? AND deleted_at IS NULL", username).First(&adminUser); result.Error != nil {
		fmt.Println("wrong or missing username")
		return adminUser, result.Error
	}
	return adminUser, nil
}


func GetAdminUserByID(id uint) (domain.AdminUser, error) {
	db := database.GormConnect()

	if result := db.Table("admin_users").
	Where("admin_users.id = ? AND admin_users.deleted_at IS NULL", id).
	Preload("UserImage").
	First(&adminUser).Debug(); result.Error != nil {
		return adminUser, result.Error
	}

	fmt.Println(adminUser.UserImage.File_name)
	return adminUser, nil
}


func UpdateAdminUser (id uint, imageID uint, request map[string]string) (domain.AdminUser, error) {
	db := database.GormConnect()
	now := time.Now()

	adminUser.ID = id
	adminUser.Username = request["username"]
	adminUser.Password = request["password"]
	
	adminUser.User_imageID = imageID
  adminUser.UpdatedAt = now

	if result := db.Updates(&adminUser); result.Error != nil {
		return adminUser, result.Error
	}
	return adminUser, nil
}