package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	var err error

	dsn := "root@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("MySQL connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("Database error %v", DB.Error)
	}

	AutoMigrateAll()
}

func AutoMigrateAll() {
	err := DB.AutoMigrate(&Album{})
	if err != nil {
		fmt.Printf("Database automigrate error %v", err)
	}
}
