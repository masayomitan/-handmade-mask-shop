package domain

import (
	_ "github.com/go-sql-driver/mysql"
	// "handmade_mask_shop/repository"
	"time"
	"gorm.io/gorm"
)

type AdminUser struct {
	ID        		 uint     `gorm:"primaryKey"`
  Username       string   `form:"username"`
	Email          string   `form:"email"`
	Password       string   `form:"password"`
	User_imageID   uint  		`form:"user_image_id"`
	
	Reset_key      string   `json:"reset_key"`
	Reset_expire   string   `json:"reset_expire"`

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    gorm.DeletedAt 		`form:"deleted_at" gorm:"default:'null'"`
	
	UserImage UserImage
}

type AdminUsers []AdminUser

type SetAdminUser struct {
	ID       interface{}
	Username interface{}
	Email    interface{}
}

type CrudAdminUser interface {
	UpdateAdminUser(id uint, request map[string]string) (AdminUser)
}