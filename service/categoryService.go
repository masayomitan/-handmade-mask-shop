package service

import (
	"fmt"
	"handmade_mask_shop/repository"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"

)

type CategoryService struct {}


func GetJsonAllCategories() []string {
	fmt.Println()
	Categories, _ := repository.GetAllCategories()
	jsonCategories, _ := json.Marshal(Categories)
	result := [] string{string(jsonCategories)}
  return result
}