package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	// "fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func GormConnect() *gorm.DB {

	err := godotenv.Load((".env"))
    if err != nil {
			panic ("envファイルの読み込みに失敗しました。")
    }

	var USER string
	var PASS string
	var PROTOCOL string
	var DBNAME string
	
	env := os.Getenv("ENV")
  if env == "dev" {
		USER = os.Getenv("DB_USER")
		PASS = os.Getenv("DB_PASS")
		DBNAME = os.Getenv("DB_DATABASE")
		PROTOCOL = os.Getenv("DB_PROTOCOL")
	}

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
		if err != nil {
			panic ("データベースとの通信に失敗しました。")
		}
	return db
}

