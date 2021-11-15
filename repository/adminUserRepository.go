package repository

import (
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AdminUserRepository struct {}

func GetLoginAdminUser(request map[string]string) (domain.AdminUser) {
  username := request["Username"]
	password := request["Password"]

	var adminUser domain.AdminUser
	db := database.GormConnect()
	if err := db.Table("admin_users").Where("username = ? AND deleted_at IS NULL", username, "NULL").First(&adminUser)
	err.Error != nil {
		fmt.Println("missing username")
		return adminUser
	}
	
	hashedPassword := adminUser.Password

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println ("missing password", err)
  }
	return adminUser
}