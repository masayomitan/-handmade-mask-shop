package domain

import (
	_ "github.com/go-sql-driver/mysql"

	"time"
)

type  UserImage struct {
	ID           uint    `gorm:"primaryKey"`
	UserID       uint    `gorm:"user_id"`
	File_name    string  `form:"file_name"`
	File_path    string  `form:"file_path"`

	CreatedAt    time.Time  `form:"created_at" gorm:"NOT NULL"`
	UpdatedAt    time.Time  `form:"updated_at" gorm:"NOT NULL"`
	DeletedAt    time.Time  `form:"deleted_at" gorm:"default:'null'"`
  
	//belongsTo
	AdminUser AdminUser
	User User
}