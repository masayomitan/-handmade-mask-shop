package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AdminUser struct {
	ID        		 uint     `gorm:"primaryKey"`
  Username       string   `form:"username"`
	Email          string   `form:"email"`
	Password       string   `form:"password"`

	Reset_key      string   `form:"reset_key"`
	Reset_expire   string   `form:"reset_expire"`

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`

}

type AdminUsers []AdminUser

type SetAdminUser struct {
	ID       interface{}
	Username interface{}
	Email    interface{}
}