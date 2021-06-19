package domain

import (
	_ "github.com/go-sql-driver/mysql"

	// "gorm.io/gorm"

	"time"
)

type User struct {
	ID        		  uint             `gorm:"primaryKey"`

	FamilyName      string
	firstName       string
	FamilyNameKana  string
	firstNameKana   string

	Email           *string
	Age          		uint8

	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Users []User

type UserUsecase interface {
	FindAll() ([]User, error)
	Update(id int) error
	Store(user User) error
	Delete(id int) error
	Search(key string) ([]User, error) // 追加
}