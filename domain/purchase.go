package domain

import (
	_ "github.com/go-sql-driver/mysql"

	// "gorm.io/gorm"

	"time"
)

type Purchase struct {

	ID           uint             `gorm:"primaryKey"`
	Item_id      uint
	User_id      uint
	Item_name    string
  
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Purcahses []Purchase

type PurcahsesUsecase interface {
	FindAll() ([]User, error)
	Update(id int) error
	Store(user User) error
	Delete(id int) error
	Search(key string) ([]User, error)
}