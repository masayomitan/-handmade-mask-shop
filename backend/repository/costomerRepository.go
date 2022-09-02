package repository

import (
	"fmt"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)


type Costomer struct {
	*domain.Costomer
}

type Costomers struct {
	*domain.Costomers
}

var costomer domain.Costomer
var costomers domain.Costomers

func GetCostomers() (*domain.Costomers, error) {
	db := database.GormConnect()
	if result := db.Table("costomers").
		Where("costomers.deleted_at IS NULL").
		Preload("CostomerImages").
		Find(&costomers); result.Error != nil {
		return &costomers, result.Error
	}
	return &costomers, nil
}


func (c *Costomer) SaveCostomer(costomer *domain.Costomer) (*Costomer, error){
	db := database.GormConnect()
	fmt.Println(costomer.Email)
	// fmt.Println(c.Email)
	// c.Email = costomer.Email
	if result := db.Create(&costomer); result.Error != nil {
		return c, result.Error
	}
	return c, nil
}

func (c *Costomer) GetCostomerByEmail(email string) (*Costomer, error) {
	if result := db.Table("costomers").
		Where("costomers.email = ?", email).
		Preload("CostomerImages").
		First(&c); result.Error != nil {
		return c, result.Error
}
	return c, nil
}