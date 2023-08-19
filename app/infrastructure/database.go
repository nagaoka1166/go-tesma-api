// app/infrastructure/database.go
package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB_SOURCE")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})	
	if err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}

	fmt.Println("DBに接続しました!")

	return db
}

func GetDB() *gorm.DB {
	return db
}